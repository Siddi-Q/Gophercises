package db

import (
	"encoding/binary"
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

// Task represents a user's task.
type Task struct {
	ID            int
	Description   string
	Completed     bool
	CompletedDate time.Time
}

// InitDB takes a path and initializes a boltdb at that path.
func InitDB(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		return err
	}

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

// CreateTask takes a task description, creates a task struct and inserts it into the database.
func CreateTask(description string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)

		task := &Task{
			ID:          id,
			Description: description,
			Completed:   false,
		}

		taskEnc, err := json.Marshal(task)

		if err != nil {
			return err
		}

		return b.Put(key, taskEnc)
	})

	if err != nil {
		return -1, err
	}

	return id, nil
}

// ReadAllTasks returns all tasks in the database.
func ReadAllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var task *Task
			err := json.Unmarshal(v, &task)

			if err != nil {
				return err
			}

			tasks = append(tasks, *task)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// ReadSomeTasks returns a list of tasks based on whether the tasks are completed or not.
func ReadSomeTasks(completed bool) ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var task *Task
			err := json.Unmarshal(v, &task)

			if err != nil {
				return err
			}

			if task.Completed == completed {
				tasks = append(tasks, *task)
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// UpdateTaskCompleted will
func UpdateTaskCompleted(key int) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		key := itob(key)
		value := b.Get(key)

		var task Task
		err := json.Unmarshal(value, &task)
		if err != nil {
			return err
		}

		task.Completed = true
		task.CompletedDate = time.Now()

		taskEnc, err := json.Marshal(task)
		if err != nil {
			return err
		}

		return b.Put(key, taskEnc)
	})

	if err != nil {
		return err
	}

	return nil
}

// UpdateTaskDescription will
func UpdateTaskDescription(key int, description string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		key := itob(key)
		value := b.Get(key)

		var task Task
		err := json.Unmarshal(value, &task)
		if err != nil {
			return err
		}

		task.Description = description

		taskEnc, err := json.Marshal(task)
		if err != nil {
			return nil
		}

		return b.Put(key, taskEnc)
	})

	if err != nil {
		return err
	}

	return nil
}

// DeleteTask will
func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

// DeleteBucket will
func DeleteBucket() error {
	return db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket(taskBucket)
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

package main

import (
	"bytes"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "phonedb"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	db, err := sql.Open("postgres", psqlInfo)
	must(err)
	err = resetDB(db, dbname)
	must(err)
	db.Close()

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	must(createPhoneNumbersTable(db))
	id, err := insertPhoneNumber(db, "1234567890")
	must(err)
	fmt.Printf("id = %d\n", id)
	id, err = insertPhoneNumber(db, "2345678901")
	must(err)
	fmt.Printf("id = %d\n", id)

	phoneNumber, err := getPhoneNumber(db, id)
	must(err)
	fmt.Printf("phone number = %s\n", phoneNumber)

	phoneNumbers, err := getAllPhoneNumbers(db)
	must(err)
	for _, pn := range phoneNumbers {
		fmt.Printf("%+v\n", pn)
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func resetDB(db *sql.DB, dbName string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + dbName)
	if err != nil {
		return err
	}
	return createDB(db, dbName)
}

func createDB(db *sql.DB, dbName string) error {
	_, err := db.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		return err
	}
	return nil
}

func createPhoneNumbersTable(db *sql.DB) error {
	sqlStatement := `
		CREATE TABLE IF NOT EXISTS phone_numbers (
			id SERIAL,
			value VARCHAR(255)
		)`
	_, err := db.Exec(sqlStatement)
	return err
}

func insertPhoneNumber(db *sql.DB, phoneNumber string) (int, error) {
	sqlStatement := `INSERT INTO phone_numbers (value) VALUES ($1) RETURNING id`
	var id int
	err := db.QueryRow(sqlStatement, phoneNumber).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func getPhoneNumber(db *sql.DB, id int) (string, error) {
	sqlStatement := `SELECT value FROM phone_numbers WHERE id=$1`
	var phoneNumber string
	err := db.QueryRow(sqlStatement, id).Scan(&phoneNumber)
	if err != nil {
		return "", err
	}
	return phoneNumber, nil
}

type phoneNumber struct {
	id     int
	number string
}

func getAllPhoneNumbers(db *sql.DB) ([]phoneNumber, error) {
	sqlStatement := `SELECT id, value FROM phone_numbers`
	rows, err := db.Query(sqlStatement)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var phoneNumbers []phoneNumber
	for rows.Next() {
		var pn phoneNumber
		if err := rows.Scan(&pn.id, &pn.number); err != nil {
			return nil, err
		}
		phoneNumbers = append(phoneNumbers, pn)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return phoneNumbers, nil
}

func normalizePhoneNumber(pn string) string {
	var buf bytes.Buffer
	for _, rune := range pn {
		if rune >= '0' && rune <= '9' {
			buf.WriteRune(rune)
		}
	}
	return buf.String()
}

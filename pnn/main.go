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

	number, err := getPhoneNumber(db, id)
	must(err)
	fmt.Printf("phone number = %s\n", number)

	phoneNumbers, err := getAllPhoneNumbers(db)
	must(err)
	for _, pn := range phoneNumbers {
		fmt.Printf("%+v\n", pn)
	}

	pn, err := findPhoneNumber(db, "1234567890")
	must(err)
	fmt.Printf("%+v\n", *pn)

	pn2 := phoneNumber{1, "111111111"}
	updatePhoneNumber(db, pn2)

	number, err = getPhoneNumber(db, pn2.id)
	must(err)
	fmt.Printf("phone number = %s\n", number)

	deletePhoneNumber(db, pn2.id)
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

func findPhoneNumber(db *sql.DB, number string) (*phoneNumber, error) {
	sqlStatement := `SELECT id, value FROM phone_numbers WHERE value=$1`
	var pn phoneNumber
	err := db.QueryRow(sqlStatement, number).Scan(&pn.id, &pn.number)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &pn, nil
}

func updatePhoneNumber(db *sql.DB, pn phoneNumber) error {
	sqlStatement := `UPDATE phone_numbers SET value=$2 WHERE id=$1`
	_, err := db.Exec(sqlStatement, pn.id, pn.number)
	return err
}

func deletePhoneNumber(db *sql.DB, id int) error {
	sqlStatement := `DELETE FROM phone_numbers WHERE id=$1`
	_, err := db.Exec(sqlStatement, id)
	return err
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

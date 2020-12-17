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

	must(db.Ping())
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

func normalizePhoneNumber(pn string) string {
	var buf bytes.Buffer
	for _, rune := range pn {
		if rune >= '0' && rune <= '9' {
			buf.WriteRune(rune)
		}
	}
	return buf.String()
}

package main

import (
	"database/sql"
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "lyx0"
	// password = ""
	dbname = "gophercises_phone"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s sslmode=disable", host, port, user)
	db, err := sql.Open("postgres", psqlInfo)
	must(err)

	err = resetDB(db, dbname)
	must(err)

	db.Close()

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	// must(db.Ping())
	must(createPhoneNumbersTable(db))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func createPhoneNumbersTable(db *sql.DB) error {
	statement := `
		CREATE TABLE IF NOT EXISTS phone_numbers (
			id SERIAL, 
			value VARCHAR(255)
		)`
	_, err := db.Exec(statement)
	return err
}

func resetDB(db *sql.DB, name string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		return err
	}
	fmt.Printf("DROPPED database: %s\n", name)
	return createDB(db, name)
}

func createDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name)
	if err != nil {
		return err
	}
	fmt.Printf("CREATED database: %s\n", name)
	return nil
}

func normalize(phone string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(phone, "")
}

// func normalize(phone string) string {
// 	var buf bytes.Buffer
// 	for _, ch := range phone {
// 		if ch >= '0' && ch <= '9' {
// 			buf.WriteRune(ch)
// 		}
// 	}
//
// 	return buf.String()
// }

package main

import (
	"database/sql"
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
	phonedb "github.com/lyx0/gophercises/phone/phonedb"
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
	must(phonedb.Reset("postgres", psqlInfo, dbname))

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	must(phonedb.Migrate("postgres", psqlInfo))

	db, err := phonedb.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	if err := db.Seed(); err != nil {
		panic(err)
	}

	// phones, err := allPhones(db)
	// must(err)
	// for _, p := range phones {
	// 	fmt.Printf("Working on... %+v\n", p)
	// 	number := normalize(p.number)
	// 	if number != p.number {
	// 		fmt.Println("Updating or removing...", number)
	// 		existing, err := findPhone(db, number)
	// 		must(err)
	// 		if existing != nil {
	// 			// delete this number
	// 			must(deletePhone(db, p.id))
	// 			fmt.Printf("Deleted ID %d\n", p.id)
	// 		} else {
	// 			// update this number
	// 			p.number = number
	// 			must(updatePhone(db, p))
	// 			fmt.Printf("Updated ID %d\n", p.id)
	// 		}
	// 	} else {
	// 		fmt.Println("No changes required")
	// 	}
	// 	// fmt.Printf("ID: %d, Number: %s\n", p.id, p.number)
	// }
}

func findPhone(db *sql.DB, number string) (*phone, error) {
	var p phone
	row := db.QueryRow("SELECT * FROM phone_numbers WHERE value=$1", number)
	err := row.Scan(&p.id, &p.number)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &p, nil
}

func deletePhone(db *sql.DB, id int) error {
	statement := `DELETE FROM phone_numbers WHERE id=$1`
	_, err := db.Exec(statement, id)
	return err

}

func updatePhone(db *sql.DB, p phone) error {
	statement := `UPDATE phone_numbers SET value=$2 WHERE id=$1`
	_, err := db.Exec(statement, p.id, p.number)
	return err
}

func must(err error) {
	if err != nil {
		panic(err)
	}
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

//	// ###### Creating/Resetting Database ########
//	db, err := sql.Open("postgres", psqlInfo)
//	must(err)
//
//	err = resetDB(db, dbname)
//	must(err)
//
//	db.Close()
//	// ###########################################

// ######### Create Table Structures ##############
// func createPhoneNumbersTable(db *sql.DB) error {
// 	statement := `
// 		CREATE TABLE IF NOT EXISTS phone_numbers (
// 			id SERIAL,
// 			value VARCHAR(255)
// 		)`
// 	_, err := db.Exec(statement)
// 	return err
// }
// ################################################

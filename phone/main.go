package main

import (
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
	phonedb "github.com/lyx0/gophercises/phone/db"
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

	phones, err := db.AllPhones()
	must(err)
	for _, p := range phones {
		fmt.Printf("Working on... %+v\n", p)
		number := normalize(p.Number)
		if number != p.Number {
			fmt.Println("Updating or removing...", number)
			existing, err := db.FindPhone(number)
			must(err)
			if existing != nil {
				must(db.DeletePhone(p.ID))
			} else {
				p.Number = number
				must(db.UpdatePhone(&p))
			}
		} else {
			fmt.Println("No changes required")
		}
	}

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

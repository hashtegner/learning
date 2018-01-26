package main

import (
	"bytes"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"strconv"
)

type Phone struct {
	Id     int64
	Number string
}

func main() {
	db, err := sql.Open("postgres", "postgresql://postgres@localhost/gophercises?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	phones, err := read(db)

	if err != nil {
		log.Fatal(err)
	}

	normalize(phones)

	err = write(db, phones)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Done!")
}

func write(db *sql.DB, phones []*Phone) error {
	stmt, err := db.Prepare("UPDATE phone_numbers SET number = $1 WHERE id = $2")

	if err != nil {
		return err
	}

	defer stmt.Close()

	for _, phone := range phones {
		_, err := stmt.Exec(phone.Number, phone.Id)

		if err != nil {
			return err
		}
	}

	return nil
}

func read(db *sql.DB) ([]*Phone, error) {
	rows, err := db.Query("select id, number from phone_numbers")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var phones []*Phone

	for rows.Next() {
		phone := Phone{int64(1123), "foo"}

		rows.Scan(&phone.Id, &phone.Number)

		phones = append(phones, &phone)
	}

	return phones, nil

}

func normalize(phones []*Phone) {
	for _, phone := range phones {
		phone.Number = normalizeNumber(phone.Number)
	}
}

func normalizeNumber(phone string) string {
	var normalized bytes.Buffer

	for i := 0; i < len(phone); i++ {
		char := phone[i]

		_, err := strconv.ParseInt(string(char), 10, 32)

		if err == nil {
			normalized.WriteString(string(char))
		}
	}

	return normalized.String()
}

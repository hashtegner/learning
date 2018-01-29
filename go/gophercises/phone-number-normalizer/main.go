package main

import (
	"fmt"
	"github.com/alesshh/learning/go/gophercises/phone-number-normalizer/normalizer"
	"github.com/alesshh/learning/go/gophercises/phone-number-normalizer/repo"
	_ "github.com/lib/pq"
)

func main() {
	conn := "postgresql://postgres@localhost/gophercises?sslmode=disable"

	repo, err := repo.Open("postgres", conn)
	must(err)

	defer repo.Close()

	err = repo.Seed()
	must(err)

	phones, err := repo.All()

	must(err)

	fmt.Println("After")
	printAll(phones)

	for _, phone := range phones {
		phone.Number = normalizer.Normalize(phone.Number)
	}

	err = repo.Update(phones)
	must(err)

	fmt.Println("Before")
	printAll(phones)

	fmt.Println("Done!")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func printAll(phones []*repo.Phone) {
	for _, phone := range phones {
		fmt.Printf("Phone %d => %s\n", phone.Id, phone.Number)
	}
}

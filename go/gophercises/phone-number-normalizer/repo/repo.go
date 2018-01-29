package repo

import "database/sql"

type Phone struct {
	Id     int64
	Number string
}

type Repo struct {
	db *sql.DB
}

func Open(driverName string, dataSourceName string) (*Repo, error) {
	db, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		return nil, err
	}

	return &Repo{db}, nil
}

func (repo Repo) Seed() error {
	phones := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
	}

	return repo.transaction(func() error {
		for _, phone := range phones {
			_, err := repo.db.Exec("INSERT INTO phone_numbers (number) VALUES ($1)", phone)

			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (repo Repo) All() ([]*Phone, error) {
	rows, err := repo.db.Query("select id, number from phone_numbers")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var phones []*Phone

	for rows.Next() {
		phone := &Phone{}

		rows.Scan(&phone.Id, &phone.Number)

		phones = append(phones, phone)
	}

	return phones, nil
}

func (repo Repo) Update(phones []*Phone) error {
	return repo.transaction(func() error {
		for _, phone := range phones {
			_, err := repo.db.Exec("UPDATE phone_numbers SET number = $1 WHERE id = $2", phone.Id, phone.Number)

			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (repo Repo) Close() error {
	return repo.db.Close()
}

type transactionFunc func() error

func (repo Repo) transaction(fn transactionFunc) error {
	db := repo.db

	tx, err := db.Begin()

	if err != nil {
		return err
	}

	err = fn()

	if err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit()
}

package resources

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

type database struct {
	db *sql.DB
}

func NewDataBase() *database {
	db, err := getDb()

	if err != nil {
		panic(err)
	}

	return &database{
		db: db,
	}
}

func getDb() (*sql.DB, error) {
	dbUrl := os.Getenv("DATABASE_URL")

	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Second * 10)

	return db, nil
}
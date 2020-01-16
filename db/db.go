package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	MySQL *sql.DB
}

type DBInterface interface {
	InitDB() error
	GetDB() *sql.DB
}

func (dbi *DB) InitDB() error {
	dbHost := os.Getenv("MYSQL_HOST")
	dbPORT := os.Getenv("MYSQL_PORT")
	dbNAME := os.Getenv("MYSQL_DATABASE")
	dbUSER := os.Getenv("MYSQL_USER")
	dbPASSWORD := os.Getenv("MYSQL_PASSWORD")

	connectStr := dbUSER + ":" + dbPASSWORD + "@tcp(" + dbHost + ":" + dbPORT + ")/" + dbNAME
	db, err := sql.Open("mysql", connectStr)

	if err != nil {
		log.Println("Error openning database")
		return err
	}
	dbi.MySQL = db
	return nil
}

func (dbi *DB) GetDB() *sql.DB {
	return dbi.MySQL
}

func CreateDB() DBInterface {
	return &DB{}
}

package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

/*
 need to install lib/pg,
 because database/sql provide generic interface, no built-in support for specific DB
 need a database drivers to implement the actual communication with the database.


  database/sql provides functions like:
 	sql.Open("postgres", connStr)
	db.Query("SELECT * FROM users")
	db.Exec("INSERT INTO users VALUES (...)")

 but how to talk to psql?
*/

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}

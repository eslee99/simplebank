package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/eslee99/simplebank/util"
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

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}

package tdb

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const DBDriver = "postgres"
const DBSource = "postgres://postgres:postgres@localhost:5432?sslmode=disable&database=tdb"

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {

	testDB, err := sql.Open(DBDriver, DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}

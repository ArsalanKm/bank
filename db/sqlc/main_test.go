package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/ArsalanKm/simple_bank/util"
	_ "github.com/lib/pq"
)

var testQueris *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadCofig("../..")

	if err != nil {
		log.Fatal("cannot load config file")
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect tot the database", err)
	}

	testQueris = New(testDB)
	os.Exit(m.Run())
}

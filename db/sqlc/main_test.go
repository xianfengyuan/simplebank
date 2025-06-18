package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/xianfengyuan/simplebank/util"
)

var testQueries *Queries
var testDB *pgxpool.Pool
var testStore Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	testStore = NewStore(testDB)

	os.Exit(m.Run())
	// This function is intentionally left empty.
	// It serves as a placeholder to ensure that the package compiles
	// and can be used in tests or other parts of the application.
	// You can add test cases or other logic here as needed.
}

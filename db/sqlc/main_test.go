package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	db "github.com/MeganViga/Plipkart/db/sqlc"
	"github.com/MeganViga/Plipkart/util"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)
var testQueries *Queries
const (
	dbEngine = "postgres"
	dbSource = "postgres://root:secret@localhost:5432/plipkart?sslmode=disable"
)
func TestMain(m *testing.M){
	db, err :=sql.Open(dbEngine,dbSource)
	if err != nil{
		log.Fatal("Unable to open db connection")
	}

	testQueries = New(db)
	os.Exit(m.Run())

}


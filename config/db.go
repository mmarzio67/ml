package config

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB
var DbUsers = map[string]User{}
var DbSessions = map[string]Session{} // session ID, session
var dbSessionsCleaned time.Time

func init() {
	//dbSource := getenv("DATABASE_URL", "postgres://ml:ml@localhost/ml?sslmode=disable")

	var err error

	dbSessionsCleaned = time.Now()

	//DB, err = sql.Open("postgres", dbSource)
	DB, err = sql.Open("postgres", "postgres://ml:medal!v!n9@localhost/ml?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your ml database.")
}

func getenv(k string, v string) string {
	if val := os.Getenv(k); val != "" {
		return val
	}
	os.Setenv(k, v)
	return v
}

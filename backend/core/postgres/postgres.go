package postgres

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/briand787b/piqlit/core/plog"
	"github.com/briand787b/piqlit/core/psql"

	"github.com/jmoiron/sqlx"

	// import the postgresql database driver
	_ "github.com/lib/pq"
)

const (
	dbHostEnvVar = "PL_DATABASE_HOST"
	dbNameEnvVar = "PL_DATABASE_NAME"
	dbUserEnvVar = "PL_DATABASE_USER"
	dbPassEnvVar = "PL_DATABASE_PASS"
	dbPortEnvVar = "PL_DATABASE_PORT"
)

var (
	db          *sqlx.DB
	connectOnce = &sync.Once{}
)

func connect() {
	// connect to postgresql
	dbHost := os.Getenv(dbHostEnvVar)
	if dbHost == "" {
		fmt.Println("WARNING: databse host is empty")
	}

	dbName := os.Getenv(dbNameEnvVar)
	if dbName == "" {
		fmt.Println("WARNING: database name is empty")
	}

	dbUser := os.Getenv(dbUserEnvVar)
	if dbUser == "" {
		fmt.Println("WARNING: database user name is empty")
	}

	dbPass := os.Getenv(dbPassEnvVar)
	if dbPass == "" {
		fmt.Println("WARNING: database password is empty")
	}

	dbPort := os.Getenv(dbPortEnvVar)
	if dbPort == "" {
		fmt.Println("WARNING: database port is empty")
	}

	connStr := fmt.Sprintf("sslmode=disable host=%s dbname=%s user=%s password=%s port=%s",
		dbHost,
		dbName,
		dbUser,
		dbPass,
		dbPort,
	)

	// DEBUG
	fmt.Println("DEBUG: db connection string: ", connStr)

	var err error
	if db, err = sqlx.Connect("postgres", connStr); err == nil {
		fmt.Println("connected to postgres!")
		return
	}

	for {
		if db, err = sqlx.Connect("postgres", connStr); err == nil {
			fmt.Println("connected to postgres!")
			return
		}

		fmt.Println("WARNING: database connection failed: ", err)
		time.Sleep(50 * time.Millisecond)
		fmt.Println("retrying...")
	}

}

// GetDB returns a pointer to the connected DB
func GetDB() *sqlx.DB {
	if db == nil {
		connectOnce.Do(connect)
	}

	return db
}

// GetExtFull returns a Postgres-backed sqlx.DB,
// wrapped in a psql.ExtFull interface
func GetExtFull(l *plog.Logger) psql.ExtFull {
	return psql.GetExtFull(l, GetDB())
}

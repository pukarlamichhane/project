package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	connectionString := "postgres://default:EwreR3fGg8ND@ep-orange-credit-11050848.us-east-1.postgres.vercel-storage.com:5432/verceldb"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil
	}

	err = db.Ping()
	if err != nil {
		return nil
	}

	return db
}

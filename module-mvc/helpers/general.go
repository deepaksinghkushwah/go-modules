package helpers

import (
	"database/sql"
)

// GetDB return db object
func GetDB() *sql.DB {
	db, err := sql.Open("mysql", "root:deepak@/test")
	if err != nil {
		panic(err)
	}

	return db
}

func CloseDB(db *sql.DB) {
	db.Close()
}

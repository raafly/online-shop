package config

import (
	"database/sql"
	"time"

	"github.com/raafly/catering/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/catering")

	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 *time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}



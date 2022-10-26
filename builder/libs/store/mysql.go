package store

import (
	"builder/builder/log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var MySQL *sqlx.DB

func initDB() (db *sqlx.DB, err error) {
	dsn := "builder:builder@tcp(127.0.0.1:3306)/builder?charset=utf8mb4&parseTime=True"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, errors.WithMessagef(err, "connect DB failed dsn: %v", dsn)
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)

	return db, nil
}

func init() {
	var err error

	MySQL, err = initDB()
	if err != nil {
		log.L.Fatal(err)
	}
}

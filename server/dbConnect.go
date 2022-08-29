package server

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func DbConnect() error {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "Users",
	}

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Println(err)
		return err
	}

	pingErr := db.Ping()
	// if pingErr != nil {
	// 	log.Println(pingErr)
	// }
	return pingErr
}

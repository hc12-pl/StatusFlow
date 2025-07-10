package database

import (
	"log"
	"go.etcd.io/bbolt"
)

var DB *bbolt.DB

func CheckForDB() {
	DB, err := bbolt.Open("../logs/log.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()
}


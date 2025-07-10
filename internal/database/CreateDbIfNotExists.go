package database

import (
	"log"

	"go.etcd.io/bbolt"
)

var DB *bbolt.DB

func InitDB() {
	var err error
	DB, err = bbolt.Open("log.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Tworzenie bucketa jeśli nie istnieje
	err = DB.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Logs"))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Database initialized")
}

package database

import (
	"fmt"
	"log"
	"time"

	"go.etcd.io/bbolt"
)

func InsertLog(text string) error {
	log.Printf("InsertLog called with %s", text)
	return DB.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("Logs"))
		if bucket == nil {
			return fmt.Errorf("bucket Logs not found")
		}
		key := []byte(time.Now().Format(time.DateTime))
		value := []byte(text)
		return bucket.Put(key, value)
	})
}


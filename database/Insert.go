package database

import (
	"time"

	"go.etcd.io/bbolt"
)

func InsertLog(DB *bbolt.DB, text string) error {
	return DB.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("Logs"))
		if err != nil {
			return err
		}
		key := []byte(time.Now().Format(time.DateTime))
		value := []byte(text)
		return bucket.Put(key, value) // tworzy nowy rekord
	})
}

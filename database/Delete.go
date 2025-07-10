package database

import (
	"go.etcd.io/bbolt"
)

func DeleteLog(DB *bbolt.DB, timestamp string) error {
	return DB.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("Logs"))
		if bucket == nil {
			return nil
		}

		key := []byte(timestamp)
		return bucket.Delete(key)
	})
}
package database

import (
	"fmt"

	"go.etcd.io/bbolt"
)

func DeleteLog(timestamp string) error {
	return DB.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("Logs"))
		if bucket == nil {
			return fmt.Errorf("bucket Logs not found")
		}

		key := []byte(timestamp)
		return bucket.Delete(key)
	})
}

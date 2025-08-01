package database

import (
	"go.etcd.io/bbolt"
	"fmt"
)

func GetAllLogs() (map[string]string, error) {
	logs := make(map[string]string)

	err := DB.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("Logs"))
		if bucket == nil {
			return fmt.Errorf("bucket Logs not found")
		}

		cursor := bucket.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			logs[string(k)] = string(v)
		}

		return nil
	})

	return logs, err
}


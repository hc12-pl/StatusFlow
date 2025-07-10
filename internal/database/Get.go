package database

import (
	"go.etcd.io/bbolt"
	"fmt"
)

func Get(k string) (string, error) {
	var result string
	err := DB.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("Logs"))
		if bucket == nil {
			return fmt.Errorf("bucket Logs not found")
		}

		key := []byte(k)
		value := bucket.Get(key)

		if value == nil {
			return fmt.Errorf("key not found")
		}
		result = string(value)
		return nil
	})
	return result, err
}

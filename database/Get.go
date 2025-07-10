package database

import (
	"go.etcd.io/bbolt"
	"fmt"
)

func Get(DB *bbolt.DB, k string) (string, error) {
	var result string
	err := DB.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("Logs"))
		if bucket == nil {
			return nil
		}

		key := []byte(k)
		value := bucket.Get(key)

		if value == nil {
			return nil
		} else {
			result = fmt.Sprintf("%x", value)
			return nil
		}
	})
	return result, err
}
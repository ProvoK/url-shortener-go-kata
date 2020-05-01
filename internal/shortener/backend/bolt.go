package backend

import (
	"errors"
	"fmt"

	bolt "go.etcd.io/bbolt"
)

type BoltBackend struct {
	DB *bolt.DB
}

type BoltConfig struct {
	Path string
}

const URLBucket string = "URLBucket"

func NewBoltBackend(conf BoltConfig) (*BoltBackend, error) {
	db, err := bolt.Open(conf.Path, 0666, nil)
	if err != nil {
		return nil, err
	}
	return &BoltBackend{DB: db}, nil
}

func (bb *BoltBackend) Store(ID, URL string) error {
	key := []byte(ID)
	err := bb.DB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(URLBucket))
		if err != nil {
			return errors.New(fmt.Sprintf("BoltBackend: Failed to use bucket %s", URLBucket))
		}
		if value := bucket.Get(key); value != nil {
			return errors.New(fmt.Sprintf("BoltBackend: Duplicate key %s", ID))
		}
		err = bucket.Put(key, []byte(URL))
		return err
	})
	return err
}

func (bb *BoltBackend) Resolve(ID string) (string, bool) {
	var URL string
	key := []byte(ID)
	err := bb.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(URLBucket))
		value := bucket.Get(key)
		if value == nil {
			return errors.New(fmt.Sprintf("BoltBackend: Cannot find key %s", ID))
		}
		URL = fmt.Sprintf("%s", value)
		return nil

	})

	if err != nil {
		return "", false
	}
	return URL, true
}

func (bb *BoltBackend) Close() {
	bb.DB.Close()
}

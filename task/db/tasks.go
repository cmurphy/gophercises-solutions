package db

import (
	"encoding/binary"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

func Init() error {
	var err error
	db, err = bolt.Open("tasks.db", 0600, nil)
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Tasks"))
		return err
	})
}

func All() ([]Task, error) {
	result := []Task{}
	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		b.ForEach(func(k, v []byte) error {
			result = append(result, Task{Key: btoi(k), Value: string(v)})
			return nil
		})
		return nil
	}); err != nil {
		return result, err
	}
	return result, nil
}

func Delete(id int) error {
	if err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		return b.Delete(itob(id))
	}); err != nil {
		return err
	}
	return nil
}

func Add(item string) error {
	if err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		id, _ := b.NextSequence()
		return b.Put(itob(int(id)), []byte(item))
	}); err != nil {
		return err
	}
	return nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

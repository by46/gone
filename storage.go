package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
)

type User struct {
	ID   int
	Name string
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func createUser(db *bolt.DB, accountID int) error {
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	u := &User{
		Name: "wendy",
	}

	tx.CreateBucket([]byte(strconv.FormatUint(uint64(accountID), 10)))
	root := tx.Bucket([]byte(strconv.FormatUint(uint64(accountID), 10)))
	bkt, err := root.CreateBucket([]byte("USERS"))
	if err != nil {
		return err
	}
	userID, err := bkt.NextSequence()
	if err != nil {
		return err
	}
	u.ID = int(userID)

	buf, err := json.Marshal(u)
	if err != nil {
		return err
	}
	err = bkt.Put([]byte(strconv.FormatUint(userID, 10)), buf)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func main() {
	db, err := bolt.Open("my.db", 0600, &bolt.Options{
		Timeout: 1 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucket([]byte("nodes"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		bucket.Put([]byte("name"), []byte("Benjamin"))
		bucket.Put([]byte("name1"), []byte("Benjamin1"))
		bucket.Put([]byte("name2"), []byte("Benjamin2"))
		return nil
	})
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("nodes"))
		v := bucket.Get([]byte("name"))
		fmt.Printf("Retrive name value : %s\n", v)
		return nil
	})

	// automatic index
	user := &User{
		Name: "Benjamin",
	}

	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("nodes"))
		id, _ := bucket.NextSequence()
		user.ID = int(id)
		buf, err := json.Marshal(user)
		if err != nil {
			return err
		}
		return bucket.Put(itob(user.ID), buf)
	})
	// iterator bucket

	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("nodes"))

		c := bucket.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("Key: %s, Value: %s\n", k, v)
		}
		return nil
	})

	// scan
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("nodes"))
		cursor := bucket.Cursor()

		prefix := []byte("name")
		for k, v := cursor.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = cursor.Next() {
			fmt.Printf("scan key: %s, value: %s\n", k, v)
		}
		return nil
	})

	// foreach

	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("nodes"))
		bucket.ForEach(func(k, v []byte) error {
			fmt.Printf(" foreach key: %s, value: %s\n", k, v)
			return nil
		})
		return nil
	})

	createUser(db, 2102)
	defer db.Close()
}

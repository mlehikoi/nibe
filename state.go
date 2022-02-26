// Persists user specific settings in Bolt database These settings include
// * User state:
//   - user name
//   - NIBE ID
//   - NIBE secret
//   - Influx DB token (Is this really necessary?)
// * The OAuth token
//   - who uses this information?
//   - why no read function?
package main

import (
	"io/ioutil"

	"github.com/boltdb/bolt"
)

const dnName string = "users.db"

// Save the user state
func Save(user, state, id, secret, influx string) {
	db, err := bolt.Open(dnName, 0600, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte(user))
		if err != nil {
			return err
		}
		if err = b.Put([]byte("state"), []byte(state)); err != nil {
			return err
		}
		if err = b.Put([]byte("id"), []byte(id)); err != nil {
			return err
		}
		if err = b.Put([]byte("secret"), []byte(secret)); err != nil {
			return err
		}
		return b.Put([]byte("influx"), []byte(influx))
	})
}

func SaveToken(user string, token []byte) error {
	return ioutil.WriteFile("token.json", token, 0600)
}

func LoadUser(user string) (string, string) {
	db, err := bolt.Open(dnName, 0600, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var id, secret []byte

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("User1"))
		id = b.Get([]byte("id"))
		secret = b.Get([]byte("secret"))
		return nil
	})
	return string(id), string(secret)
}

// LoadConfig loads the Influx config
func LoadConfig() string {
	db, err := bolt.Open(dnName, 0600, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var influx []byte

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Influx"))
		influx = b.Get([]byte("influx"))
		return nil
	})
	return string(influx)
}

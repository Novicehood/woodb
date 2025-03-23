package main

import (
	"log"
	"time"

	"github.com/rosedblabs/rosedb/v2"
)

// this file shows how to use the Expiry/TTL feature of woodb.
func main() {
	// specify the options
	options := woodb.DefaultOptions
	options.DirPath = "/tmp/rosedb_ttl"

	// open a database
	db, err := woodb.Open(options)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = db.Close()
	}()

	// when you put a key-value pair, you can specify the ttl.
	err = db.PutWithTTL([]byte("name"), []byte("woodb"), time.Second*5)
	if err != nil {
		panic(err)
	}
	// now you can get the ttl of the key.
	ttl, err := db.TTL([]byte("name"))
	if err != nil {
		panic(err)
	}
	println(ttl.String())

	_ = db.Put([]byte("name2"), []byte("rosedb2"))
	//and you can also set the ttl of the key after you put it.
	err = db.Expire([]byte("name2"), time.Second*2)
	if err != nil {
		panic(err)
	}
	ttl, err = db.TTL([]byte("name2"))
	if err != nil {
		log.Println(err)
	}
	println(ttl.String())
}

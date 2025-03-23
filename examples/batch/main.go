package main

import (
	"github.com/rosedblabs/rosedb"
)

// this file shows how to use the batch operations of woodb

func main() {
	// specify the options
	options := woodb.DefaultOptions
	options.DirPath = "/tmp/rosedb_batch"

	// open a database
	db, err := woodb.Open(options)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = db.Close()
	}()

	// create a batch
	batch := db.NewBatch(woodb.DefaultBatchOptions)

	// set a key
	_ = batch.Put([]byte("name"), []byte("woodb"))

	// get a key
	val, _ := batch.Get([]byte("name"))
	println(string(val))

	// delete a key
	_ = batch.Delete([]byte("name"))

	// commit the batch
	_ = batch.Commit()

	// if you want to cancel batch, you must call rollback().
	// _= batch.Rollback()

	// once a batch is committed, it can't be used again
	// _ = batch.Put([]byte("name1"), []byte("rosedb1")) // don't do this!!!
}

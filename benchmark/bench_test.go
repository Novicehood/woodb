package benchmark

import (
	"errors"
	"math/rand"
	"os"
	"testing"

	"github.com/rosedblabs/rosedb/v2"
	"github.com/rosedblabs/rosedb/v2/utils"
	"github.com/stretchr/testify/assert"
)

var db *woodb.DB

func openDB() func() {
	options := woodb.DefaultOptions
	options.DirPath = "/tmp/rosedbv2"

	var err error
	db, err = woodb.Open(options)
	if err != nil {
		panic(err)
	}

	return func() {
		_ = db.Close()
		_ = os.RemoveAll(options.DirPath)
	}
}

func BenchmarkPutGet(b *testing.B) {
	closer := openDB()
	defer closer()

	b.Run("put", benchmarkPut)
	b.Run("get", bencharkGet)
}

func BenchmarkBatchPutGet(b *testing.B) {
	closer := openDB()
	defer closer()

	b.Run("batchPut", benchmarkBatchPut)
	b.Run("batchGet", benchmarkBatchGet)
}

func benchmarkPut(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		err := db.Put(utils.GetTestKey(i), utils.RandomValue(1024))
		assert.Nil(b, err)
	}
}

func benchmarkBatchPut(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	batch := db.NewBatch(woodb.DefaultBatchOptions)
	defer batch.Commit()
	for i := 0; i < b.N; i++ {
		err := batch.Put(utils.GetTestKey(i), utils.RandomValue(1024))
		assert.Nil(b, err)
	}
}

func benchmarkBatchGet(b *testing.B) {
	for i := 0; i < 10000; i++ {
		err := db.Put(utils.GetTestKey(i), utils.RandomValue(1024))
		assert.Nil(b, err)
	}

	b.ResetTimer()
	b.ReportAllocs()
	batch := db.NewBatch(woodb.DefaultBatchOptions)
	defer batch.Commit()
	for i := 0; i < b.N; i++ {
		_, err := batch.Get(utils.GetTestKey(rand.Int()))
		if err != nil && !errors.Is(err, woodb.ErrKeyNotFound) {
			b.Fatal(err)
		}
	}
}

func bencharkGet(b *testing.B) {
	for i := 0; i < 10000; i++ {
		err := db.Put(utils.GetTestKey(i), utils.RandomValue(1024))
		assert.Nil(b, err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := db.Get(utils.GetTestKey(rand.Int()))
		if err != nil && !errors.Is(err, woodb.ErrKeyNotFound) {
			b.Fatal(err)
		}
	}
}

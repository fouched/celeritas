package cache

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/dgraph-io/badger/v4"
	"github.com/gomodule/redigo/redis"
	"log"
	"os"
	"testing"
	"time"
)

var testRedisCache RedisCache
var testBadgerCache BadgerCache

func TestMain(m *testing.M) {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()

	pool := redis.Pool{
		MaxIdle:     10,
		MaxActive:   100,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", s.Addr())
		},
	}

	testRedisCache.Conn = &pool
	testRedisCache.Prefix = "test-celeritas"

	defer testRedisCache.Conn.Close()

	// clear and create badger database
	_ = os.RemoveAll("./testdata/tmp/badger")
	if _, err := os.Stat("./testdata/tmp"); os.IsNotExist(err) {
		err := os.Mkdir("./testdata/tmp", 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = os.Mkdir("./testdata/tmp/badger", 0755)
	if err != nil {
		log.Fatal(err)
	}

	db, _ := badger.Open(badger.DefaultOptions("./testdata/tmp/badger"))
	testBadgerCache.Conn = db

	os.Exit(m.Run())
}

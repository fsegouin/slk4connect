package db

import (
	"github.com/fsegouin/slk4connect/Godeps/_workspace/src/github.com/fzzy/radix/redis"
	"log"
	"sync"
)

var instance *redis.Client
var once sync.Once

func GetInstance() *redis.Client {
	once.Do(func() {
		client, err := redis.Dial("tcp", "localhost:6379")

		if err != nil {
			log.Fatal(err)
		}

		instance = client
	})

	return instance
}

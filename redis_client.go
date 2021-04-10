package main

import (
	"context"
	"encoding/json"
	. "github.com/esakat/observe_my_hatebu/data"
	"github.com/go-redis/redis/v8"
	"log"
)

var ctx = context.Background()
var rdbHatebuEntry *redis.Client

func createRedisClient() {
	rdbHatebuEntry = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: "",
		DB:       config.RedisDB,
	})
}

func HasEid(eid string) bool {
	result, err := rdbHatebuEntry.Exists(ctx, eid).Result()
	if err != nil {
		panic(err)
	}
	if result > 0 {
		return true
	} else {
		return false
	}
}

func InsertMyHatebuEntry(entry *MyEntry) {
	entryJsonString, _ := json.Marshal(entry)
	var m = make(map[string]interface{})
	if err := json.Unmarshal(entryJsonString, &m); err != nil {
		log.Println(err)
		return
	}
	_, err := rdbHatebuEntry.HSet(ctx, entry.EntryID, m).Result()
	if err != nil {
		panic(err)
	}
}

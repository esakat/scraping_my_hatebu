package main

import (
	"github.com/kelseyhightower/envconfig"
	"log"
	"time"
)

type Config struct {
	HatebuUserName      string        `required:"true" split_words:"true"`
	UpdateDuration      time.Duration `default:"1m" split_words:"true"`
	BotToken            string        `required:"true" split_words:"true"`
	ChannelID           string        `required:"true" split_words:"true"`
	MentionUser         string        `required:"true" split_words:"true"`
	EntryCollectionName string        `required:"true" split_words:"true"`
	ProjectID           string        `required:"true" split_words:"true"`
}

var config Config

func init() {
	err := envconfig.Process("app", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	createFirestoreClient()
}

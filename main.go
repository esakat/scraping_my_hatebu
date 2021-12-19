package main

import (
	. "github.com/esakat/observe_my_hatebu/data"
	"log"
	"time"
)

func main() {

	// Get New Entry
	myBookmarks := scrape(config.HatebuUserName)
	for _, bookmark := range myBookmarks {
		// notify to slack
		timestamp, err := notifyToChannel(bookmark, config.MentionUser)
		if err != nil {
			log.Println(err)
			continue
		}

		// push to firestore
		AddMyHatebuEntry(&MyEntry{
			EntryID:         bookmark.EntryID,
			URL:             bookmark.URL,
			ThreadTimestamp: timestamp,
			UpdateTimestamp: bookmark.Timestamp,
		})

		// for slack waiting time
		time.Sleep(time.Second)
	}

}

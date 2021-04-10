package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func scrape(username string) (myBookmarks []MyBookmark) {

	// Create URL
	url := fmt.Sprintf("https://b.hatena.ne.jp/%s/bookmark", username)

	// Get Request
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		return
	}

	// Read Body
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	// Check Encoding
	det := chardet.NewTextDetector()
	detResult, err := det.DetectBest(buf)
	if err != nil {
		return
	}

	// Translate Encoding
	bReader := bytes.NewReader(buf)
	reader, err := charset.NewReaderLabel(detResult.Charset, bReader)
	if err != nil {
		return
	}

	// Parse HTML
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return
	}

	// Search bookmark-items
	doc.Find(".bookmark-item").Each(func(i int, s *goquery.Selection) {
		// GET Entry URL
		href, _ := s.Find(".centerarticle-entry-title").Find("a").Attr("href")

		// GET Entry Ttile
		title := s.Find(".centerarticle-entry-title").Find("a").Text()

		// GET Entry Own Comment
		comment := s.Find(".js-comment").Text()
		// GET Entry Own Tags
		var tags []string
		s.Find(".centerarticle-reaction-tags").Each(func(j int, tag *goquery.Selection) {
			tags = append(tags, strings.TrimSpace(tag.Text()))
		})

		category := s.Find("a[data-gtm-click-label='user-bookmark-category']")
		categoryId, _ := category.Attr("href")
		categoryId = categoryId[10:]
		categoryName := category.Text()

		// GET Entry ID
		eid, _ := s.Find(".centerarticle-reaction").Attr("id")
		eid = strings.Replace(eid, "bookmark-", "", 1)

		// GET Entry Timestamp
		timestamp := s.Find(".centerarticle-reaction-timestamp").Text()

		// check redis set, if it is already exist, do nothing.
		if !HasEid(eid) && isToday(timestamp) {
			myBookmarks = append(myBookmarks, MyBookmark{
				EntryID:   eid,
				URL:       href,
				Title:     title,
				Timestamp: time.Now().Format("2006/01/02 15:04"),
				Comment:   comment,
				Tags:      tags,
				Category: Category{
					Name: categoryName,
					ID:   categoryId,
				},
			})
			log.Printf("eid: %s has been bookmarked.\n", eid)
		}
	})

	return
}

func isToday(timestamp string) bool {
	format := "2006/01/02"
	now := time.Now().Format(format)
	return timestamp == now
}

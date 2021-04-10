package main

type MyBookmark struct {
	EntryID   string   `json:"eid"`
	URL       string   `json:"url"`
	Title     string   `json:"title"`
	Timestamp string   `json:"thread_timestamp"`
	Comment   string   `json:"comment"`
	Tags      []string `json:"tags"`
	Category  Category `json:"category"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

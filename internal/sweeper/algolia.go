package sweeper

import (
	"io"
	"log"
	"time"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

type Config struct {
	EventIndex *search.Index
}

type Record struct {
	ObjectID      string `json:"objectID"`
	DateTimestamp int64  `json:"dateTimestamp"`
	ID            int    `json:"id"`
	Title         string `json:"title"`
}

func (c Config) SweepAlgoliaRecords() {
	it, err := c.EventIndex.BrowseObjects()
	var record Record

	if err != nil {
		log.Fatalf("Error browsing Algolia records %s", err)
	}
	deletedRecordsCount := 0

	for {
		_, err := it.Next(&record)

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Error with event %s, error: %s", record.Title, err)
			continue
		}

		eventDate := time.UnixMilli(record.DateTimestamp)

		if isMoreThanOneMonthAgo(eventDate) {
			c.EventIndex.DeleteObject(record.ObjectID)
			log.Printf("Record %s deleted", record.ObjectID)
			deletedRecordsCount++
		}
	}

	log.Printf("%d deleted records", deletedRecordsCount)
}

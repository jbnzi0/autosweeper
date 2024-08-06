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
		log.Fatalf("Error browsing Algolia records %v", err)
	}
	deletedRecordsCount := 0

	for {
		_, err := it.Next(&record)

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Error with event %s, error: %v", record.Title, err)
			continue
		}

		eventDate := time.UnixMilli(record.DateTimestamp)

		if isMoreThanOneMonthAgo(eventDate) {
			_, err := c.EventIndex.DeleteObject(record.ObjectID)

			if err != nil {
				log.Printf("Error deleting record %s, error: %v", record.ObjectID, err)
				continue
			}

			log.Printf("Record %s with ID %s deleted", record.Title, record.ObjectID)
			deletedRecordsCount++
		}
	}

	log.Printf("%d deleted records", deletedRecordsCount)
}

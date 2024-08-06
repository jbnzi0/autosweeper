package sweeper

import (
	"io"
	"log"

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

	for {
		_, err := it.Next(&record)

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Error with event %s, error: %s", record.Title, err)
			continue
		}

		log.Println(record.Title, record.DateTimestamp)

	}
}

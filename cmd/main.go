package main

import (
	"flag"
	"log"
	"os"

	"github.com/algobroom/internal/sweeper"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/joho/godotenv"
)

func main() {
	withEnv := flag.Bool("with-env", false, "Parse env file")
	flag.Parse()

	if *withEnv {
		err := godotenv.Load("./.env")
		if err != nil {
			log.Println("Env file not found or error loading .env file")
		}
	}

	client := search.NewClient(os.Getenv("ALGOLIA_APP_ID"), os.Getenv("ALGOLIA_API_KEY"))
	index := client.InitIndex(os.Getenv("ALGOLIA_INDEX"))

	c := sweeper.Config{
		EventIndex: index,
	}

	c.SweepAlgoliaRecords()
}

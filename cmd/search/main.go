package main

import (
	"fmt"
	"os"

	"github.com/meilisearch/meilisearch-go"
)

func main() {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://127.0.0.1:7700",
		APIKey: "masterKey",
	})

	searchRes, err := client.Index("movies").Search("philoudelphia",
		&meilisearch.SearchRequest{
			ShowMatchesPosition: true,
			Limit:               20,
		})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(searchRes.TotalPages)
	for _, v := range searchRes.Hits {
		fmt.Println(v)
	}
}

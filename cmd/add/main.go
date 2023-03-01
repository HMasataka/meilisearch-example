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
	// An index is where the documents are stored.
	index := client.Index("movies")

	// If the index 'movies' does not exist, Meilisearch creates it when you first add the documents.
	documents := []map[string]interface{}{
		{"id": 1, "title": "Carol", "genres": []string{"Romance", "Drama"}},
		{"id": 2, "title": "Wonder Woman", "genres": []string{"Action", "Adventure"}},
		{"id": 3, "title": "Life of Pi", "genres": []string{"Adventure", "Drama"}},
		{"id": 4, "title": "Mad Max: Fury Road", "genres": []string{"Adventure", "Science Fiction"}},
		{"id": 5, "title": "Moana", "genres": []string{"Fantasy", "Action"}},
		{"id": 6, "title": "Philadelphia", "genres": []string{"Drama"}},
		{"id": 7, "title": "Carol", "genres": []string{"Romance", "Drama"}},
		{"id": 8, "title": "Wonder Woman", "genres": []string{"Action", "Adventure"}},
		{"id": 9, "title": "Life of Pi", "genres": []string{"Adventure", "Drama"}},
		{"id": 10, "title": "Mad Max: Fury Road", "genres": []string{"Adventure", "Science Fiction"}},
		{"id": 11, "title": "Moana", "genres": []string{"Fantasy", "Action"}},
		{"id": 11, "title": "Philadelphia", "genres": []string{"Drama"}},
		{"id": 12, "title": "Carol", "genres": []string{"Romance", "Drama"}},
		{"id": 13, "title": "Wonder Woman", "genres": []string{"Action", "Adventure"}},
		{"id": 14, "title": "Life of Pi", "genres": []string{"Adventure", "Drama"}},
		{"id": 15, "title": "Mad Max: Fury Road", "genres": []string{"Adventure", "Science Fiction"}},
		{"id": 16, "title": "Moana", "genres": []string{"Fantasy", "Action"}},
		{"id": 17, "title": "Philadelphia", "genres": []string{"Drama"}},
		{"id": 19, "title": "Carol", "genres": []string{"Romance", "Drama"}},
		{"id": 20, "title": "Wonder Woman", "genres": []string{"Action", "Adventure"}},
		{"id": 21, "title": "Life of Pi", "genres": []string{"Adventure", "Drama"}},
		{"id": 22, "title": "Mad Max: Fury Road", "genres": []string{"Adventure", "Science Fiction"}},
		{"id": 23, "title": "Moana", "genres": []string{"Fantasy", "Action"}},
		{"id": 24, "title": "Philadelphia", "genres": []string{"Drama"}},
		{"id": 25, "title": "Carol", "genres": []string{"Romance", "Drama"}},
		{"id": 25, "title": "Wonder Woman", "genres": []string{"Action", "Adventure"}},
		{"id": 26, "title": "Life of Pi", "genres": []string{"Adventure", "Drama"}},
		{"id": 27, "title": "Mad Max: Fury Road", "genres": []string{"Adventure", "Science Fiction"}},
		{"id": 28, "title": "Moana", "genres": []string{"Fantasy", "Action"}},
		{"id": 29, "title": "Philadelphia", "genres": []string{"Drama"}},
		{"id": 30, "title": "Philadelphia", "genres": []string{"Drama"}},
		{"id": 31, "title": "Philadelphia", "genres": []string{"Drama"}},
		{"id": 32, "title": "Philadelphia", "genres": []string{"Drama"}},
		{"id": 33, "title": "Philadelphia", "genres": []string{"Drama"}},
		{"id": 34, "title": "Philadelphia", "genres": []string{"Drama"}},
		{"id": 35, "title": "Philadelphia", "genres": []string{"Drama"}},
		{"id": 36, "title": "Philadelphia", "genres": []string{"Drama"}},
		{"id": 37, "title": "Philadelphia", "genres": []string{"Drama"}},
		{"id": 38, "title": "Philadelphia", "genres": []string{"Drama"}},
		{"id": 39, "title": "Philadelphia", "genres": []string{"Drama"}},
	}

	task, err := index.AddDocuments(documents)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(task.TaskUID)
}

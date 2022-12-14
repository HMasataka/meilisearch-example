# meilisearch

## Run with Docker

```bash
docker run -it --rm \
    -p 7700:7700 \
    -e MEILI_MASTER_KEY='MASTER_KEY'\
    -v $(pwd)/meili_data:/meili_data \
    getmeili/meilisearch:v0.30 \
    meilisearch --env="development"
```

```bash
curl \
   -X POST 'http://localhost:7700/indexes/movies/documents?primaryKey=id' \
   -H 'Content-Type: application/json' \
   -H 'Authorization: Bearer MASTER_KEY' \
   --data-binary @movies.json
```

```bash
curl \
    -X POST 'http://localhost:7700/indexes/movies/search' \
    -H 'Content-Type: application/json' \
    -H 'Authorization: Bearer MASTER_KEY' \
    --data-binary '{ "q": "botman" }'
```

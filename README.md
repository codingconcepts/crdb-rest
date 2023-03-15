# crdb-rest
A barebones example of accessing CockroachDB via REST

### Dependencies

* [Go](https://go.dev)

### Running locally

Start a cluster
```
$ cockroach start-single-node \
    --listen-addr=localhost:26257 \
    --http-addr=localhost:8080 \
    --insecure
```

Create a table
```
$ cockroach sql --insecure < create.sql
```

Run the server
```
$ go run main.go
```

### Requests

Create a todo
```
$ curl -X POST "http://localhost:3000/v1/todos" \
    -H "Content-Type: application/json" \
    -d '{
        "title": "todo e"
    }'
```

Fetch all todos
```
$ curl "http://localhost:3000/v1/todos?per_page=1"
```

Fetch a todo
```
$ curl "http://localhost:3000/v1/todos/03036918-ba3d-4264-84ed-94e0c6d7433e"
```

Delete a todo
```
$ curl -X DELETE "http://localhost:3000/v1/todos/cbdcbe9a-ce24-41ce-853a-a0f3c7ec7ec0"
```
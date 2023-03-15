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

```

Fetch all todos
```

```

Fetch a todo
```

```

Delete a todo
```

```
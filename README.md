## An API in GO for concurrent requests employing the Hexagonal architecture

A model for a concurrent Go API facilitating requests to an external API.

### Requirements

- Go v1.18+
- MySQL (see below).


#### MySQL & Docker

We've added a
`docker-compose.yaml` file with a MySQL container already set up.

To run it, just execute:

```sh
docker-compose up -d 
```

You can also use your own MySQL instance. Note that those applications
expects a MySQL instance to be available on `localhost:3308`,
identified by `go-api:go-api` and with a `go-api` database.

To set up your database, you can execute the `schema.sql` file
present on the `sql` directory. It's automatically loaded if
you use the provided `docker-compose.yaml` file.

#### Tests

To execute all tests, just run:

```sh
go test ./requests/... 
```
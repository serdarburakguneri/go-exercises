# simple-postgres-client

The purpose of this exercise is to implement repository pattern and test it with a running postgres instance.

## How to run

Start a docker instance of your favourite Postgres version.
Modify the connection string in PostgresClient.OpenConnection function (host, port, username, password and dbname:)
with what you have running.

In the main folder "simple-postgres-client"

```bash
go build .
```

```bash
go run .
```

to create a customer table with a record, you can uncomment populateDB method in main.go 
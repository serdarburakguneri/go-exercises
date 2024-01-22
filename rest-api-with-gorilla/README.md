# rest api with gorilla

This is an example REST API implemented in GO.

Instead of a single "main.go" file, I am creating different files and grouping related files under
some packages for a more structured project..

## How to run

In the main folder "rest-api-with-gorilla"

```bash
go build .
```

```bash
go run .
```

## testing the endpoints

* GET (all)

```bash
curl http://localhost:8080/scooter 
```

* GET (by id)

```bash
curl http://localhost:8080/scooter/56ff36c1-b913-45f6-8fc3-d238569b6995
```


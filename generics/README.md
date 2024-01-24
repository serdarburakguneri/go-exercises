# generics

With this exercise, I am discovering generics and predicates.

My Filter function

```
func Filter[T any](collection []T, predicate func(T) bool) []T {
```

receives a collection of any type, runs a predicate on its each element and returns a collection
with the elements that can pass the predicate.


## How to run
n generics folder:

```bash
go build main.go
```

```bash
go run main.go
```
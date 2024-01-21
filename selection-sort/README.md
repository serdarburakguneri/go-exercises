# bubble-sort

Here is a selection sort algorithm implementation in Go.

Within the second loop, the algorithm finds the index of the min element and swaps if it's smaller
than the current element in the first loop.

Here we can see the shape of the collection in each step of the algorithm:


[1 5 6 3 99 44 55 22]

[1 3 6 5 99 44 55 22]

[1 3 5 6 99 44 55 22]

[1 3 5 6 99 44 55 22]

[1 3 5 6 22 44 55 99]

[1 3 5 6 22 44 55 99]

[1 3 5 6 22 44 55 99]

## How to run

```bash
go build main.go
```

```bash
go run main.go
```
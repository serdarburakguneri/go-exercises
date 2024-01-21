# bubble-sort

Here is a bubble sort algorithm implementation in Go.

Running 2 for loops, the algorithm compares each element of the array 
with the next one and swaps them if the next one is smaller.

Here we can see the shape of the collection in each step of the algorithm:


[1 5 6 3 99 44 55 22]

[1 5 3 6 44 55 22 99]

[1 3 5 6 44 22 55 99]

[1 3 5 6 22 44 55 99]

[1 3 5 6 22 44 55 99]

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
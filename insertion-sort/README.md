# bubble-sort

An insertion sort algorithm implementation in Go.

The goal is the move larger elements to the right in every iteration so that 
the element at current index can be in correct order.

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
# dependency-injection

Here I tested an interface definition and implementing it by two structs.
(As a java guy, I am uncomfortable not seeing an 'implements' keyword :) ) 

Next step was defining a struct that has an object implementing the interface.
With a constructor, I let the clients of this struct to decide which implementation
of the interface to use.

This is basically a constructor injection. With an IoC container, we can create miracles with 
ctor injection.


## How to run

In the main folder "dependency-injection"

```bash
go build .
```

```bash
go run .
```



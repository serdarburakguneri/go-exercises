# dependency-injection-with-wire

In this module, I tested google wire library for dependency injection.

The module contains a ScooterController struct to simulate a rest controller. It works with a
ScooterService as an interface.

I have two structs implementing ScooterService interface and within wire.go file, I decide which
implementation to be used with the ScooterController.

After running "wire" command, we can see the generated "wire_gen.go" file has an init method with desired
implementation of ScooterService.

So, it brings a structure for dependency injection. Having this configuration in only one place may help
to use the same implementer everywhere in the project. Respecting this style, we can have a kind of IOC mechanism.

Alternatively, we may think of having another layer, like defining application scoped java beans, and we can configure
some functions to return desired types by deciding which implementations to inject. And by consuming these functions only 
for getting instances, we can have a similar effect.

## How to run

In the main folder "dependency-injection-with-wire":

* Uncomment the "Initialize" code in wire.go.

* Remove wire_gen.go file.

```bash
wire
```

```bash
go build .
```

```bash
go run .
```

and then you will see the generated wire_gen.go file.
# paytabs_task

This component is developed using go-chi https://github.com/go-chi/chi router and oapi-codegen https://github.com/deepmap/oapi-codegen to simulate money transfer between accounts and using go mod for dependancies management.

# Use api/api.yml file in the swagger ui to see the description of the client interface



## Setup of the component

Must have golang installed version >= 12.0.0

make file consists of 4 steps: generate, test, build, run
you can run all of them 

```bash
make all
```



## Test consume accounts data and make transferes by run

```bash
make test
```
or

go test -v ./tests


# Start the http server on port 8080:

```bash
make run
```

# Build and run docker image

```bash
docker build --tag transfer-service .
```

```bash
docker run -p 9090:9090 transfer-service
```

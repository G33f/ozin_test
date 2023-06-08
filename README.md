## Description

This is a gRPC service that has two endpoints
    The first **MakeURLShorter** takes a **CrateRequest** as an argument, and it returns a **CrateResponse**. Both types have a **url** field, which is of type string 
    The service takes the url from the **CrateRequest** and generates a new, unique url 
(hash) based on it. It saves it into a DB which type can be chosen, but more about that later. After that it returns 
the new url in the **CrateResponse** body.
    The second endpoint, **GetOriginalURL**, does the opposite. It takes the url, in the body of the **GetRequest** it generated 
looks for it in the repository, then returns it in the body of **GetResponse**.

You can choose the type of repository, you can go **inMemory** or in **PostgreSQL**. This can be done by using flag 
-storage_type, which can be specified at startup (``-storage_type=PostgreSQL`` or ``-storage_type=InMemory``). 
If you start the service without this flag, it will work in **InMemory** mode.

## Startup guide

It is very simple, you need to have ``docker`` (**https://docs.docker.com/desktop/install/linux-install/**) 
and ``docker-compose`` (**https://docs.docker.com/engine/install/ubuntu/**).

Next, simply call:
```shell
make compose-up-in-memory
```
or
```shell
make compose-up-postgresql
```
or you can do it yourself by running
```shell
go run cmd/api/main.go
```
```shell
go run cmd/api/main.go -storage_type=PostgreSQL
```
```shell
go run cmd/api/main.go -storage_type=InMemory
```


## service check

There are different options for checking the service. Here are a couple of examples.

>Use the **gRPCurl** utility: https://pkg.go.dev/github.com/fullstorydev/grpcurl#section-readme
>
>Use the api platform **postman** or **insomnia**
>
>Use your own api that implements the client part of gRPC service
>
>The proto file itself can be found in the `proto` directory

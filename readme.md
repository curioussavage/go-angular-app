# About

This is an example of a Go server that provides an HTTP API and swagger docs.


# Build

The angular client for the API has been checked in so it should not be necessary to
rebuild it. If necessary that can be done by installing [swagger-codegen](https://swagger.io/tools/swagger-codegen/)

```
swagger-codegen generate -i /path/to/project/go-angular-app/server/docs/swagger.json -l typescript-angular -o /path/to/project/go-angular-app/projects/api-client
```

If the go API is changed you also need to run the [swag](https://github.com/swaggo/swag) to the spec.


```
swag init -g cmd/server/main.go
```

# Run

to run the server:

```
cd server/
go run cmd/server/main.go

```

to run the client:

```
cd client/
npm install
npm run start
```

# Test

server

```
cd server/
go test ./models ./controllers
```

client

```
npm run test
```

# notes

The following features were not implemented comprehensively to save time.

* test suite for both server and client
* client side validation of forms.
* server side validation of requests

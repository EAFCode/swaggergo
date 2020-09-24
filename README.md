# Generating swagger documentation with Golang - Product API

Swagger documentation is generated from the code using [go-swagger](https://github.com/go-swagger/go-swagger) package.

1) step: Get dependency witch contains a golang implementation of Swagger 2.0 
```
go get -u github.com/go-swagger/go-swagger/cmd/swagger
```
2) step: Adding [metadata](https://goswagger.io/use/spec/meta.html) in your code.

3) step: Generate the file swagger.yaml using the following command:
```
swagger generate spec -o ./swagger.yaml --scan-models
```
4) step: If you want swagger-ui import the package [go-openapi](https://github.com/go-openapi/runtime/tree/master/middleware) in your code
```
import "github.com/go-openapi/runtime/middleware"
```
5) step: running your code:
```
go run main.go
```
Swagger documentation can be viewed using the ReDoc UI in your browser at:
```
swagger-ui:   http://localhost:8080/docs.
swagger-spec: http://localhost:8080/swagger.yaml
```

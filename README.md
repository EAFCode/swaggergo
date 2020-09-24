# Generating swagger documentation with Golang - Product API

Swagger documentation is generated from the code using [go-swagger](https://github.com/go-swagger/go-swagger) package.

1) step: Get dependency witch contains a golang implementation of Swagger 2.0 
```
go get -u github.com/go-swagger/go-swagger/cmd/swagger
```
2) step: Adding metadata(route, param or model) in your code.
```
https://goswagger.io/use/spec/model.html
```
3) step: Generate the file swagger.yaml using the following command:
```
swagger generate spec -o ./swagger.yaml --scan-models
```
4) step: If you want swagger-ui add package [go-openapi](https://github.com/go-openapi/runtime/tree/master/middleware) your code
```
github.com/go-openapi/runtime/middleware"
```
5) step: running your code and put this link in browser:
```
go run main.go
```
Swagger documentation can be viewed using the ReDoc UI in your browser at:
```
swagger-ui:   http://localhost:8080/docs.
swagger-spec: http://localhost:8080/swagger.yaml
```

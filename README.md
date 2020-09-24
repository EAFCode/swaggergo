# Generating swagger documentation with Golang - Product API

Swagger documentation is generated from the code using go-swagger package.

1) step: Get dependency in your  GOPATH
```
go get -u github.com/go-swagger/go-swagger/cmd/swagger
```
2) step: Adding metadata(route, param or model) in your code.
```
https://goswagger.io/use/spec/model.html
```
3) step: Generete the file swagger.yaml using the following command:
```
swagger generate spec -o ./swagger.yaml --scan-models
```
4) step: If you want swagger-ui add package go-openapi in your code
```
github.com/go-openapi/runtime/middleware"
```
see https://github.com/go-openapi/runtime/tree/master/middleware

5) step: running your code and put this link in your browser:
```
go run main.go
```
Swagger documentation can be viewed using the ReDoc UI in your browser at:
```
swagger-ui:   http://localhost:8080/docs.
swagger-spec: http://localhost:8080/swagger.yaml
```

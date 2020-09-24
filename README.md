# Generate swagger documentation with Golang

Swagger documentation is generated from the code annotations inside the source using go-swagger.

1) step: Get dependency in your directory GOPATH

go get -u github.com/go-swagger/go-swagger/cmd/swagger

2) step: Adding meta data in your route, param or model.

https://goswagger.io/use/spec/model.html

3) step: Genereted your spec( swagger.yaml) using this command in your cmd:

swagger generate spec -o ./swagger.yaml --scan-models

4) step: If you want swagger-ui add package  go-openapi in your code

github.com/go-openapi/runtime/middleware"

see https://github.com/go-openapi/runtime/tree/master/middleware

5) step: running your code and put this link in your browser:
go run main.go

Swagger documentation can be viewed using the ReDoc UI in your browser at:

swagger-ui:   http://localhost:8080/docs.
swagger-spec: http://localhost:8080/swagger.yaml

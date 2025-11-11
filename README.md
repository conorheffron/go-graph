# go-graph

[![Go](https://github.com/conorheffron/go-graph/actions/workflows/go.yml/badge.svg)](https://github.com/conorheffron/go-graph/actions/workflows/go.yml)

## Overview
 - Track to do items with Go & GraphQL.

 ### Quick start build steps
 ```shell
go install github.com/99designs/gqlgen@latest

go mod init go-graph

go get github.com/99designs/gqlgen@v0.17.81

go run github.com/99designs/gqlgen init

go mod tidy

go build .

./go-graph
```
### Should see the following in console:
```shell
2025/10/25 17:37:16 connect to http://localhost:8080/ for GraphQL playground
```

### Start Server (without Build Steps)
```shell
go run ./server.go  
```

### Go to GraphiQL UI
 - http://localhost:8080/

### Sample Mutations & Query (GraphQL format)
```graphql
query myToDos {
  todos {
    id
    text
    done
  }
}

mutation addToDo {
  createTodo(text: "Learn GraphQL") {
    id
    text
    done
  }
}

mutation addToDo2 {
  createTodo(text: "Shopping") {
    id
    text
    done
  }
}

mutation addToDo3 {
  createTodo(text: "Exercise") {
    id
    text
    done
  }
}
```

### Sample Mutations & Query (curl format)

#### Muation (Add 'Learn GraphQL' to do)
```shell
curl -X POST -H "Content-Type: application/json" \
-d '{"query": "mutation { createTodo(text: \"Learn GraphQL\") { id text done } }"}' \
http://localhost:8080/query
```
#### Response
```json
{"data":{"createTodo":{"id":"8053753236064475949","text":"Learn GraphQL","done":false}}}
```        

#### Muation (Add 'Shopping' to do)                                               
```shell
curl -X POST -H "Content-Type: application/json" \
-d '{"query": "mutation { createTodo(text: \"Shopping\") { id text done } }"}' \
http://localhost:8080/query
```
#### Response
```json
{"data":{"createTodo":{"id":"3260314713893700176","text":"Shopping","done":false}}}
```

#### Query (Get TODOs list)  
```shell
curl -X POST -H "Content-Type: application/json" \
-d '{"query": "query { todos { id text done } }"}' \                            
http://localhost:8080/query
```
#### Response
```json
{"data":{"todos":[{"id":"8053753236064475949","text":"Learn GraphQL","done":false},{"id":"3260314713893700176","text":"Shopping","done":false}]}}
```

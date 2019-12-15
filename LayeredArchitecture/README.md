# Layered Architecture

```
server/
├── Dockerfile
├── application
│   └── todo.go
├── cmd
│   └── server
│       └── main.go
├── domain
│   ├── model
│   │   └── todo.go
│   └── repository
│       └── todo.go
├── go.mod
├── go.sum
├── httputil
│   └── error.go
├── infrastructure
│   └── persistence
│       └── todo.go
└── presentation
    ├── controller
    │   └── todo.go
    └── view
        └── json.go
```
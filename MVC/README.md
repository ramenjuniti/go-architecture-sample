# MVCパターン

```
server
├── Dockerfile
├── cmd
│   └── server
│       └── main.go
├── controller
│   └── todo.go
├── go.mod
├── go.sum
├── httputil
│   └── error.go
├── model
│   └── todo.go
└── view
    └── json.go
```

## Model

ViewやControllerに置くべきでないロジックを実装

## View

サーバーの出力に関するロジックを実装

## Controller

ModelとViewを操作して以下のような任意の処理を実行

* HTTPのリクエストを処理してModelとViewに渡す
* DBのトランザクション制御

など
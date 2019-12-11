FROM golang:alpine

WORKDIR /go/src/server
COPY . .

RUN go build -o server ./cmd/server/main.go

CMD ["./server"]
FROM golang:1.7
WORKDIR /go/src/github.com/therealplato/muid-demo
ADD *.go ./
CMD go run main.go

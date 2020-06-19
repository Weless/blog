#！ /bin/bash
#set GOOS=linux
#set GOARCH=amd64
GOOS=linux GOARCH=amd64 go build -o build/gin main.go routes.go
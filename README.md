#go-upload

Simple File Upload server written in Go while learning how to use Go :) 

The server uses Basic Auth:
Username: upload
Password are the numbers of the current date - Example: 20151118 - November 18 2015

Dependencies:
sqlite3

First install Sqlite3 Go driver:
```
go get github.com/mattn/go-sqlite3
```

Compile & Run using:
```
go build && ./go-upload
```

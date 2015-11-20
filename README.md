#go-upload

Simple File Upload server written in Go while learning how to use Go :) 

The server uses Basic Auth:

Username: upload

Password are the numbers of the current date - Example: 20151118 - November 18 2015

#Runtime dependencies:
sqlite3

#Compiling:
1. Install the Go programming language for your platform https://golang.org/
2. Install Sqlite for your platform https://www.sqlite.org/
3. Install Sqlite3 Go driver:
```
go get github.com/mattn/go-sqlite3
```
4. Compile & Run using:
```
go build && ./go-upload
```

# go-upload


Simple File Upload server written in Go while learning how to use Go :) 


The server uses Basic Auth:


Username: upload

Password: pass123


## Runtime dependencies:

* sqlite3

## Compiling:

* Install the Go programming language for your platform https://golang.org/
* Install Sqlite for your platform https://www.sqlite.org/
* Install Sqlite3 Go driver:

```
go get github.com/mattn/go-sqlite3
```

* Compile & Run using:

```
go build && ./go-upload
```

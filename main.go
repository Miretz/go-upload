package main

import (
	"fmt"
	"net/http"
	"io"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	"text/template"
	"time"
	"crypto/md5"
	"html"
	"log"
)

const port int = 3000
const uploadDir string = "./files/"
const username = "upload"
const pass = "pass123"

type File struct {
	Id          string
	Name        string
	Path        string
	Comment     string
	Create_date string
}

func init() {
	InitializeDB()
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if IsUnauthorized(w, r) {
		return
	}

	data := ListFilesInDB()
	t, _ := template.ParseFiles("./tmpl/index.gtpl")
	t.Execute(w, data)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if IsUnauthorized(w, r) {
		return
	}

	//display upload form
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("./tmpl/upload.gtpl")
		t.Execute(w, token)
		return
	}

	//do not allow other than POST
	if r.Method != "POST" {
		return
	}

	//upload file to filesystem
	filename, err := UploadFileIntoFS(r)
	if err != nil {
		log.Println(err)
	}

	//prepare additional attributes and insert into DB
	create_date := time.Now().Format("Mon Jan _2 15:04:05 2006")
	comment := r.FormValue("comment")

	InsertFileIntoDB(filename, comment, create_date)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if IsUnauthorized(w, r) {
		return
	}

	if r.Method == "POST" {
		id := r.URL.Query().Get("id")
		id = html.EscapeString(id)
		path := GetFilePathFromDB(id)
		DeleteFromFS(path)
		DeleteFromDB(id)
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func ServeStaticFiles(w http.ResponseWriter, r *http.Request) {
	if IsUnauthorized(w, r) {
		return
	}

	fs := http.FileServer(http.Dir("files"))
	fileHandler := http.StripPrefix("/files/", fs).ServeHTTP
	fileHandler(w, r)
}

func main() {
	fmt.Printf("Starting server on http://localhost:%d\n", port)
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/upload", UploadHandler)
	http.HandleFunc("/delete", DeleteHandler)
	http.HandleFunc("/files/", ServeStaticFiles)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

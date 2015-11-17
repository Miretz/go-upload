package main

import (
	"fmt"
	"database/sql"
	"net/http"
	"io"
	"log"
	"os"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	"text/template"
	"time"
	"crypto/md5"
	"html"
)

const port int = 3000
const uploadDir string = "./files/"
const dbfile string = "./data.db"

var db *sql.DB

const crTbQ string = `CREATE TABLE IF NOT EXISTS files 
(
	id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
	name VARCHAR(255) NOT NULL,
	path TEXT NOT NULL,
	comment TEXT,
	create_date TEXT NOT NULL	
)`


func init(){

	//Init database
	var err error
	db, err = sql.Open("sqlite3", dbfile)
	if err != nil {
		log.Fatal(err)
	}

	//Create table if not exists
	_, err = db.Exec(crTbQ)
	if err != nil {
		log.Printf("%q: %s\n", err, crTbQ)
		return
	}
}

type File struct {
	Id string
	Name string
	Path string
	Comment string
	Create_date string
}

func IndexHandler(w http.ResponseWriter, r *http.Request){
	rows, _ := db.Query("SELECT id, name, path, comment, create_date FROM files")
	defer rows.Close()
	data := make([]File,0)
	for rows.Next(){
		var id string
		var name string 
		var path string
		var comment string
		var create_date string
		rows.Scan(&id, &name, &path, &comment, &create_date)
		f := File{id, name, path, comment, create_date}
		data = append(data, f)
	}
	t, _ := template.ParseFiles("./tmpl/index.gtpl")
	t.Execute(w, data)
}

func UploadHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
        t, _ := template.ParseFiles("./tmpl/upload.gtpl")
		t.Execute(w, token)
	} else {
        //64MB
		r.ParseMultipartForm(64 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		f, err := os.OpenFile(uploadDir+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
		    fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		create_date := time.Now().Format("Mon Jan _2 15:04:05 2006")
		comment := r.FormValue("comment")
		db.Exec("INSERT INTO files(name, path, comment, create_date) VALUES ($1, $2, $3, $4)", 
				handler.Filename, "/files/"+handler.Filename, html.EscapeString(comment), create_date)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}

func DeleteHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		id := r.URL.Query().Get("id")
		id = html.EscapeString(id)
		var path string
		db.QueryRow("SELECT path FROM files WHERE id = $1", id).Scan(&path)
		os.Remove("." + path)
		db.Exec("DELETE FROM files WHERE id = $1", id)
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func main(){
	fmt.Printf("Starting server on http://localhost:%d\n", port)
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/upload", UploadHandler)
	http.HandleFunc("/delete", DeleteHandler)
	fs := http.FileServer(http.Dir("files"))
	http.Handle("/files/", http.StripPrefix("/files/", fs))
	http.ListenAndServe(":" + strconv.Itoa(port), nil)
}
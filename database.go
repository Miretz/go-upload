package main

import (
	"database/sql"
	"log"
	"html"
)

const databaseFile string = "./data.db"
const createTableQuery string = `CREATE TABLE IF NOT EXISTS files
(
	id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
	name VARCHAR(255) NOT NULL,
	path TEXT NOT NULL,
	comment TEXT,
	create_date TEXT NOT NULL
)`

const listFilesQuery string = "SELECT id, name, path, comment, create_date FROM files"
const insertFileQuery string = "INSERT INTO files(name, path, comment, create_date) VALUES ($1, $2, $3, $4)"
const getPathByIdQuery string = "SELECT path FROM files WHERE id = $1"
const deleteByIdQuery string = "DELETE FROM files WHERE id = $1"

var db *sql.DB

func InitializeDB(){
	//Init database
	var err error
	db, err = sql.Open("sqlite3", databaseFile)
	if err != nil {
		log.Fatal(err)
	}

	//Create table if not exists
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Printf("%q: %s\n", err, createTableQuery)
		return
	}
}

func ListFilesInDB() []File {
	data := make([]File,0)

	rows, err := db.Query(listFilesQuery)
	if err != nil {
		log.Printf("%q: %s\n", err, listFilesQuery)
		return data
	}

	defer rows.Close()
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

	return data
}

func InsertFileIntoDB(filename, comment, create_date string){
	_, err := db.Exec(insertFileQuery, filename, "/files/"+filename, html.EscapeString(comment), create_date)
	if err != nil {
		log.Printf("%q: %s\n", err, insertFileQuery)
	}
}

func GetFilePathFromDB(id string) string {
	var path string
	db.QueryRow(getPathByIdQuery, id).Scan(&path)
	return path
}

func DeleteFromDB(id string){
	db.Exec(deleteByIdQuery, id)
}
package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"

	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"

	"github.com/labstack/echo/v4"
)

var (
	AppName   = "meet"
	TableName = "users"
	db        *sql.DB
)

func createTable(db *sql.DB) {
	sts := `
DROP TABLE IF EXISTS users;
CREATE TABLE IF NOT EXISTS users(
    userid INTEGER PRIMARY KEY AUTOINCREMENT,
    Username  VARCAR(250) NOT NULL ,
    Password  varcar(250) NOT NULL,
    Email     VARCAR(250) NOT NULL,
    Gender    VARCAR(250) NOT NULL,
    Age       INT(2) NOT NULL,
    Height    INT NOT NULL,
    Weight    INT NOT NULL,
    Lang      VARCAR(250) NOT NULL,
    Profess   VARCAR(250) NOT NULL,
    Contry    VARCAR(250) NOT NULL,
    Descript  TEXT NOT NULL,
    Photos    TEXT NOT NULL
);
`
	_, err := db.Exec(sts)
	if err != nil {
		log.Fatal(err)
	}

}

// init database
func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./database.sql")
	if err != nil {
		log.Println("open database error: ", err)
	}
	createTable(db)
	fmt.Println("table users created")
	return db
}

// CREATE DATABASE ;
func CreateDB(dbName string, db *sql.DB) {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		panic(err)
	}
}

// init templates

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// path file is depends to enveronment.

func Templs(path string) *Template {
	return &Template{templates: template.Must(template.ParseFiles(listFiles(path)...))}
}

// listFiles return list filenames os spicific dir
// use paht.wolkFile insteade

func listFiles(dir string) (list []string) {

	f, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	sublist := make([]string, 0)
	root := dir + "/"
	for _, v := range files {
		root = dir + "/"
		if v.IsDir() {
			root = root + v.Name()
			sublist = listFiles(root)
			//for _, filename := range sublist {
			//	list = append(list, filename)
			//}
			continue
		}
		list = append(list, root+v.Name())
	}
	for _, f := range sublist {
		list = append(list, f)
	}

	return list
}

// folder when photos is stored.

func PhotoFold() string {
	//if os.Getenv("USERNAME") == "fedor" {
	//	return "/home/fedor/repo/files/"
	//}
	return "../files/"
}

// where is assets  path ?
func Assets() string {
	home, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	if home != "/Users/fedora/repo/meet" {
		return "/root/social/assets"
	}
	return "assets"
}

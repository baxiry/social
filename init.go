package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
)

var (
	AppName   = "meet"
	TableName = "users"
	db        *sql.DB
)

func ConnectDB() *sql.DB {
	db, err := sql.Open(
		"mysql", "root:123456@tcp(127.0.0.1:3306)/?charset=utf8&parseTime=True&loc=Local")
	if err != nil {

		log.Println("open database error: ", err)
		switch {
		case strings.Contains(err.Error(), "connection refused"):
			// TODO handle errors by code of error not by strings.

			//cmd := exec.Command("mysql.server", "restart")
			// for systemd linux : exec.Command("sudo", "service", "mariadb", "start")
			//cmd.Stdin = strings.NewReader(os.Getenv("JAWAD"))
			//err = cmd.Run()
			if err != nil {
				fmt.Println("error when run database cmd ", err)
			}
		default:
			log.Println("not knowen err at db.Ping() func")
			log.Println("unknown this error", err)
			os.Exit(1)
			//return nil
		}
	}

	return db
}

func CreateTable(dbName, tableName string, db *sql.DB) {
	//CREATE DATABASE ;
	//tbname := dbName + "." + tableName   // ` + tbname + `
	_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS social.users  (
    userid int unsigned NOT NULL AUTO_INCREMENT,
    username varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    email varchar(255) UNIQUE NOT NULL,
    photos text NOT NULL DEFAULT "",
    number_photos int NOT NULL DEFAULT 0,
    PRIMARY KEY (userid)
);`)
	if err != nil {
		panic(err)
	}
}

func CreateDB(dbName string, db *sql.DB) {
	//CREATE DATABASE ;
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

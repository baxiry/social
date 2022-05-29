package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// path file is depends to enveronment.
func templ() *Template {
	var p string
	//if os.Getenv("USERNAME") != "fedor" {
	//	p = "/root/store/"
	//}
	files := []string{
		p + "tmpl/home.html", p + "tmpl/upacount.html", p + "tmpl/acount.html", p + "tmpl/login.html", p + "tmpl/sign.html",
		p + "tmpl/404.html", p + "tmpl/updateProd.html",
		p + "tmpl/stores.html", p + "tmpl/mystore.html", p + "tmpl/notfound.html",
		p + "tmpl/upload.html", p + "tmpl/product.html", p + "tmpl/products.html",
		p + "tmpl/part/header.html", p + "tmpl/part/footer.html", p + "tmpl/updatefotos.html",
	}
	return &Template{templates: template.Must(template.ParseFiles(files...))}
}

// folder when photos is stored.

func photoFold() string {
	//if os.Getenv("USERNAME") == "fedor" {
	//	return "/home/fedor/repo/files/"
	//}
	return "../files/"
}

// where assets  path ?
func assets() string {
	home, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	if home != "/Users/fedora/repo/store" {
		return "/root/store/assets"
	}
	return "assets"
}

// init database
var (
	db  *sql.DB
	err error
)

func setdb() *sql.DB {
	db, err = sql.Open(
		"mysql", "root:123456@tcp(127.0.0.1:3306)/?charset=utf8&parseTime=True&loc=Local")
	if err != nil { // why no error when db is not runinig ??
		fmt.Println("run mysql server", err)
		// TODO report this error.

		// wehen db is stoped no error is return.
		// we expecte errore no database is runing

		// my be this error is fixed with panic ping pong bellow
	}

	if err = db.Ping(); err != nil {
		// TODO handle this error: dial tcp 127.0.0.1:3306: connect: connection refused
		fmt.Println("mybe database is not runing or error is: ", err)
		os.Exit(1)
	}
	return db
}

var catigories = map[string][]string{
	"cars":      {"mersides", "volswagn", "shefrole", "ford", "jarary", "jawad"},
	"animals":   {"dogs", "sheeps", "elephens", "checkens", "lions"},
	"motors":    {"harly", "senteroi", "basher", "hddaf", "mobilite"},
	"mobiles":   {"sumsung", "apple", "oppo", "netro", "nokia"},
	"computers": {"dell", "toshipa", "samsung", "hwawi", "hamed"},
	"services":  {"penter", "developer", "cleaner", "shooter", "gamer"}, //services
	"others":    {"somthing", "another-somth", "else", "anythings"},
}

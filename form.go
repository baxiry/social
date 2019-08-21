package main

import (
	"fmt"
	"html/template"
	"os"
	// "gopkg.in/mgo.v2"
    "github.com/gin-gonic/gin"
)

func getData(c *gin.Context) {
		title := c.PostForm("title")
		articl := c.PostForm("articl")

        fmt.Printf("title is : %s,  articl is : %s", title, articl)
	}
    c.JSON(200, gin.H{title: articl})
}

func postData(c *gin.Context) {
    c.JSON(200, gin.H{"postmsg": "posted hello world"})
}


func main() {
    r := gin.Default()
    r.GET("/", getData)
    r.POST("/", postData)
    r.Run()
}














type Book struct {
	Name    string
	Subject string
	Author  string
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	CheckeErr("err in create session: maybe mongodb not active", err)

	defer session.Close()

	tmpl := template.Must(template.ParseFiles("asset/form.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
	    tmpl.Execute(w, nil)
	    return
	}

		c := session.DB("test").C("Book")

		book := Book{
			Name:    r.FormValue("name"),
			Subject: r.FormValue("subject"),
			Author:  r.FormValue("author"),
		}
		err = c.Insert(&Book{book.Name, book.Subject, book.Author})

		// do something with details
		fmt.Println(book.Name)
		fmt.Println(book.Subject)

		tmpl.Execute(w, struct{ Ok bool }{true})
	})

	http.ListenAndServe(":8080", nil)
}

func CheckeErr(str string, err error) {
	if err != nil {
		fmt.Println(str, err)
		os.Exit(0)
	}
}

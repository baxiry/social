package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func main() {

	db = setdb()
	defer db.Close()

	e := echo.New()

	// TODO store secret key in envrenment
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.Renderer = templ()

	// files
	e.Static("/assets", assets())
	e.Static("/fs", photoFold())

	// account and verefy
	e.GET("/", homePage)
	e.GET("/sign", signPage)
	e.POST("/sign", signup)
	e.GET("/login", loginPage)
	e.POST("/login", login)
	e.GET("/user/:id", Profile)
	e.GET("/fotos", PhotosPage)
	e.POST("/upfotos", UpPhotos)
	e.GET("/upacount", UpdatePage)
	e.POST("/upacount", Update)

	e.GET("/messages", MessagesPage)
	e.GET("/example", examplePage)
	e.GET("/activity", Activity)
	e.GET("/search", Search)

	//e.POST("/updatefotos/:id", updateProdFotos)

	//e.GET("/:catigory/:id", getOneProd) // whech is beter ? :catigory or /product ?

	e.Logger.Fatal(e.Start(":8080"))
}

func examplePage(c echo.Context) error {

	sess, _ := session.Get("session", c)
	username := sess.Values["username"]
	userid := sess.Values["userid"]

	data := make(map[string]interface{}, 3)

	data["username"] = username
	data["userid"] = userid
	users := getRecentUsers()
	for _, u := range users {
		fmt.Println(u.UserId)
	}
	data["users"] = users
	fmt.Println(c.Render(http.StatusOK, "example.html", data))
	return nil
}

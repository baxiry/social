package main

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func main() {

	db = setdb()
	defer db.Close()

	e := echo.New()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.Renderer = templ()

	// files
	e.Static("/assets", assets())
	e.Static("/fs", photoFold())

	// account and verefy
	e.GET("/", homePage)
	e.GET("/sign", SignPage)
	e.POST("/sign", Signup)
	e.GET("/login", LoginPage)
	e.POST("/login", Login)
	e.GET("/user/:id", Profile)
	e.GET("/fotos", PhotosPage)
	e.POST("/upfotos/:id", UpPhotos)
	e.GET("/upacount", UpdatePage)
	e.POST("/upacount", Update)

	e.GET("/messages", MessagesPage)
	e.GET("/activity", Activity)
	e.GET("/search", Search)

	//e.POST("/updatefotos/:id", updateProdFotos)

	//e.GET("/:catigory/:id", getOneProd) // whech is beter ? :catigory or /product ?

	e.Logger.Fatal(e.Start(":8080"))

}

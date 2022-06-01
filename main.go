package main

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func main() {

	db := setdb()
	defer db.Close()

	e := echo.New()
	//SetupRoutes()
	//e.GET("/a", contentHandler, contentRewrite)
	//e.HTTPErrorHandler = customHTTPErrorHandler

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
	e.GET("/acount/:id", acount)
	e.GET("/upacount", updateAcount)
	e.POST("/upacount", updateAcountInfo)

	// store and product
	e.GET("/mystore", myStores)
	e.GET("/stores", stores)
	e.GET("/:catigory", getProds) // ??
	e.GET("/product/:id", getOneProd)
	e.GET("/update/:id", updateProdPage)
	e.POST("/update/:id", updateProd)
	e.GET("/delete/:id", deleteProd)

	e.GET("/updatefotos/:id", updateFotosPage)
	e.POST("/updatefotos/:id", updateProdFotos)

	e.GET("/upload", createProductPage)
	e.POST("/upload", createProduct)
	//e.GET("/:catigory/:id", getOneProd) // whech is beter ? :catigory or /product ?

	// not found pages
	//e.GET("/:ok/:ok/:ok", notFoundPage)
	//e.GET("/:ok/", notFoundPage)

	e.Logger.Fatal(e.Start(":8080"))
}

package main

import (
	"net/http"
	"social/im"

	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func main() {

	db = ConnectDB()
	defer db.Close()

	e := echo.New()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.Renderer = Templs("templates")

	// files
	e.Static("/assets", Assets())
	e.Static("/fs", PhotoFold())

	// account and verefy
	e.GET("/", HomePage)
	e.GET("/sign", SignPage)
	e.POST("/sign", Signup)
	e.GET("/login", LoginPage)
	e.POST("/login", Login)
	e.GET("/user/:id", Profile)
	e.GET("/fotos", PhotosPage)
	e.POST("/upfotos/:id", UpPhotos)
	e.GET("/upacount", UpdatePage)
	e.POST("/upacount", Update)

	//e.GET("/messages", im.MessagesPage)
	e.GET("/activity", ActivityPage)
	e.GET("/search", SearchPage)
	e.GET("/ws", WsHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func WsHandler(e echo.Context) error {
	//if r.Header.Get("Origin")!="http://"+r.Host {http.Error(w,"Origin not allowed",-1);return}

	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(e.Response().Writer, e.Request(), nil)
	if err != nil {
		http.Error(e.Response().Writer, "Could not open websocket connection", 404)
		return err
	}

	go im.ServeMessages(conn)
	return nil
}

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func login(c echo.Context) error {
	femail := c.FormValue("email")
	fpass := c.FormValue("password")
	userid, name, email, pass := getUsername(femail)

	if pass == fpass && femail == email {
		//userSession[email] = name
		NewSession(c, name, userid)
		return c.Redirect(http.StatusSeeOther, "/") // 303 code
		// TODO redirect to latest page
	}
	// TODO flush this message
	fmt.Println(c.Render(200, "login.html", "username or pass is wrong!"))
	return nil
}

// signup sing up new user handler
func signup(c echo.Context) error {
	name := c.FormValue("username")
	pass := c.FormValue("password")
	email := c.FormValue("email")
	phon := c.FormValue("phon")
	err := insertUser(name, pass, email, phon)
	if err != nil {
		//fmt.Println(err)
		return c.Render(200, "sign.html", "wrrone")
	}
	return c.Redirect(http.StatusSeeOther, "/login") // 303 code
}

func signPage(c echo.Context) error {
	return c.Render(200, "sign.html", "hello")
}

func loginPage(c echo.Context) error {
	fmt.Println(c.Render(200, "login.html", "hello"))
	return nil
}

// insertUser register new user in db
func insertUser(user, pass, email, phon string) error {
	insert, err := db.Query(
		"INSERT INTO stores.users(username, password, email, phon) VALUES ( ?, ?, ?, ? )",
		user, pass, email, phon)

	// if there is an error inserting, handle it
	if err != nil {
		return err
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
	return nil
}

// get user info
func getUsername(femail string) (int, string, string, string) {
	var name, email, password string
	var userid int
	err := db.QueryRow(
		"SELECT userid, username, email, password FROM stores.users WHERE email = ?",
		femail).Scan(&userid, &name, &email, &password)
	if err != nil {
		fmt.Println("no result or", err.Error())
	}
	return userid, name, email, password
}

// newSession creates new session
func NewSession(c echo.Context, username string, userid int) {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60, // = 1h,
		HttpOnly: true,    // no websocket or any thing else
	}
	sess.Values["username"] = username
	sess.Values["userid"] = userid
	sess.Save(c.Request(), c.Response())
}

package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// signup sing up new user handler
func signup(c echo.Context) error {
	email := c.FormValue("email")
	pass := c.FormValue("password")
	fmt.Println(email, pass)
	err := insertUser(email, pass)
	if err != nil {
		fmt.Println(err)
		return c.Render(200, "sign.html", err.Error())
	}
	return c.Redirect(http.StatusSeeOther, "/login") // 303 code
}

// insertUser register new user in db
func insertUser(email, pass string) error {
	insert, err := db.Query(
		"INSERT INTO social.users(email, password) VALUES ( ?, ?)",
		email, pass)

	// if there is an error inserting, handle it
	if err != nil {
		return err
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
	return nil
}

func login(c echo.Context) error {
	femail := c.FormValue("email")
	fpass := c.FormValue("password")
	userid, username, pass := selectUser(femail)

	if pass == fpass && pass != "" {
		//userSession[email] = name
		NewSession(c, username, userid)
		return c.Redirect(http.StatusSeeOther, "/") // 303 code
		// TODO redirect to latest page
	}
	// TODO flush this message
	fmt.Println(c.Render(200, "login.html", "username or pass is wrong!"))
	return nil
}

// select User info
func selectUser(femail string) (int, string, string) {
	var username, password string
	var userid int
	err := db.QueryRow(
		"SELECT userid, username, password FROM social.users WHERE email = ?",
		femail).Scan(&userid, &username, &password)
	if err != nil {
		fmt.Println(err.Error())
	}
	return userid, username, password
}

func signPage(c echo.Context) error {
	return c.Render(200, "sign.html", "hello")
}

func loginPage(c echo.Context) error {
	fmt.Println(c.Render(200, "login.html", "hello"))
	return nil
}

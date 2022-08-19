package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// signup sing up new user handler
func Signup(c echo.Context) error {

	email := c.FormValue("email")
	pass := c.FormValue("password")
	m := c.FormValue("man")
	f := c.FormValue("femane")

	gender := ""
	if m == "on" {
		gender = "m"
	}
	if f == "on" {
		gender = "f"
	}

	fmt.Println(email, pass, "gender is : ", gender)

	err := insertUser(email, pass, gender)
	if err != nil {
		fmt.Println(err)
		return c.Render(200, "sign.html", err.Error())
	}
	return c.Redirect(http.StatusSeeOther, "/login") // 303 code
}

// insertUser register new user in db
func insertUser(email, pass, gender string) error {

	insert, err := db.Query(
		"INSERT INTO social.users(email, password, gender) VALUES ( ?, ?, ?)",
		email, pass, gender)

	// if there is an error inserting, handle it
	if err != nil {
		return err
	}
	// be careful deferring Queries if you are using transactions
	insert.Close()
	return nil
}

func Login(c echo.Context) error {
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

func SignPage(c echo.Context) error {
	return c.Render(200, "sign.html", "")
}

func LoginPage(c echo.Context) error {
	fmt.Println(c.Render(200, "login.html", ""))
	return nil
}

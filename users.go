package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// gets all user information for update this info
func getUserInfo(userid int) (string, string, string) {
	var username, email, photos string
	err := db.QueryRow(
		"SELECT username, email,photos FROM social.users WHERE userid = ?",
		userid).Scan(&username, &email, &photos)
	if err != nil {
		fmt.Println("no result or", err.Error())
	}
	return username, email, photos
}

// update user info in db
func updateUserInfo(name, email, phon string, userid int) error {

	//Update db
	stmt, err := db.Prepare("update stores.users set username=?, email=?, phon=? where userid=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// execute
	res, err := stmt.Exec(name, email, phon, userid)
	if err != nil {
		return err
	}

	a, err := res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println("efected foto update: ", a) // 1
	return nil
}

// updateAcount updates Acount information
func updateAcountInfo(c echo.Context) error {
	//data := make(map[string]interface{},1)
	sess, _ := session.Get("session", c)
	uid := sess.Values["userid"]
	if uid == nil {
		// login first
		return c.Redirect(http.StatusSeeOther, "/login") // 303 code
	}

	name := c.FormValue("name")
	email := c.FormValue("email")
	phon := c.FormValue("phon")
	fmt.Println("name+email+phon is :", name, email, phon)

	err := updateUserInfo(name, email, phon, uid.(int))
	if err != nil {
		fmt.Println("error at update db function", err)
	}

	// update session information
	NewSession(c, name, uid.(int))

	// redirect to acoun page
	userid := strconv.Itoa(uid.(int))

	return c.Redirect(303, "/acount/"+userid)
}

// updateAcount updates Acount information
func updateAcount(c echo.Context) error {
	data := make(map[string]interface{}, 1)
	sess, _ := session.Get("session", c)

	uid := sess.Values["userid"]
	username := sess.Values["username"]

	data["username"] = username

	if uid == nil {
		// login first
		return c.Redirect(http.StatusSeeOther, "/login") // 303 code
	}

	data["username"], data["email"], data["phon"] = getUserInfo(uid.(int))

	data["userid"] = uid

	fmt.Println(data)

	return c.Render(200, "upacount.html", data)
}

// acount render profile of user.
func acount(c echo.Context) error {
	sess, _ := session.Get("session", c)
	data := make(map[string]interface{}, 2)
	data["username"] = sess.Values["username"]
	fmt.Println("username is ", data["username"])
	data["userid"] = sess.Values["userid"]
	fmt.Println("user id or user is : ", data["userid"])
	// TODO get all info like foto from db

	if data["userid"] == nil {
		return c.Redirect(http.StatusSeeOther, "/login") // 303 code
	}
	return c.Render(200, "acount.html", data)
}

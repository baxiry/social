package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// gets all user information for update this info
func getOneUser(userid int) (string, string, string, string) {
	var name, email, phon, avatar string
	err := db.QueryRow(
		"SELECT username, email,phon, linkavatar FROM stores.users WHERE userid = ?",
		userid).Scan(&name, &email, &phon, &avatar)
	if err != nil {
		fmt.Println("no result or", err.Error())
	}
	return name, email, phon, avatar
}

// update user info in db
func GetUsers(name, email, phon string, userid int) error {

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
func Search_one(c echo.Context) error {
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

	err := updateUserInfo(name, uid.(int))
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
func Search(c echo.Context) error {
	data := make(map[string]interface{}, 1)
	sess, _ := session.Get("session", c)

	uid := sess.Values["userid"]
	if uid == nil {
		// login first
		return c.Redirect(http.StatusSeeOther, "/login") // 303 code
	}

	username := sess.Values["username"]
	data["userid"] = uid
	data["username"] = username

	fmt.Println(data)

	return c.Render(200, "upacount.html", data)
}

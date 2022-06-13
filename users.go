package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Profile(c echo.Context) error {
	username, userid, err := GetSession(c)
	if err != nil {
		//println(c.Redirect(http.StatusSeeOther, "/login"))
		fmt.Println("error of upacount handler is ", err)
		//return nil

	}

	id, _ := strconv.Atoi(c.Param("id"))

	data := make(map[string]interface{}, 1)
	data["user"] = getUserInfo(id)

	// for session
	data["userid"] = userid
	data["username"] = username

	fmt.Println(c.Render(200, "user.html", data))
	return nil
}

// updateAcount updates Acount information
func UpdatePage(c echo.Context) error {
	fmt.Println("update account")

	username, userid, err := GetSession(c)
	if err != nil {
		println(c.Redirect(http.StatusSeeOther, "/login"))
		fmt.Println("error of upacount handler is ", err)
		return nil
	}

	data := make(map[string]interface{}, 1)

	data["username"] = username
	data["userid"] = userid
	data["user"] = getUserInfo(userid)

	fmt.Println(data)

	fmt.Println(c.Render(200, "upacount.html", data))
	return nil
}

// updateAcount updates Acount information
func Update(c echo.Context) error {
	fmt.Println("update account")

	username, userid, err := GetSession(c)
	if err != nil {
		println(c.Redirect(http.StatusSeeOther, "/login"))
		fmt.Println("error of upacount handler is ", err)
		return nil
	}

	data := make(map[string]interface{}, 1)

	data["username"] = username
	data["userid"] = userid
	data["user"] = getUserInfo(userid)

	fmt.Println(data)

	fmt.Println(c.Render(200, "user.html", data))
	return nil
}

// getUserIfor from db
func getUserInfo(userid int) (user User) {
	err := db.QueryRow(
		"SELECT username, email, photos FROM social.users WHERE userid = ?",
		userid).Scan(&user.Username, &user.Email, &user.Photos)

	if err != nil {
		fmt.Println("no result or", err.Error())
	}
	return user
}

// update user info in db
func updateUserInfo(field string, userid int) error {

	//Update db
	stmt, err := db.Prepare("update social.users set " + field + "=? where userid=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// execute
	res, err := stmt.Exec(field, userid)
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
func UpdateInfo(c echo.Context) error {
	//data := make(map[string]interface{},1)
	sess, _ := session.Get("session", c)
	uid := sess.Values["userid"]
	if uid == nil {
		// login first
		return c.Redirect(http.StatusSeeOther, "/login") // 303 code
	}
	fmt.Println("we are on update user info post")
	name := c.FormValue("name")

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

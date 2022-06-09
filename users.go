package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func UserProfile(c echo.Context) error {

	data := make(map[string]interface{}, 1)

	id, _ := strconv.Atoi(c.Param("id"))

	data["name"], data["email"], data["photo"] = getUserInfo(id)

	fmt.Println(c.Render(200, "user.html", data))
	return nil
}

func Profile(c echo.Context) error {

	data := make(map[string]interface{}, 1)

	data["username"], data["userid"], _ = GetSession(c)

	id, _ := strconv.Atoi(c.Param("id"))

	data["name"], data["email"], data["photo"] = getUserInfo(id)

	fmt.Println(c.Render(200, "user.html", data))
	return nil
}

// updateAcount updates Acount information
func updateAcount(c echo.Context) error {
	sess, _ := session.Get("session", c)
	uid := sess.Values["userid"]
	if uid == nil {
		// login first
		return c.Redirect(http.StatusSeeOther, "/login") // 303 code
	}
	username := sess.Values["username"]

	data := make(map[string]interface{}, 1)
	data["username"] = username
	data["username"], data["email"], data["photos"] = getUserInfo(uid.(int))
	data["userid"] = uid
	fmt.Println(data)
	fmt.Println(c.Render(200, "upacount.html", data))
	return nil
}

// getUserIfor from db
func getUserInfo(userid int) (string, string, string) {
	var username, email, photos string
	err := db.QueryRow(
		"SELECT username, email, photos FROM social.users WHERE userid = ?",
		userid).Scan(&username, &email, &photos)

	if err != nil {
		fmt.Println("no result or", err.Error())
	}
	return username, email, photos
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
func updateAcountInfo(c echo.Context) error {
	//data := make(map[string]interface{},1)
	sess, _ := session.Get("session", c)
	uid := sess.Values["userid"]
	if uid == nil {
		// login first
		return c.Redirect(http.StatusSeeOther, "/login") // 303 code
	}

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

// acount render profile of user.
func acount(c echo.Context) error {
	sess, _ := session.Get("session", c)
	userid := sess.Values["userid"]

	if userid == nil {
		return c.Redirect(http.StatusSeeOther, "/login") // 303 code
	}

	data := make(map[string]interface{}, 2)

	data["username"] = sess.Values["username"]
	data["userid"] = userid

	_, _, data["photos"] = getUserInfo(userid.(int))

	return c.Render(200, "acount.html", data)
}

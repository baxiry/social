package main

import (
	"fmt"
	"meet/auth"
	"meet/helps"
	"net/http"
	"strings"

	"github.com/blockloop/scan"
	"github.com/labstack/echo/v4"
)

func HomePage(c echo.Context) error {

	username, userid, err := auth.GetSession(c)

	helps.PrintError("get session", err)
	println("userid is ", userid)
	println("username is ", username)

	data := make(map[string]interface{}, 3)
	data["username"] = username
	data["userid"] = userid
	users := getRecentUsers()

	fmt.Print("profile info", ProfileInfo(userid), "\n\n")
	for i := range users {
		photos := strings.Split(users[i].Photos, "; ")
		users[i].Photos = setAvatar(users[i].Gender, photos[0])
	}

	data["users"] = users
	fmt.Println(c.Render(http.StatusOK, "home.html", data))
	return nil
}

// getCatigories get all photo name of catigories.
func getRecentUsers() (users []User) {
	rows, err := db.Query("SELECT userid, username, email, photos, gender from social.users;")
	defer rows.Close()

	err = scan.Rows(&users, rows)
	helps.PrintError("error from schan.Rows: ", err)

	return users
}

func ProfileInfo(userid int) (profile User) {
	rows, err := db.Query("select * from social.users where userid = ?", userid)
	if err != nil {
		fmt.Println("\nat query func owner id db select ", err)
	}
	defer rows.Close() // ??
	err = scan.Rows(&profile, rows)
	println(err)
	return profile
}

func setAvatar(gen, photo string) string {
	if photo != "" {
		return photo
	}
	if gen == "m" {
		return "bman.jpg"
	}
	if gen == "f" {
		return "bwoman.jpg"

	}
	return ""
}

package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type User struct {
	UserId       uint64
	Username     string
	Password     string
	Email        string
	Gender       string
	Photos       string
	NemberPhotos int
}

func homePage(c echo.Context) error {
	sess, _ := session.Get("session", c)
	username := sess.Values["username"]
	userid := sess.Values["userid"]

	data := make(map[string]interface{}, 3)

	data["username"] = username
	data["userid"] = userid
	users := getRecentUsers()
	for _, u := range users {
		fmt.Println(u)
	}
	data["users"] = users
	fmt.Println(c.Render(http.StatusOK, "home.html", data))
	return nil
}

//getCatigories get all photo name of catigories.
func getRecentUsers() (users []User) {
	res, err := db.Query(
		"SELECT userid, username, email, photos from social.users;")
	if err != nil {
		fmt.Println("getResentUsers error : ", err)
	}
	defer res.Close()

	var u User
	for res.Next() {
		res.Scan(&u.UserId, &u.Username, &u.Email, &u.Photos)
		users = append(users, u)
	}

	for _, u := range users {
		fmt.Println(u)
	}
	return users
}

/*

type Profile struct{}

func Profile(profileId int) (profile Profile) {
	rows, err := db.Query("select productID, catigory, title, description, photos, price from stores.products where ownerid = ?", ownerid)
	if err != nil {
		fmt.Println("at query func owner id db select ", err)
	}
	defer rows.Close() // ??

	// iterate over rows
	for rows.Next() {
		err = rows.Scan(&p.ProductId, &p.Catigory, &p.Title, &p.Description, &p.Photo, &p.Price)
		if err != nil {
			fmt.Println("err when getting Porducts from db. at rews.Next()", err)
			return nil
		}
		if p.Photo == "" {
			fmt.Println("no fotots")
		}
		products = append(products, p)

	}
	return profile
}

*/

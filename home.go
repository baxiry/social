package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// notFoundPage
//func notFoundPage(c echo.Context) error {
//  return c.Render(200, "notfound.html", nil)
//}

func homePage(c echo.Context) error {

	sess, _ := session.Get("session", c)
	username := sess.Values["username"]
	uid := sess.Values["userid"]
	//fmt.Println("name is : ", name)

	data := make(map[string]interface{}, 3)
	data["username"] = username
	data["userid"] = uid
	fmt.Println(c.Render(http.StatusOK, "home.html", data))
	return nil
}

// getCatigories get all photo name of catigories.
// res, err := db.Query(
//		"SELECT productId, title, photos, price FROM stores.products WHERE catigory = ?", catigory)
/*
	for res.Next() {
		res.Scan(&p.ProductId, &p.Title, &picts, &p.Price)
		list := strings.Split(picts, "];[")
		p.Photo = list[0]
		items = append(items, p)
		// TODO we need just avatar photo
	}

*/

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

package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// TODO redirect to latest page after login.

// getCatigories get all photo name of catigories.
func getProductes(catigory string) ([]Product, error) {
	var p Product
	var picts string
	res, err := db.Query(
		"SELECT productId, title, photos, price FROM stores.products WHERE catigory = ?", catigory)
	if err != nil {
		fmt.Println("error in getProducts func", err)
		return nil, err
	}
	defer res.Close() // free result row

	items := make([]Product, 0)
	for res.Next() {
		res.Scan(&p.ProductId, &p.Title, &picts, &p.Price)
		list := strings.Split(picts, "];[")
		p.Photo = list[0]
		items = append(items, p)
		// TODO we need just avatar photo
	}
	return items, nil
}

// select All product from db
func myProducts(ownerid int) []Product {
	rows, err := db.Query("select productID, catigory, title, description, photos, price from stores.products where ownerid = ?", ownerid)
	if err != nil {
		fmt.Println("at query func owner id db select ", err)
	}
	defer rows.Close() // ??

	var products = []Product{}
	var p = Product{}

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
	return products
}

// perhaps is beter ignoring this feater ??!
func myStores(c echo.Context) error { // TODO rename to myproduct ??

	sess, _ := session.Get("session", c)
	username := sess.Values["username"]

	if username == nil {
		return c.Redirect(http.StatusSeeOther, "/login") // 303 code
	}

	data := make(map[string]interface{}, 3)
	userid := sess.Values["userid"]

	data["username"] = username
	data["userid"] = userid

	data["products"] = myProducts(userid.(int))

	fmt.Println(c.Render(200, "mystore.html", data))
	return nil
}

// getProduct get & render all data of one product.
func getProds(c echo.Context) error {
	data := make(map[string]interface{})

	sess, _ := session.Get("session", c)

	userid := sess.Values["userid"]

	catigory := c.Param("catigory") // TODO home or catigory.html ?
	fmt.Println("animals end point, caticory is : ", catigory)

	data["username"] = sess.Values["username"]
	data["userid"] = userid
	data["subCatigories"] = catigories[catigory] // from router.go
	data["products"], err = getProductes(catigory)
	if err != nil {
		fmt.Println("in gitCatigories: ", err)
	}

	fmt.Println(c.Render(http.StatusOK, "products.html", data))
	return nil
}

// perhaps is beter ignoring this feater ??!
func stores(c echo.Context) error {
	sess, _ := session.Get("session", c)
	userid := sess.Values["userid"]
	data := make(map[string]interface{}, 2)
	username := sess.Values["username"]

	data["username"] = username
	data["userid"] = userid
	return c.Render(200, "stores.html", data)
}

// TODO url := c.Request().URL  we need change url path ? example /cats/ to /cats

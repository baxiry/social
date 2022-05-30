package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type Product struct {
	ProductId   int
	Title       string
	Catigory    string
	Description string
	Photo       string
	Photos      []string
	Price       string
}

// TODO redirect to latest page after login.
func updateProdPage(c echo.Context) error {
	// TODO whish is beter all data of product or jast photo ?
	data := make(map[string]interface{})
	sess, _ := session.Get("session", c)
	data["username"] = sess.Values["username"]
	data["userid"] = sess.Values["userid"]

	// User ID from path `users/:id`
	pid := c.Param("id") // TODO home or catigory.html ?
	productId, _ := strconv.Atoi(pid)

	data["product"], err = selectProduct(productId)

	err = c.Render(http.StatusOK, "updateProd.html", data)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

// upload photos
func createProductPage(c echo.Context) error {
	// check session first
	sess, _ := session.Get("session", c)
	userid := sess.Values["userid"]
	if userid == nil {
		// TODO flash here
		return c.Redirect(http.StatusSeeOther, "/login") // 303 code
	}

	data := make(map[string]interface{}, 3)
	username := sess.Values["username"]
	data["username"] = username
	data["userid"] = userid
	return c.Render(200, "upload.html", data)
}

// select product from db
func selectProduct(productId int) (Product, error) {
	var p Product
	var picts string
	err := db.QueryRow(
		"SELECT title, catigory, description, photos, price FROM stores.products WHERE productId = ?",
		productId).Scan(&p.Title, &p.Catigory, &p.Description, &picts, &p.Price)
	if err != nil {
		return p, err
	}

	list := strings.Split(picts, "];[")
	fmt.Println("list fotos is :", list)
	// TODO split return 2 item in some casess, is this a bug ?
	p.Photos = filter(list)
	p.ProductId = productId
	return p, nil
}

// delete Producte from db.
func deleteProducte(productId int) error {
	res, err := db.Exec("DELETE FROM stores.products WHERE productId=?", productId)
	if err != nil {
		return err
	}

	affectedRows, err := res.RowsAffected()

	if err != nil {
		return err
	}
	fmt.Println("affectedRows: ", affectedRows)
	// defer res // TODO I need understand this close in mariadb
	return nil
}

// db
func updateProduct(title, catig, descr, price, photos string, productId int) error {
	// TODO chane price type.

	//Update db
	stmt, err := db.Prepare("update  stores.products set  title=?,  catigory=?, description=?,  price=?,  photos=? where productId=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// execute
	_, err = stmt.Exec(title, catig, descr, price, photos, productId)

	if err != nil {
		return err
	}
	/*
		a, err := res.RowsAffected()
		if err != nil {
			fmt.Println("error is: ", err)
			return err
		}
	*/
	return nil
}

// insert product to db
func insertProduct(title, catigory, details, picts string, ownerid int, price float64) error {
	insert, err := db.Query(
		"INSERT INTO stores.products(ownerID, title, catigory, description, price, photos) VALUES ( ?, ?, ?, ?, ?, ?)",
		ownerid, title, catigory, details, price, picts)
	// if there is an error inserting, handle it
	if err != nil {
		return err
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close() // TODO why we need closeing this connection ?

	return nil
}

// upload uploads new product
func createProduct(c echo.Context) error {
	// TODO: how upload this ?.  definde uploader by session
	sess, _ := session.Get("session", c)
	ownerid := sess.Values["userid"]
	if ownerid == nil {
		return err
	}
	// TODO mybe we need handle when session expired befoar appload

	title := c.FormValue("title")
	catigory := c.FormValue("catigory")
	details := c.FormValue("description")
	//price, e := strconv.Atoi(c.FormValue("price"))
	price, err := strconv.ParseFloat(c.FormValue("price"), 32)
	if err != nil {
		fmt.Println("error at ParseFloat", err)
		return err
	}

	// Read files, Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]
	//fmt.Println("files is :", files[0].Filename)
	picts := ""
	for _, v := range files {
		picts += v.Filename
		picts += "];["
		fmt.Println(picts)
		// TODO Rename pictures.
	}

	//  func insertProduct(title, catigory, details, picts string, ownerid, int64, price float32) error {
	err = insertProduct(title, catigory, details, picts, ownerid.(int), price)

	if err != nil {
		fmt.Println("error in insert product", err)
	}

	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			fmt.Println("error at file.Open() file is :", err)

			return err
		}
		defer src.Close()
		// Destination
		dst, err := os.Create(photoFold() + file.Filename)
		if err != nil {
			fmt.Println("error at io.Create file is :", err)
			return err
		}
		defer dst.Close()
		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			fmt.Println("error at io.Copy file is :", err)
			return err
		}
	}

	// TODO redirect to home or to acount ??
	err = c.Redirect(http.StatusSeeOther, "/") // 303 code
	if err != nil {
		fmt.Println("redirect err", err)
		return nil
	}
	return nil
}

// delete product
func deleteProd(c echo.Context) error {
	// TODO we need checkout sesston ?

	sess, _ := session.Get("session", c)
	ownerid := sess.Values["userid"]
	if ownerid == nil {
		return c.Redirect(http.StatusSeeOther, "/mystore")
	}

	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	err = deleteProducte(i)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// TODO return string to ajax resever
	return c.String(http.StatusOK, "success!")
}

// update Prodact
func updateProd(c echo.Context) error {

	pid := c.Param("id")
	id, err := strconv.Atoi(pid)
	if err != nil {
		fmt.Println("id error", err)
	}

	title := c.FormValue("title")
	catig := c.FormValue("catigory")
	descr := c.FormValue("description")
	price := c.FormValue("price")
	photos := c.FormValue("files")

	err = updateProduct(title, catig, descr, price, photos, id)
	if err != nil {
		// TODO send error to client with ajax
		fmt.Println("error when update product: ", err)
		return err
	}
	err = c.Redirect(http.StatusSeeOther, "/mystore")
	return nil
}

// TODO redirect to latest page after login.
func getOneProd(c echo.Context) error {

	sess, _ := session.Get("session", c)
	userid := sess.Values["userid"]
	username := sess.Values["username"]

	data := make(map[string]interface{})

	id := c.Param("id") // TODO home or catigory.html ?
	productId, _ := strconv.Atoi(id)

	data["product"], err = selectProduct(productId)

	if err != nil {
		fmt.Println("with gitCatigories: ", err)
	}
	data["userid"] = userid
	data["username"] = username

	// User ID from path `users/:id`
	return c.Render(http.StatusOK, "product.html", data)
}

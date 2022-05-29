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

// updateFotosPage router fo update Fotos Page
func updateFotosPage(c echo.Context) error {
	data := make(map[string]interface{})
	sess, _ := session.Get("session", c)   // TODO i need session ?
	data["username"] = sess.Values["name"] // TODO use user id instead name
	if data["username"] == nil {
		fmt.Println("session name is nil redirect to login")
		c.Redirect(303, "/login")
	}

	pid := c.Param("id")
	productId, _ := strconv.Atoi(pid)

	data["productFotos"], err = getProductFotos(productId)
	data["productId"] = productId
	fmt.Printf("product is : %#v", data["productFotos"])
	if err != nil {
		fmt.Println(err)
	}
	err := c.Render(http.StatusOK, "updatefotos.html", data)
	if err != nil {
		fmt.Println("\nerr is : ", err)
	}
	return nil
}

// update fotos name in database
func updateProductFotos(photos string, productId int) error {

	//Update db
	stmt, err := db.Prepare("update  stores.products set photos=? where productId=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// execute
	_, err = stmt.Exec(photos, productId)
	if err != nil {
		return err
	}

	return nil
}

// updateFotos updates photos of products
func updateProdFotos(c echo.Context) error {

	pid := c.Param("id")
	id, err := strconv.Atoi(pid)
	if err != nil {
		fmt.Println("id error", err)
	}

	// from her :
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
		// TODO Rename pictures.
	}

	// databas function
	err = updateProductFotos(picts, id)

	if err != nil {
		fmt.Println("error in update product foto", err)
	}

	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			fmt.Println("err in file.Open()")
			return err
		}
		defer src.Close()
		// Destination
		dst, err := os.Create(photoFold() + file.Filename)
		if err != nil {
			fmt.Println("err in io.Create()")
			return err
		}
		defer dst.Close()
		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			fmt.Println("err in io.Copy()")
			return err
		}
	}

	err = c.Redirect(http.StatusSeeOther, "/mystore")
	if err != nil {
		fmt.Println("\nerr when update product photo", err)
	}
	return nil
}

// selecte fotos from db
func getProductFotos(productId int) ([]string, error) {
	var picts string
	err := db.QueryRow(
		"SELECT photos FROM stores.products WHERE productId = ?",
		productId).Scan(&picts)
	if err != nil {
		return nil, err
	}
	list := strings.Split(picts, "];[")
	// TODO split return 2 item in some casess, is this a bug ?
	fotos := filter(list)
	return fotos, nil
}

// some tools
func filter(slc []string) []string {
	res := make([]string, 0)
	for _, v := range slc {
		if v != "" {
			res = append(res, v) // TODO this need improve fo performence
		}
	}
	return res
}

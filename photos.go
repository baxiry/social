package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

// phootosPage router fo update Fotos Page
func PhotosPage(c echo.Context) error {

	username, userid, err := GetSession(c)
	if err != nil {
		fmt.Println("session name is nil redirect to login", http.StatusSeeOther)
		c.Redirect(http.StatusSeeOther, "/login") // 303 code
		return nil
	}

	data := make(map[string]interface{}, 3)
	photos, err := getProductFotos(userid)
	data["photos"] = photos
	data["username"] = username

	fmt.Println("fotos is : ", data)

	fmt.Println(c.Render(http.StatusOK, "upfotos.html", data))
	return nil
}

// update fotos name in database
func UpdatePhotos(photos string, userid int) error {

	//Update db
	stmt, err := db.Prepare("update  social.users set photos=? where userid=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// execute
	_, err = stmt.Exec(photos, userid)
	if err != nil {
		return err
	}

	return nil
}

// updateFotos updates photos of products
func UpPhotos(c echo.Context) error {

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
	err = UpdatePhotos(picts, id)

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

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
	data["catigories"] = catigories
	err := c.Render(http.StatusOK, "home.html", data)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return nil
}

/* TODO handle error
func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
    errorPage := fmt.Sprint("/404.html", code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
    fmt.Println(err)
    //c.Redirect(303, "notfound.html")
    c.Redirect(http.StatusSeeOther, "/notfound") // 303 code
    return
}
*/

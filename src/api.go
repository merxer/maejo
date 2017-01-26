package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type User struct {
	Firstname string
	Lastname string
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func get_users(c echo.Context) error {
	pat := User{
		"thawatchai",
		"singngam",
	}
	return c.JSON(http.StatusOK, pat)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", index)
	e.GET("/users", get_users)

	e.Logger.Fatal(e.Start(":1323"))
}

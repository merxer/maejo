package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname string	`json:"lastname"`
	Username string	`json:"username,omitempty"`
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func get_users(c echo.Context) error {
	pat := User{
		Firstname: "thawatchai",
		Lastname: "singngam",
	}
	return c.JSON(http.StatusOK, pat)
}

func get_users_id(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", index)
	e.GET("/users", get_users)
	e.GET("/users/:id", get_users_id)

	e.Logger.Fatal(e.Start(":1323"))
}

package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func get_users(c echo.Context) error {
	return c.XML(http.StatusOK, "pat")
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", index)
	e.GET("/users", get_users)

	e.Logger.Fatal(e.Start(":1323"))
}

package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()
	e.GET("/", index)
	e.Logger.Fatal(e.Start(":1323"))
}

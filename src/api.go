package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/mgo.v2"

	"./helper"
	db "./helper/db"
	"./models"
)

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func get_users(c echo.Context) error {
	user := models.User{}
	result, _ := user.Read_from_db()
	if len(result) == 0 {
		return c.NoContent(http.StatusNoContent)
	}
	return c.JSON(http.StatusOK, result)
}

func get_users_id(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func create_user(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	user.Save_to_db()
	return c.NoContent(http.StatusCreated)
}

func init() {
	mongo_session, err := mgo.Dial("localhost:27017")
	helper.Check(err)

	mongo_session.SetMode(mgo.Monotonic, true)
	db.Mongo_session = mongo_session
	db.Users_collection = mongo_session.DB("maejo").C("users")
}

func main() {
	defer db.Mongo_session.Close()

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", index)
	e.GET("/users", get_users)
	e.GET("/users/:id", get_users_id)
	e.POST("/users", create_user)

	e.Logger.Fatal(e.Start(":1323"))
}

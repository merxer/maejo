package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"./helper"
	db "./helper/db"
	"./models"
)

const (
	API_SERVER          = ":1323"
	DATABASE_SERVER     = "localhost:27017"
	DATABASE_NAME       = "maejo"
	DATABASE_COLLECTION = "users"
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
	user := new(models.User)
	id := c.Param("id")
	user.Id = bson.ObjectIdHex(id)
	result, _ := user.Read_by_id()
	return c.JSON(http.StatusOK, result)
}

func delete_user_by_id(c echo.Context) error {
	user := new(models.User)
	id := c.Param("id")
	user.Id = bson.ObjectIdHex(id)
	user.Delete_by_id()
	return c.NoContent(http.StatusOK)
}

func create_user(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	user.Save_to_db()
	return c.NoContent(http.StatusCreated)
}

func delete_user_by_keys(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	user.Delete_by_keys()
	return c.NoContent(http.StatusOK)
}

func update_user_by_id(c echo.Context) error {
	user := new(models.User)
	id := c.Param("id")
	user.Id = bson.ObjectIdHex(id)

	if err := c.Bind(user); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	user.Update_by_id()
	return c.NoContent(http.StatusOK)
}


func init() {
	mongo_session, err := mgo.Dial(DATABASE_SERVER)
	helper.Check(err)

	mongo_session.SetMode(mgo.Monotonic, true)
	db.Mongo_session = mongo_session
	db.Users_collection = mongo_session.DB(DATABASE_NAME).C(DATABASE_COLLECTION)
}

func main() {
	defer db.Mongo_session.Close()

	e := echo.New()
	e.Use(middleware.Logger())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/", index)
	e.GET("/users", get_users)
	e.GET("/users/:id", get_users_id)
	e.POST("/users", create_user)
	e.DELETE("/users/:id", delete_user_by_id)
	e.DELETE("/users", delete_user_by_keys)
	e.PUT("/users/:id", update_user_by_id)


	e.Logger.Fatal(e.Start(API_SERVER))
}

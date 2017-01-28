package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/mgo.v2"

	"./helper"
)

var (
	MongoSession    *mgo.Session
	UsersCollection *mgo.Collection
)

type User struct {
	Firstname string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Username  string `json:"username,omitempty" bson:"username,omitempty"`
	Password  string `json:"password,omitempty" bson:"password,omitempty"`
}

func (u *User) save_to_db() error {
	err := UsersCollection.Insert(&u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) read_from_db() ([]User, error) {
	result := []User{}
	err := UsersCollection.Find(nil).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func get_users(c echo.Context) error {
	user := User{}
	result, _ := user.read_from_db()
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
	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	user.save_to_db()
	return c.NoContent(http.StatusCreated)
}

func init() {
	MongoSession, err := mgo.Dial("localhost:27017")
	helper.Check(err)

	MongoSession.SetMode(mgo.Monotonic, true)
	UsersCollection = MongoSession.DB("maejo").C("users")
}

func main() {
	defer MongoSession.Close()

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", index)
	e.GET("/users", get_users)
	e.GET("/users/:id", get_users_id)
	e.POST("/users", create_user)

	e.Logger.Fatal(e.Start(":1323"))
}

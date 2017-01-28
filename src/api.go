package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/mgo.v2"
)

type User struct {
	Firstname string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname string	`json:"lastname,omitempty" bson:"lastname,omitempty"`
	Username string	`json:"username,omitempty" bson:"username,omitempty"`
	Password string	`json:"password,omitempty" bson:"password,omitempty"`
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

func create_user(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, user)
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("maejo").C("users")
	err = c.Insert(&User{
		"thawatchai",
		"singngam",
		"merxer",
		"passw0rd",
	})

	if err != nil {
		panic(err)
	}


	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", index)
	e.GET("/users", get_users)
	e.GET("/users/:id", get_users_id)
	e.POST("/users", create_user)

	e.Logger.Fatal(e.Start(":1323"))
}

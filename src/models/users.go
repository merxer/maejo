package models

import (
	"gopkg.in/mgo.v2/bson"

	db "../helper/db"
)

type User struct {
	Id bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Firstname string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Username  string `json:"username,omitempty" bson:"username,omitempty"`
	Password  string `json:"password,omitempty" bson:"password,omitempty"`
}

func (u *User) Save_to_db() error {
	err := db.Users_collection.Insert(&u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Read_from_db() ([]User, error) {
	result := []User{}
	err := db.Users_collection.Find(nil).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
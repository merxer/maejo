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

type Query struct {
	Key User	`json:"key" bson:"key"`
	Change User `json:"change" bson:"change"`
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

func (u *User) Read_by_id() (*User, error) {
	err := db.Users_collection.Find(bson.M{"_id": u.Id}).One(&u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) Delete_by_id() (*User, error) {
	err := db.Users_collection.RemoveId(u.Id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) Delete_by_keys() (*User, error) {
	err := db.Users_collection.Remove(&u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) Update_by_id() (*User, error) {
	change := bson.M{"$set": &u}
	err := db.Users_collection.UpdateId(u.Id, change)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (q *Query) Update_by_keys() error {
	err := db.Users_collection.Update(&q.Key, &q.Change)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) IsNotDuplicate() bool {
	err := db.Users_collection.Find(bson.M{"username": u.Username}).One(&u)
	if err != nil {
		return true
	}
	return false
}

func (u *User) Login() (*User, error){
	err := db.Users_collection.Find(bson.M{
				"username": u.Username,"password": u.Password}).One(&u)
	if err != nil {
		return nil, err
	}
	return u,nil
}

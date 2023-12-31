package service

import (
	"errors"
	"goseed/models/db"
	"goseed/models/entity"

	"github.com/goonode/mogo"
	"labix.org/v2/mgo/bson"
)

//Userservice is to handle user relation db query
type Userservice struct{}

//Create is to register new user
func (userservice Userservice) Create(user *(entity.User)) error {
	conn := db.GetConnection()
	defer conn.Session.Close()

	doc := mogo.NewDoc(entity.User{}).(*(entity.User))
	err := doc.FindOne(bson.M{"email": user.Email}, doc)
	if err == nil {
		return errors.New("Already Exist")
	}
	userModel := mogo.NewDoc(user).(*(entity.User))
	err = mogo.Save(userModel)
	if vErr, ok := err.(*mogo.ValidationError); ok {
		return vErr
	}
	return err
}

// Delete a user from DB
func (userservice Userservice) Delete(email string) error {
	user, _ := userservice.FindByEmail(email)
	conn := db.GetConnection()
	defer conn.Session.Close()
	err := user.Remove()
	return err
}

//Find user
func (userservice Userservice) Find(user *(entity.User)) (*entity.User, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	doc := mogo.NewDoc(entity.User{}).(*(entity.User))
	err := doc.FindOne(bson.M{"email": user.Email}, doc)

	if err != nil {
		return nil, err
	}
	return doc, nil
}



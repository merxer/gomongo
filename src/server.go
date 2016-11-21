package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"./models"
)

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return session
}

func addUser(mongoSession *mgo.Session) {
	u := models.User{bson.NewObjectId(),"nu", "Pitsanulok","P@ssw0rd"}
	err := mongoSession.DB("test").C("users").Insert(u)
	if err != nil {
		panic(err)
	}
}

func readUsers(mongoSession *mgo.Session) {
	u := []models.User{}
	err := mongoSession.DB("test").C("users").Find(bson.M{"firstname": "nu"}).All(&u)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s,%s\n", u[0].Firstname, u[0].Lastname)
}

func updateUser(mongoSession *mgo.Session) {
	err := mongoSession.DB("test").C("users").Update(bson.M{"firstname": "nu"}, bson.M{"$set": bson.M{"password": "1234567890"}})
	if err != nil {
		panic(err)
	}
}

func deleteUser(mongoSession *mgo.Session) {
	err := mongoSession.DB("test").C("users").Remove(bson.M{"firstname": "nu"})
	if err != nil {
		panic(err)
	}
}


func main() {
	mongoSession := getSession()
	mongoSession.SetMode(mgo.Monotonic, true)

	addUser(mongoSession)
	readUsers(mongoSession)
	updateUser(mongoSession)
	deleteUser(mongoSession)
}

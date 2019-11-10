package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type person struct {
	Name  string
	Phone string
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")

	err = c.Insert(&person{"alex", "+6282358074"}, &person{"Ibob", "+628539357"})

	if err != nil {
		log.Fatal(err)
	}

	result := person{}

	err = c.Find(bson.M{"name": "alex"}).One(&result)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone : ", result.Phone)
}

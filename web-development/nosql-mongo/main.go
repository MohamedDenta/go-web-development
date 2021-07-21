package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

func main() {

	session, err := mgo.Dial(fmt.Sprint("mongodb+srv://<username>:<password>@cluster0.gzooj.mongodb.net/myFirstDatabase"))
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// Optional. Switch the session to a montonic behavior .
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Person{"Ale", "+201234567890"}, &Person{"wael", "+201234567890"}, &Person{"err", "+201234567890"})
	if err != nil {
		log.Fatal(err)
	}
	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
}

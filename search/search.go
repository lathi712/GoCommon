package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//User type
type User struct {
	Name string
	Role string
}

func search(searchQuery string) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	d := session.DB("users").C("user")
	res := searchQueryNew(searchQuery, d)
	if res != nil {
		for i := 0; i < len(res); i++ {
			fmt.Println(res[i].Name)
		}
	} else {
		searchNotexist := fmt.Sprintf("No User exits with %s username", searchQuery)
		fmt.Println(searchNotexist)
	}
}

func searchQueryNew(query string, d *mgo.Collection) []User {

	var res []User
	err := d.Find(bson.M{"name": bson.RegEx{query + `+`, ""}}).All(&res)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
func main() {
	var i string
	//_, err := fmt.Scanf("%s", &i)
	i = os.Args[1]
	//fmt.Print(i)
	var userName = i
	search(userName)

}

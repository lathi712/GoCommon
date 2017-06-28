package main

import (
	"fmt"
	"log"
	"net/url"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func urlMaker() {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("users").C("deployment")

	var deps []Deploy
	err = c.Find(bson.M{}).All(&deps)
	if err != nil {
		//ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
		log.Println("Failed get all books: ", err)
		return
	}

	//	respBody, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(respBody)
	fmt.Println(deps[0].Depname)
	fmt.Println(len(deps))
	//u.Scheme = "http"
	//u.Host = "10.207.139.17:4200"
	//u.Opaque = "10.207.139.17:4200/charts"

	u, err := url.Parse("http://10.207.139.17:4000/charts?")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(deps); i++ {

		q := u.Query()
		q.Set("depid", deps[i].Depid)
		q.Set("depname", deps[i].Depname)
		q.Set("projname", deps[i].Projname)
		q.Set("username", deps[i].Username)
		q.Set("gitcodeurl", deps[i].Gitcodeurl)
		q.Set("namespace", deps[i].Namespace)
		q.Set("nginxpath", deps[i].Nginxpath)
		q.Set("automataenv", deps[i].Automataenv)
		q.Set("subenv", deps[i].Subenv)
		q.Set("kubedepapplabel", deps[i].Kubedepapplabel)
		q.Set("kubeservname", deps[i].Kubeservname)
		q.Set("clustername", deps[i].Clustername)
		q.Set("zone", deps[i].Zone)
		q.Set("createdat", deps[i].Createdat)
		u.RawQuery = q.Encode()
		fmt.Println()
		fmt.Println(u)
	}

}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"goji.io"
	"goji.io/pat"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message: %q}", message)
}

func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}

func server() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	ensureIndex(session)

	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/books"), allBooks(session))
	// mux.HandleFunc(pat.Post("/books"), addBook(session))
	// mux.HandleFunc(pat.Get("/books/:isbn"), bookByISBN(session))
	// mux.HandleFunc(pat.Put("/books/:isbn"), updateBook(session))
	// mux.HandleFunc(pat.Delete("/books/:isbn"), deleteBook(session))

	http.ListenAndServe("localhost:8081", mux)
}

func ensureIndex(s *mgo.Session) {
	session := s.Copy()
	defer session.Close()

	c := session.DB("store").C("books")

	index := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

func allBooks(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		c := session.DB("users").C("deployment")

		var books []Deploy
		err := c.Find(bson.M{}).All(&books)
		if err != nil {
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed get all books: ", err)
			return
		}

		respBody, err := json.MarshalIndent(books, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		ResponseWithJSON(w, respBody, http.StatusOK)
	}
}

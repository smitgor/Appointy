package main

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homepage)
	uc := controllers.NewUserController(getSession())
	http.hendle("/users", uc.CreateUser)
	http.hendle("/users/:id", uc.GetUser)
	http.hendle("/posts", uc.CreatePost)
	http.hendle("/posts/:id", uc.GetPost)
	http.hendle("/posts/users/:id", uc.GetListOfPost)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhos:27017")
	if err != nil {
		panic(err)
	}
	return s
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/smitgor/Appointy/controllers"
	"gopkg.in/mgo.v2"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	uc := controllers.NewUserController(getSession())
	http.HandleFunc("/", homepage)
	http.HandleFunc("/users", uc.CreateUser)
	http.HandleFunc("/users/:id", uc.GetUser)
	//http.HandleFunc("/posts", uc.CreatePost)
	//http.HandleFunc("/posts/:id", uc.GetPost)
	//http.HandleFunc("/posts/users/:id", uc.GetListOfPost)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}

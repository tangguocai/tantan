package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	. "tantan/dbhelper"
	"tantan/handle"
	"tantan/model"
)
var (
	hostname   string
	port       int
)

func init(){
	PgInit()
	user := &model.UserDetail{}
	models := []interface{}{user}
	if err := Pg.CreateTabel(models);err != nil{
		panic(err)
	}
	flag.StringVar(&hostname, "hostname", "0.0.0.0", "The hostname or IP on which the REST server will listen")
	flag.IntVar(&port, "port", 8080, "The port on which the REST server will listen")
}

func main(){
	defer func() {
		Pg.Conn.Close()  // close connect after down
	}()

	flag.Parse()
	var address = fmt.Sprintf("%s:%d", hostname, port)
	log.Println("REST service listening on", address)

	// register router
	router := mux.NewRouter().StrictSlash(true)
	router.
		HandleFunc("/api/service", handle.MyGetHandler).
		Methods("GET")
	router.
		HandleFunc("/api/service/{servicename}",handle.MyPostHandler).
		Methods("POST")
	router.
		HandleFunc("/users",handle.CreateUserHandler).
		Methods("POST")
	router.
		HandleFunc("/users",handle.ListAllUsersHandler).
		Methods("GET")
	router.
		HandleFunc("/users/{user_id}/relationship/{other_user_id}",handle.RelationShipHandler).
		Methods("PUT")
	router.
		HandleFunc("/users/{user_id}/relationship",handle.ListUserRelationShipHandler).
		Methods("GET")

	// start server listening
	err := http.ListenAndServe(address, router)
	if err != nil {
		log.Fatalln("ListenAndServe err:", err)
	}

	log.Println("Server end")
}

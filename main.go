package main

import (
	"github.com/gorilla/mux"
	"labix.org/v2/mgo"
	"log"
	"net/http"
	"time"
)

const MongoConnectionHost  = "127.0.0.1"
const MongoConnectionPort  = 27017
const MongoConnectionDb    = "ggbuilds"

const BuildsCollectionName = "builds"

var (
	mongoSession  *mgo.Session
	mongoDatabase *mgo.Database
	repo           buildRepo

	router = mux.NewRouter()
)

func main() {
	var err error

	// Set up mongo database
	if mongoSession, err = mgo.Dial(MongoConnectionHost); err != nil {
		panic(err)
	}
	log.Println("Connected to MongoDB server")

	// Choose database
	mongoDatabase = mongoSession.DB(MongoConnectionDb)

	// Set up repo collections
	repo.Collection = mongoDatabase.C(BuildsCollectionName)

	// Set up web server handlers
	// START OMIT
	route("/builds", handleBuilds).Methods("GET")
	route("/", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "./index.html")	
	})
	// END OMIT

	http.Handle("/", router)

	log.Printf("Starting GGBuilds API Server")
	panic(http.ListenAndServe(":8080", nil))
}

func route(pattern string, handler func(http.ResponseWriter, *http.Request)) *mux.Route {
	handler = logRequest(handler)
	return router.HandleFunc(pattern, handler)
}

func logRequest(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var s = time.Now()
		handler(w, r)
		log.Printf("%s %s %6.3fms", r.Method, r.RequestURI, (time.Since(s).Seconds()*1000))
	}
}
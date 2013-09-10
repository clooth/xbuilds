// XBuilds
package main

import (
	"flag"
	"fmt"
	"time"
	"net/http"
	"encoding/json"
	"log"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"github.com/gorilla/mux"
  	"github.com/wsxiaoys/terminal/color"
)

//
// Constants
//

const MongoConnectionHost      = "mongodb://xbuilds:foobar987@paulo.mongohq.com:10020/app18043605"
const MongoConnectionDb        = "app18043605"
const BuildsCollectionName     = "builds"
const BuildStepsCollectionName = "buildSteps"


//
// Declared Variables
//

var (
	// Mongo session
	mongoSession  *mgo.Session
	mongoDatabase *mgo.Database

	// Repositories
	buildsRepo     BuildsRepository

	// Router
	router = mux.NewRouter()
)

//
// Utility Functions
//

// Response is a simple way to return JSON objects
type Response map[string]interface{}
func (r Response) String() (s string) {
        b, err := json.Marshal(r)
        if err != nil {
                s = ""
                return
        }
        s = string(b)
        return
}

// WriteJson will wrap and serialize the given object into a JSON object
func RespondWithJSON(rw http.ResponseWriter, v interface{}) {
	response := Response{}

	if _, err := json.Marshal(v); err != nil {
		log.Printf("Error when giving JSON response: %v", err)
		response["success"] = false
		response["data"] = nil
	} else {
		response["success"] = true
		response["data"] = v
	}

	//rw.Header().Set("Content-Length", strconv.Itoa(len(response)))
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, response)
	return
}

// Simplified mux route creator with logging
func route(pattern string, handler func(http.ResponseWriter, *http.Request)) *mux.Route {
	handler = logRequest(handler)
	return router.HandleFunc(pattern, handler)
}

// Logging for requests
func logRequest(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var s = time.Now()
		
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Origin")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == "OPTIONS" && len(r.Header["Origin"]) > 0 {
			return
		}

		handler(w, r)
		color.Printf("@c%s %s %s %6.3fms\n", r.Method, r.RequestURI, r.RemoteAddr, (time.Since(s).Seconds()*1000))
	}
}

//
// Model Objects
//

type (
	Builds []Build
	Build  struct {
		Id      bson.ObjectId `json:"id"          bson:"_id"`
		Name 	string        `json:"n"           bson:"n"`
		Created time.Time     `json:"c"           bson:"c"`
		Updated time.Time     `json:"u,omitempty" bson:"u,omitempty"`
	}
)


//
// Model Repositories
//

type (
	BuildsRepository struct {
		Collection *mgo.Collection
	}
)

func (r BuildsRepository) All() (builds Builds, err error) {
	err = r.Collection.Find(bson.M{}).All(&builds)
	
	if err != nil {
		return nil, err
	}

	return builds, nil
}

func (r BuildsRepository) Get(id string) (build Build, err error) {
	bid := bson.ObjectIdHex(id)
	result := Build{}
	
	err = r.Collection.FindId(bid).One(&result)
	
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r BuildsRepository) Create(build *Build) (err error) {
	if build.Id.Hex() == "" {
		build.Id = bson.NewObjectId()
	}
	
	if build.Created.IsZero() {
		build.Created = time.Now()
	}
	
	build.Updated = time.Now()

	_, err = r.Collection.UpsertId(build.Id, build)
	
	return err
}

func (r BuildsRepository) Update(build *Build) (err error) {
	var change = mgo.Change{
		ReturnNew: true,
		Update: bson.M{
			"$set": bson.M{
				"u": time.Now(),
				"n": build.Name,
			}}}

	_, err = r.Collection.FindId(build.Id).Apply(change, build)

	return err
}

//
// HTTP Handlers
//

func handleBuildsIndex(rw http.ResponseWriter, req *http.Request) {
	var (
		builds Builds
		err    error
	)

	if builds, err = buildsRepo.All(); err != nil {
		log.Printf("Builds index error: %v", err)
		http.Error(rw, "500 Internal server Error", 500)
		return
	}

	RespondWithJSON(rw, builds)
}

//
// Main Application Start Point
//

func main() {
	var err error

	introMessage()

	// Parse command line
	var port string
	flag.StringVar(&port, "p", "1337", "Port to run the server on")
	flag.Parse()

	// Set up mongo database
	if mongoSession, err = mgo.Dial(MongoConnectionHost); err != nil {
		panic(err)
	}
	color.Println("")
	color.Println("@g- Connected to Database")

	// Choose database
	mongoDatabase = mongoSession.DB(MongoConnectionDb)

	// Set up repo collections
	buildsRepo.Collection = mongoDatabase.C(BuildsCollectionName)

	// Set up web server handlers

	//route("/builds/{id}", handleBuild).Methods("GET")
	route("/builds", handleBuildsIndex).Methods("GET")
	route("/", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "./index.html")	
	})
	http.Handle("/", router)

	color.Println("@g- Initialized Router")

	color.Println("@g- Server Online and Listening")
	color.Println("")
	panic(http.ListenAndServe(":"+port, nil))
}

func introMessage() {
	intro := `@m__  ______  _   _ ___ _     ____  ____  
\ \/ / __ )| | | |_ _| |   |  _ \/ ___| 
 \  /|  _ \| | | || || |   | | | \___ \ 
 /  \| |_) | |_| || || |___| |_| |___) |
/_/\_\____/ \___/|___|_____|____/|____/ 
`

	color.Println(intro)
}

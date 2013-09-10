package main

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type (
	Builds []Build
	Build  struct {
		Id      bson.ObjectId `json:"id"          bson:"_id"`
		Name 	string        `json:"n"           bson:"n"`
		Created time.Time     `json:"c"           bson:"c"`
		Updated time.Time     `json:"u,omitempty" bson:"u,omitempty"`
	}
)
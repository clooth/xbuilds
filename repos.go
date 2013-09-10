package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type (
	buildRepo struct {
		Collection *mgo.Collection
	}
)

func (r buildRepo) All() (builds Builds, err error) {
	err = r.Collection.Find(bson.M{}).All(&builds)
	
	if err != nil {
		return nil, err
	}

	return builds, nil
}

func (r buildRepo) Create(build *Build) (err error) {
	if build.Id.Hex() == "" {
		build.Id = bson.NewObjectId()
	}
	
	if build.Created.IsZero() {
		build.Created = time.Now()
	}
	
	build.Updated = time.Now()

	_, err = r.Collection.UpsertId(build.Id, build)
	return
}

func (r buildRepo) Update(build *Build) (err error) {
	var change = mgo.Change{
		ReturnNew: true,
		Update: bson.M{
			"$set": bson.M{
				"u": time.Now(),
				"n": build.Name,
			}}}

	_, err = r.Collection.FindId(build.Id).Apply(change, build)

	// Missing return here
}
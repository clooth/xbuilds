package main

import (
	"log"
	"net/http"
)

// Builds handlers
func handleBuilds(writer http.ResponseWriter, r *http.Request) {
	var (
		builds Builds
		err    error
	)

	if builds, err = repo.All(); err != nil {
		log.Printf("%v", err)
		http.Error(writer, "500 Internal Server Error", 500)
		return
	}

	writeJson(writer, builds)
}

// TODO: Players handlers
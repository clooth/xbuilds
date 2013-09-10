package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func writeJson(rw http.ResponseWriter, v interface{}) bool {
	// avoid json vulnerabilities, always wrap v in an object literal
	doc := map[string]interface{}{"d": v}

	if data, err := json.Marshal(doc); err != nil {
		log.Printf("Error marshalling json: %v", err)
	} else {
		rw.Header().Set("Content-Length", strconv.Itoa(len(data)))
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(data)
	}

	return true
}
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"./../point"
)

//TraceProcessorHandler reads the data received
func TraceProcessorHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var sq point.Point
	err = json.Unmarshal(data, &sq)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

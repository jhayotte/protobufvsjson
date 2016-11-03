package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"../point"
	pb "../pointpb"
	"github.com/golang/protobuf/proto"
)

//TraceHandler shows what is received
func TraceHandler(w http.ResponseWriter, r *http.Request) {
	//symbol := r.FormValue("symbol")
	location := point.Location()

	data, err := json.MarshalIndent(&location, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

//TraceProcessorHandler reads data received
func TraceProcessorHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var sq pb.Location
	err = proto.Unmarshal(data, &sq)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"time"

	"./../point"
	pb "./../pointpb"
	"github.com/golang/protobuf/proto"
)

func protoBody() ([]byte, error) {
	var sq pb.Location

	for i := 0; i <= 1000; i++ {
		sq.Points = append(sq.Points, &pb.Point{
			AssetID:  "VehicleX",
			Lat:      1.22215,
			Lng:      1.223,
			Datetime: time.Now().String(),
		})
	}

	return proto.Marshal(&sq)
}

func jsonBody() ([]byte, error) {
	var sq point.Trace

	for i := 0; i <= 1000; i++ {
		sq.Points = append(sq.Points, &point.Point{
			AssetID:  "VehicleX",
			Lat:      1.22215,
			Lng:      1.223,
			Datetime: time.Now(),
		})
	}

	return json.Marshal(&sq)
}

func main() {
	log.SetFlags(0)

	format := flag.String("format", "json", "The encoding to use. json or proto")
	duration := flag.Int("duration", 60, "Duration to run in seconds.")
	flag.Parse()

	var data []byte
	var err error

	switch *format {
	case "json":
		data, err = jsonBody()
	case "proto":
		data, err = protoBody()
	}

	if err != nil {
		log.Fatalf("Failed to create request body: %v", err)
	}

	log.Printf("Sending requests for %d seconds using %s encoding. Request size: %d\n", *duration, *format, len(data))

	requestTotal := 0
	go func() {
		for {
			resp, err := http.Post("http://127.0.0.1:10000/trace", "", bytes.NewReader(data))
			if err != nil {
				log.Fatalf("Failed http request: %v", err)
			}
			if resp.StatusCode != http.StatusOK {
				log.Fatalf("Error response code: %v", resp.StatusCode)
			}
			requestTotal++
		}
	}()

	time.Sleep(time.Duration(*duration) * time.Second)
	log.Printf("Request Total: %d\n", requestTotal)
}

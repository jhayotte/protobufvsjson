package point

import "time"

//Trace represents a serie of point
type Trace struct {
	Points []*Point
}

//Point represents a physical latitude and longitude
type Point struct {
	AssetID  string
	Lat      float64
	Lng      float64
	Datetime time.Time
}

//Location returns a point for an AssetID at a specific time
func Location() Point {
	p := Point{
		AssetID:  "VehicleX",
		Lat:      1.22215,
		Lng:      1.223,
		Datetime: time.Now(),
	}

	return p
}

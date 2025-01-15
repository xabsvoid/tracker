package model

type LatLng struct {
	Latitude  float64
	Longitude float64
}

func NewLatLng(latitude, longitude float64) LatLng {
	return LatLng{
		Latitude:  latitude,
		Longitude: longitude,
	}
}

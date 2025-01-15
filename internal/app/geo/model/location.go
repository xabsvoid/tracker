package model

import "time"

type Location struct {
	UUID     UUID
	LatLng   LatLng
	DeadLine time.Time
}

func NewLocation(
	uuid UUID,
	latLng LatLng,
	deadLine time.Time,
) Location {
	return Location{
		UUID:     uuid,
		LatLng:   latLng,
		DeadLine: deadLine,
	}
}

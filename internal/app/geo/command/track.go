package command

import (
	"context"
	"time"

	"github.com/xabsvoid/tracker/internal/app/geo/model"
	"github.com/xabsvoid/tracker/internal/app/geo/repository"
)

type Track struct {
	// dependencies
	locationRepo repository.Location
	// request
	uuid     model.UUID
	latLng   model.LatLng
	deadline time.Time
}

func NewTrack(
	locationRepo repository.Location,
	uuid model.UUID,
	latLng model.LatLng,
	deadline time.Time,
) *Track {
	return &Track{
		locationRepo: locationRepo,
		uuid:         uuid,
		latLng:       latLng,
		deadline:     deadline,
	}
}

func (c *Track) Do(ctx context.Context) error {
	location := model.NewLocation(c.uuid, c.latLng, c.deadline)

	return c.locationRepo.Set(ctx, location)
}

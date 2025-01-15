package command

import (
	"context"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/xabsvoid/tracker/internal/app/geo/model"
	"github.com/xabsvoid/tracker/internal/app/geo/repository"
)

type Track struct {
	clock        clock.Clock
	locationRepo repository.Location
	uuid         model.UUID
	latLng       model.LatLng
	ttl          time.Duration
}

func NewTrack(
	clock clock.Clock,
	locationRepo repository.Location,
	uuid model.UUID,
	latLng model.LatLng,
	ttl time.Duration,
) *Track {
	return &Track{
		clock:        clock,
		locationRepo: locationRepo,
		uuid:         uuid,
		latLng:       latLng,
		ttl:          ttl,
	}
}

func (c *Track) Do(ctx context.Context) error {
	deadline := c.clock.Now().Add(c.ttl)
	location := model.NewLocation(c.uuid, c.latLng, deadline)

	return c.locationRepo.Set(ctx, location)
}

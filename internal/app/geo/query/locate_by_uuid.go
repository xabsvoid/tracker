package query

import (
	"context"

	"github.com/benbjohnson/clock"
	"github.com/xabsvoid/tracker/internal/app/geo/model"
	"github.com/xabsvoid/tracker/internal/app/geo/repository"
)

type LocateByUUID struct {
	clock        clock.Clock
	locationRepo repository.Location
	uuid         model.UUID

	location model.Location
}

func NewLocateByUUID(
	clock clock.Clock,
	locationRepo repository.Location,
	uuid model.UUID,
) *LocateByUUID {
	return &LocateByUUID{
		clock:        clock,
		locationRepo: locationRepo,
		uuid:         uuid,
	}
}

func (q *LocateByUUID) GetLocation() model.Location {
	return q.location
}

func (q *LocateByUUID) Do(ctx context.Context) error {
	location, err := q.locationRepo.GetByUUID(ctx, q.uuid)
	if err != nil {
		return err
	}

	timeNow := q.clock.Now()
	if !timeNow.Before(location.DeadLine) {
		return model.ErrNotFound
	}

	q.location = location

	return nil
}

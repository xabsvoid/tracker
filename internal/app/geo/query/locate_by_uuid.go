package query

import (
	"context"
	"time"

	"github.com/xabsvoid/tracker/internal/app/geo/model"
	"github.com/xabsvoid/tracker/internal/app/geo/repository"
)

type LocateByUUID struct {
	// dependencies
	locationRepo repository.Location
	// request
	uuid     model.UUID
	deadline time.Time
	// response
	location model.Location
}

func NewLocateByUUID(
	locationRepo repository.Location,
	uuid model.UUID,
	deadline time.Time,
) *LocateByUUID {
	return &LocateByUUID{
		locationRepo: locationRepo,
		uuid:         uuid,
		deadline:     deadline,
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

	if !q.deadline.Before(location.DeadLine) {
		return model.ErrNotFound
	}

	q.location = location

	return nil
}

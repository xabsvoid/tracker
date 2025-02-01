package repository

import (
	"context"
	"sync"
	"time"

	"github.com/xabsvoid/tracker/internal/app/geo/model"
)

type Location struct {
	limit       int
	mutex       sync.RWMutex
	locationMap map[string]model.Location
}

func NewLocation(limit int) *Location {
	return &Location{
		limit:       limit,
		locationMap: make(map[string]model.Location, limit),
	}
}

func (r *Location) GetByUUID(_ context.Context, uuid model.UUID) (model.Location, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	location, ok := r.locationMap[string(uuid)]
	if !ok {
		return model.Location{}, model.ErrNotFound
	}

	return location, nil
}

func (r *Location) Set(_ context.Context, location model.Location) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, ok := r.locationMap[string(location.UUID)]
	if !ok && len(r.locationMap) >= r.limit {
		return model.ErrOverflow
	}

	r.locationMap[string(location.UUID)] = location

	return nil
}

func (r *Location) Cleanup(timeNow time.Time, limit int) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	keys := make([]string, 0, limit)

	for key, location := range r.locationMap {
		if !timeNow.Before(location.DeadLine) {
			keys = append(keys, key)
			if len(keys) == cap(keys) {
				break
			}
		}
	}

	for _, key := range keys {
		delete(r.locationMap, key)
	}
}

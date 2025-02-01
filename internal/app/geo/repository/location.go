package repository

import (
	"context"

	"github.com/xabsvoid/tracker/internal/app/geo/model"
)

type Location interface {
	GetByUUID(ctx context.Context, uuid model.UUID) (model.Location, error)
	Set(ctx context.Context, location model.Location) error
}

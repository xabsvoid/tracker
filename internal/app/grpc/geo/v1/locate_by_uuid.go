package v1

import (
	"context"
	"errors"

	"github.com/xabsvoid/tracker/internal/app/geo/model"
	"github.com/xabsvoid/tracker/internal/app/geo/query"
	geov1 "github.com/xabsvoid/tracker/pkg/api/geo/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) LocateByUUID(ctx context.Context, req *geov1.LocateByUUIDRequest) (*geov1.LocateByUUIDResponse, error) {
	uuid := model.NewUUID(req.GetUuid())
	if err := uuid.Validate(); err != nil {
		return nil, status.New(codes.InvalidArgument, "uuid: "+err.Error()).Err()
	}

	handler := query.NewLocateByUUID(s.clock, s.locationRepo, uuid)

	if err := handler.Do(ctx); err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.New(codes.NotFound, err.Error()).Err()
		}

		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	location := handler.GetLocation()

	return &geov1.LocateByUUIDResponse{
		Latlng: &geov1.LatLng{
			Latitude:  location.LatLng.Latitude,
			Longitude: location.LatLng.Longitude,
		},
		Deadline: timestamppb.New(location.DeadLine),
	}, nil
}

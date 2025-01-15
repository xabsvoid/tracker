package v1

import (
	"context"

	"github.com/xabsvoid/tracker/internal/app/geo/command"
	"github.com/xabsvoid/tracker/internal/app/geo/model"
	geov1 "github.com/xabsvoid/tracker/pkg/api/geo/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) Track(ctx context.Context, req *geov1.TrackRequest) (*geov1.TrackResponse, error) {
	uuid := model.NewUUID(req.GetUuid())
	if err := uuid.Validate(); err != nil {
		return nil, status.New(codes.InvalidArgument, "uuid: "+err.Error()).Err()
	}

	latLngReq := req.GetLatlng()
	latLng := model.NewLatLng(latLngReq.GetLatitude(), latLngReq.GetLongitude())

	handler := command.NewTrack(s.clock, s.locationRepo, uuid, latLng, s.trackTTL)

	if err := handler.Do(ctx); err != nil {
		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	return &geov1.TrackResponse{}, nil
}

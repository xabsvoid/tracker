package v1

import (
	"context"
	geov1 "github.com/xabsvoid/tracker/pkg/api/geo/v1"
	"google.golang.org/grpc"
	"log"
)

type Service struct {
	geov1.UnimplementedGeoServiceServer
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) RegisterService(srv grpc.ServiceRegistrar) {
	geov1.RegisterGeoServiceServer(srv, s)
}

func (s *Service) Track(_ context.Context, req *geov1.TrackRequest) (*geov1.TrackResponse, error) {
	log.Printf("Track: %s\n", req.String())
	return &geov1.TrackResponse{}, nil
}

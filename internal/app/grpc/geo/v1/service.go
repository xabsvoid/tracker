package v1

import (
	"time"

	"github.com/benbjohnson/clock"
	"github.com/xabsvoid/tracker/internal/app/geo/repository"
	geov1 "github.com/xabsvoid/tracker/pkg/api/geo/v1"
	"google.golang.org/grpc"
)

type Service struct {
	geov1.UnimplementedGeoServiceServer

	clock        clock.Clock
	locationRepo repository.Location
	trackTTL     time.Duration
}

func NewService(
	clock clock.Clock,
	locationRepo repository.Location,
	trackTTL time.Duration,
) *Service {
	return &Service{
		clock:        clock,
		locationRepo: locationRepo,
		trackTTL:     trackTTL,
	}
}

func (s *Service) RegisterService(srv grpc.ServiceRegistrar) {
	geov1.RegisterGeoServiceServer(srv, s)
}

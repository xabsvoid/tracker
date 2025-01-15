package main

import (
	"context"
	"log"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/xabsvoid/tracker/internal/app/grpc"
	geov1 "github.com/xabsvoid/tracker/internal/app/grpc/geo/v1"
	mem "github.com/xabsvoid/tracker/internal/app/mem/geo/repository"
)

const (
	LocationRepositoryLimit           = 100
	LocationRepositoryCleanupInterval = 10 * time.Second
	LocationRepositoryCleanupLimit    = 10
	GeoTrackTTL                       = time.Minute
)

func main() {
	ctx := context.Background()

	rtClock := clock.New()

	locationRepo := mem.NewLocation(LocationRepositoryLimit)
	locationRepoTicker := time.NewTicker(LocationRepositoryCleanupInterval)
	defer locationRepoTicker.Stop()
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-locationRepoTicker.C:
				locationRepo.Cleanup(rtClock.Now(), LocationRepositoryCleanupLimit)
			}
		}
	}()

	geoV1Service := geov1.NewService(rtClock, locationRepo, GeoTrackTTL)

	server := grpc.NewServer("tcp", ":50051", geoV1Service)

	if err := server.Run(); err != nil {
		log.Print(err)
	}
}

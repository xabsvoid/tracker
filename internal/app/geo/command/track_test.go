package command

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/xabsvoid/tracker/internal/app/geo/model"
	"github.com/xabsvoid/tracker/internal/pkg/mock"
)

func TestTrack_Do(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	timeNow := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	uuid := model.NewUUID("6ba7b811-9dad-11d1-80b4-00c04fd430c8")

	latLng := model.NewLatLng(1.0, 1.0)

	ttl := time.Hour

	location := model.NewLocation(uuid, latLng, timeNow.Add(ttl))

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mockClock := mock.NewMockClock(t)
		mockClock.EXPECT().Now().Return(timeNow)
		t.Cleanup(func() {
			mockClock.AssertExpectations(t)
		})

		mockLocationRepo := mock.NewMockLocation(t)
		mockLocationRepo.EXPECT().Set(ctx, location).Return(nil)
		t.Cleanup(func() {
			mockLocationRepo.AssertExpectations(t)
		})

		handler := NewTrack(mockClock, mockLocationRepo, uuid, latLng, ttl)
		err := handler.Do(ctx)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		mockClock := mock.NewMockClock(t)
		mockClock.EXPECT().Now().Return(timeNow)
		t.Cleanup(func() {
			mockClock.AssertExpectations(t)
		})

		errTest := errors.New("test error")

		mockLocationRepo := mock.NewMockLocation(t)
		mockLocationRepo.EXPECT().Set(ctx, location).Return(errTest)
		t.Cleanup(func() {
			mockLocationRepo.AssertExpectations(t)
		})

		handler := NewTrack(mockClock, mockLocationRepo, uuid, latLng, ttl)
		err := handler.Do(ctx)
		require.ErrorIs(t, err, errTest)
	})
}

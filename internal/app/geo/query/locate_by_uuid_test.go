package query

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabsvoid/tracker/internal/app/geo/model"
	"github.com/xabsvoid/tracker/internal/pkg/mock"
)

func TestLocateByUUID_Do(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	timeNow := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	uuid := model.NewUUID("6ba7b811-9dad-11d1-80b4-00c04fd430c8")

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		latLng := model.NewLatLng(1.0, 1.0)

		deadline := time.Now().Add(time.Hour)

		location := model.NewLocation(uuid, latLng, deadline)

		mockLocationRepo := mock.NewMockLocation(t)
		mockLocationRepo.EXPECT().GetByUUID(ctx, uuid).Return(location, nil)
		t.Cleanup(func() {
			mockLocationRepo.AssertExpectations(t)
		})

		handler := NewLocateByUUID(mockLocationRepo, uuid, timeNow)
		err := handler.Do(ctx)
		require.NoError(t, err)
		assert.Equal(t, location, handler.GetLocation())
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		errTest := errors.New("test error")

		mockLocationRepo := mock.NewMockLocation(t)
		mockLocationRepo.EXPECT().GetByUUID(ctx, uuid).Return(model.Location{}, errTest)
		t.Cleanup(func() {
			mockLocationRepo.AssertExpectations(t)
		})

		handler := NewLocateByUUID(mockLocationRepo, uuid, timeNow)
		err := handler.Do(ctx)
		require.ErrorIs(t, err, errTest)
	})

	t.Run("deadline", func(t *testing.T) {
		t.Parallel()

		latLng := model.NewLatLng(1.0, 1.0)

		deadline := timeNow.Add(-time.Hour)

		location := model.NewLocation(uuid, latLng, deadline)

		mockLocationRepo := mock.NewMockLocation(t)
		mockLocationRepo.EXPECT().GetByUUID(ctx, uuid).Return(location, nil)
		t.Cleanup(func() {
			mockLocationRepo.AssertExpectations(t)
		})

		handler := NewLocateByUUID(mockLocationRepo, uuid, timeNow)
		err := handler.Do(ctx)
		require.ErrorIs(t, err, model.ErrNotFound)
	})
}

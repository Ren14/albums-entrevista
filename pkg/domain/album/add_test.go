package album

import (
	"github.com/go-playground/assert/v2"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestAdd_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	fakeRepo := NewMockStorageRepository(ctrl)

	service := &Service{
		StorageRepository: fakeRepo,
		CurrencyService:   nil,
	}
	testAlbum := Album{
		ID:     "1",
		Title:  "test",
		Artist: "test",
		Price:  1,
	}

	fakeRepo.EXPECT().Add(gomock.Any()).Return(nil)

	response := service.Add(testAlbum)

	assert.Equal(t, response, nil)
}

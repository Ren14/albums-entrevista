package album

import (
	"github.com/go-playground/assert/v2"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestGetAllUSD_Success(t *testing.T) {
	ctrl := gomock.NewController(t)

	fakeStorage := NewMockStorageRepository(ctrl)
	fakeCurrencyService := NewMockCurrencyService(ctrl)

	service := &Service{
		StorageRepository: fakeStorage,
		CurrencyService:   fakeCurrencyService,
	}

	fakeStorage.EXPECT().GetAll().Return([]Album{
		{
			ID:     "1",
			Title:  "test",
			Artist: "test",
			Price:  2,
		},
		{
			ID:     "2",
			Title:  "test",
			Artist: "test",
			Price:  4,
		},
	})

	fakeCurrencyService.EXPECT().GetUSDPrice().Return(25.50)

	albums := service.GetAllUSD()

	assert.Equal(t, albums, []Album{
		{
			ID:     "1",
			Title:  "test",
			Artist: "test",
			Price:  0.0784313725490196,
		},
		{
			ID:     "2",
			Title:  "test",
			Artist: "test",
			Price:  0.1568627450980392,
		},
	})
}

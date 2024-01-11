package currencies

import (
	"github.com/go-playground/assert/v2"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestGetUSDPrice_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	fakeRepo := NewMockCurrencyRepo(ctrl)
	service := &CurrencyService{
		CurrencyRepo: fakeRepo,
	}

	fakeRepo.EXPECT().GetPrice().Return(100.50)

	price := service.GetUSDPrice()

	assert.Equal(t, price, 100.50)
}

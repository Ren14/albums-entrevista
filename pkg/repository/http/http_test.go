package http

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestGetPrice_Success(t *testing.T) {
	r := &Currency{}
	price := r.GetPrice()
	assert.Equal(t, price, 1110.50)
}

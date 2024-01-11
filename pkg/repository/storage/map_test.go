package storage

import (
	"github.com/go-playground/assert/v2"
	"github.com/rontivero/entrevistatecnica/pkg/domain/album"
	"testing"
)

func TestAdd_GetByID_Success(t *testing.T) {
	fakeMap := &Map{}

	err := fakeMap.Add(album.Album{
		ID:     "1",
		Title:  "test",
		Artist: "test",
		Price:  1,
	})

	assert.Equal(t, err, nil)
}

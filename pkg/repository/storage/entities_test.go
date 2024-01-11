package storage

import (
	"github.com/go-playground/assert/v2"
	"github.com/rontivero/entrevistatecnica/pkg/domain/album"
	"testing"
	"time"
)

func TestFromDomain(t *testing.T) {
	album := fromDomain(album.Album{
		ID:     "1",
		Title:  "test",
		Artist: "test",
		Price:  1,
	})

	assert.Equal(t, album, Album{
		ID:        "1",
		Title:     "test",
		Artist:    "test",
		Price:     1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	})
}

func TestToDomain(t *testing.T) {
	repoAlbum := Album{
		ID:        "1",
		Title:     "test",
		Artist:    "test",
		Price:     1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	domainAlbum := repoAlbum.toDomain()

	assert.Equal(t, domainAlbum, album.Album{
		ID:     "1",
		Title:  "test",
		Artist: "test",
		Price:  1,
	})
}

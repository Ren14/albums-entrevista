package storage

import (
	"github.com/rontivero/entrevistatecnica/pkg/domain/album"
	"time"
)

type Album struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Artist    string    `json:"artist"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// albums slice to seed record album data.
var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

var mapAlbums = map[int]Album{
	1: {
		ID:        "1",
		Title:     "Oktubre",
		Artist:    "Patricio Rey y sus redonditos de ricota",
		Price:     49.56,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	2: {
		ID:        "2",
		Title:     "La mosca y la sopa",
		Artist:    "Patricio Rey y sus redonditos de ricota",
		Price:     50.50,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

func fromDomain(a album.Album) Album {
	return Album{
		ID:        a.ID,
		Title:     a.Title,
		Artist:    a.Artist,
		Price:     a.Price,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

func (a Album) toDomain() *album.Album {
	return &album.Album{
		ID:     a.ID,
		Title:  a.Title,
		Artist: a.Artist,
		Price:  a.Price,
	}
}

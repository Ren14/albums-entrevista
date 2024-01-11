package storage

import (
	"errors"
	"github.com/rontivero/entrevistatecnica/pkg/domain/album"
	"strconv"
)

func (m Map) Add(album album.Album) error {
	a, err := m.GetByID(album.ID)
	if err != nil {
		if err.Error() != "error not found" {
			return err
		}
	}

	if a != nil {
		return nil //manejar idempotencia
	}

	mapAlbums[getIntID(album.ID)] = fromDomain(album)
	return nil
}

func (m Map) GetByID(albumID string) (*album.Album, error) {
	a, ok := mapAlbums[getIntID(albumID)]
	if ok {
		return a.toDomain(), nil
	}

	return nil, errors.New("error not found")
}

func (m Map) GetAll() []album.Album {
	var sliceAlbums []album.Album
	for _, a := range mapAlbums {
		sliceAlbums = append(sliceAlbums, *a.toDomain())
	}
	return sliceAlbums
}

func getIntID(albumID string) int {
	id, err := strconv.Atoi(albumID)
	if err != nil {
		// ... handle error
		panic(err)
	}

	return id
}

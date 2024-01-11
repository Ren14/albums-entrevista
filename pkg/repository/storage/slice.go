package storage

import (
	"errors"
	"github.com/rontivero/entrevistatecnica/pkg/domain/album"
)

func (s *Slice) Add(album album.Album) error {
	a, err := s.GetByID(album.ID)
	if err != nil {
		if err.Error() != "error not found" {
			return err
		}
	}

	if a != nil {
		return nil // manejar idempotencia
	}

	newAlbum := fromDomain(album)
	albums = append(albums, newAlbum)

	return nil
}

func (s *Slice) GetByID(albumID string) (*album.Album, error) {
	for _, a := range albums {
		if a.ID == albumID {
			return a.toDomain(), nil
		}
	}

	return nil, errors.New("error not found")
}

func (s *Slice) GetAll() []album.Album {
	var response []album.Album
	for _, a := range albums {
		response = append(response, *a.toDomain())
	}

	return response
}

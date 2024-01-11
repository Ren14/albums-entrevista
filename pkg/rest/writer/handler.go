package writer

import (
	"github.com/gin-gonic/gin"
	"github.com/rontivero/entrevistatecnica/pkg/domain/album"
	"net/http"
)

type AlbumService interface {
	Add(album album.Album) error
}

type Handler struct {
	AlbumService AlbumService
}

func NewHandler(albumSrv AlbumService) *Handler {
	return &Handler{
		AlbumService: albumSrv,
	}
}

func (h *Handler) postAlbums(c *gin.Context) {
	var newAlbum album.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	err := h.AlbumService.Add(newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error creating album", "err": err.Error()})
	} else {
		c.IndentedJSON(http.StatusCreated, newAlbum)
	}
}

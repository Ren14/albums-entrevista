package reader

import (
	"github.com/gin-gonic/gin"
	"github.com/rontivero/entrevistatecnica/pkg/domain/album"
	"net/http"
)

type AlbumService interface {
	GetAll() []album.Album
	GetByID(albumID string) (*album.Album, error)
	GetAllUSD() []album.Album
}

type Handler struct {
	AlbumService AlbumService
}

func NewHandler(albumSrv AlbumService) *Handler {
	return &Handler{
		AlbumService: albumSrv,
	}
}

func (h *Handler) getAlbums(c *gin.Context) {
	albums := h.AlbumService.GetAll()
	c.IndentedJSON(http.StatusOK, albums)
}

func (h *Handler) getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	album, err := h.AlbumService.GetByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	} else {
		c.IndentedJSON(http.StatusOK, *album)
	}
}

func (h *Handler) getAlbumsPriceInUSD(c *gin.Context) {
	albums := h.AlbumService.GetAllUSD()
	c.IndentedJSON(http.StatusOK, albums)
}

func ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "pong")
}

package reader

import "github.com/gin-gonic/gin"

func (h *Handler) GetRoutes(router *gin.Engine) {
	router.GET("/ping", ping)
	router.GET("/albums", h.getAlbums)
	router.GET("/albums/:id", h.getAlbumByID)
	router.GET("/albums/usd", h.getAlbumsPriceInUSD)
}

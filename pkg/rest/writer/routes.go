package writer

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetRoutes(router *gin.Engine) {
	router.POST("/albums", h.postAlbums)

}

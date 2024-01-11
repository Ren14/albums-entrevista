package writer

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/rontivero/entrevistatecnica/pkg/domain/album"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostAlbums_Success(t *testing.T) {
	ctrl := gomock.NewController(t)

	fakeAlbumService := NewMockAlbumService(ctrl)
	handler := Handler{AlbumService: fakeAlbumService}

	fakeAlbumService.EXPECT().Add(gomock.Any()).Return(nil)
	body := album.Album{
		ID:     "1",
		Title:  "test",
		Artist: "test",
		Price:  1,
	}
	validJsonAlbum, _ := json.Marshal(body)

	request := httptest.NewRequest(http.MethodGet, "https://example.com/api/get/post/", bytes.NewReader(validJsonAlbum))

	g := gin.Context{
		Request:  request,
		Writer:   nil, // TODO: debo revisar como inyectar esta dependencia
		Params:   nil,
		Keys:     nil,
		Errors:   nil,
		Accepted: nil,
	}

	handler.postAlbums(&g)

	assert.Equal(t, g.Writer.Status(), http.StatusCreated)
}

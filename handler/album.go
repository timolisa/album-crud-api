package handler

import (
	"net/http"
	"simba/album-store-api/album"

	"github.com/gin-gonic/gin"
)

type AlbumController struct {
	albumService album.Service
}

func NewAlbumController(albumService album.Service) *AlbumController {
	return &AlbumController{albumService}
}

func (h *AlbumController) GetAllAlbums(c *gin.Context) {
	albums, err := h.albumService.FindAll()

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
	}

	c.IndentedJSON(http.StatusFound, gin.H{"success": true, "message": albums})
}

func (h *AlbumController) GetAlbumById(c *gin.Context) {
	id := c.Param("id")
	album, err := h.albumService.FindByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.IndentedJSON(http.StatusFound, gin.H{"success": true, "payload": album})
}

func (h *AlbumController) CreateAlbum(c *gin.Context) {
	var newAlbum album.AlbumRequest

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	response, err := h.albumService.Create(newAlbum)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "error creating album"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"success": true, "payload": response})
}

func (h *AlbumController) UpdateAlbum(c *gin.Context) {
	var newAlbum album.AlbumRequest
	id := c.Param("id")

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	response, err := h.albumService.Update(id, newAlbum)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "error creating album"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"success": true, "payload": response})
}

func (h *AlbumController) DeleteAlbum(c *gin.Context) {
	id := c.Param("id")
	message := h.albumService.Delete(id)

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": message})

}

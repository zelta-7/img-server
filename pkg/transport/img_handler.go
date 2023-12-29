package transport

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zelta-7/img-server/pkg/consumer"
	"github.com/zelta-7/img-server/pkg/producer"
	"github.com/zelta-7/img-server/pkg/repository"
	"github.com/zelta-7/img-server/pkg/service"
)

type Input struct {
	Url string `json:"url" binding:"required"`
}

type ImgHandler struct {
	imgService service.ImgService
}

func NewImgHandler(imgService service.ImgService) *ImgHandler {
	return &ImgHandler{imgService: imgService}
}

func (h *ImgHandler) GetImg(c *gin.Context) {
	url := c.Param("imageName")

	compressedPath, err := h.imgService.GetImg(url)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "image not found"})
		return
	}

	c.File(compressedPath)
}

func (h *ImgHandler) PostImg(c *gin.Context) {
	url := c.PostForm("url")

	if url == " " {
		fmt.Println("")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "URL IS REQUIRED"})
		return
	}

	err := producer.QueueImage(url)
	if err != nil {
		fmt.Println("Error queuing image: ", err)
	}

	compressedPath, err := consumer.Consume("imageQueue")
	if err != nil {
		fmt.Println("Error consuming image: ", err)
	}

	newId := uint(uuid.New().ID())
	for _, path := range compressedPath {
		record := repository.ImageRecord{
			Id:             newId,
			Url:            url,
			CompressedPath: path,
		}
		err := h.imgService.PostImg(record)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(200, gin.H{"status": "Image queued for processing"})
}

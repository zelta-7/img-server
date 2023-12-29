package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/zelta-7/img-server/pkg/service"
)

func Route(imgService service.ImgService) (*gin.Engine, error) {
	imgHandler := NewImgHandler(imgService)

	r := gin.Default()

	r.POST("/post", imgHandler.PostImg)
	r.GET("/get/:imageName", imgHandler.GetImg)

	r.Run(":8080")

	return r, nil
}

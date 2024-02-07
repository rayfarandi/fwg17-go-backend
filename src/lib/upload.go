package lib

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Upload(c *gin.Context, field string, dest string) *string {
	file, _ := c.FormFile(field)
	fileExt := map[string]string{
		"image/jpg":  ".jpg",
		"image/png":  ".png",
		"image/jpeg": ".jpeg",
	}
	fileType := file.Header["Content-Type"][0]
	log.Println(file.Header["Content-Type"][0])

	fileName := fmt.Sprintf("%v%v", uuid.NewString(), fileExt[fileType])
	fileDes := fmt.Sprintf("uploads/%v%v", dest, fileName)

	c.SaveUploadedFile(file, fileDes)
	return &fileName
}

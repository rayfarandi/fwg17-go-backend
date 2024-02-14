package lib

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Upload(c *gin.Context, field string, dest string) (string, error) {
	file, _ := c.FormFile(field)

	invalidType := "Invalid File Type, only .jpg .png .jpeg"
	invalidSize := "Invalid File Size more than 1.5MB"

	fileExt := map[string]string{
		"image/jpg":  ".jpg",
		"image/png":  ".png",
		"image/jpeg": ".jpeg",
	}
	fileType := file.Header["Content-Type"][0]

	typeExt := fileExt[fileType]

	if typeExt == "" {
		fmt.Println(fileExt)
		return "", errors.New(invalidType)
	}

	if file.Size > 1500000 {
		return "", errors.New(invalidSize)
	}

	fileName := fmt.Sprintf("uploads/%v/%v%v", dest, uuid.NewString(), typeExt)

	c.SaveUploadedFile(file, fileName)
	return fileName, nil
}

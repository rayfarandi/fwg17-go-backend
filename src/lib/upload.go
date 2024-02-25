package lib

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Upload(c *gin.Context, field string, dest string) (string, error) {

	file, err := c.FormFile(field)

	if err != nil {
		fmt.Println(err)
		return "", errors.New("field file not found")
	}

	ext := map[string]string{
		"image/jpeg": ".jpg",
		"image/png":  ".png",
	}

	fileType := file.Header["Content-Type"][0]
	extension := ext[fileType]
	if extension == "" {
		fmt.Println(fileType)
		return "", errors.New("only jpeg, jpg, and png file")
	}

	fileSize := file.Size
	if fileSize > 1500000 {
		fmt.Println(fileSize)
		return "", errors.New("the file size exceeds the maximum limit of 1.5 MB")
	}

	fileName := fmt.Sprintf("uploads/%v/%v%v", dest, uuid.NewString(), extension)

	c.SaveUploadedFile(file, fileName)

	return fileName, nil
}

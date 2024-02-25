package controllers_customer

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/rayfarandi/fwg17-go-backend/src/lib"
	"github.com/rayfarandi/fwg17-go-backend/src/models"
	"github.com/rayfarandi/fwg17-go-backend/src/service"

	"github.com/KEINOS/go-argonize"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	id := int(claims["id"].(float64))

	user, err := models.FindOneUsers(id)
	if err != nil {
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
				Success: false,
				Message: "User not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &service.Response{
		Success: true,
		Message: "Detail user",
		Results: user,
	})
}

func UpdateProfile(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	id := int(claims["id"].(float64))

	isUserExist, error := models.FindOneUsers(id)
	if error != nil {
		fmt.Println(isUserExist, error)
		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: "no data found",
		})
		return
	}

	data := service.UserForm{}
	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	plain := []byte(data.Password)
	hash, _ := argonize.Hash(plain)
	data.Password = hash.String()
	data.Id = id

	_, err = c.FormFile("picture")
	if err == nil {
		_ = os.Remove("./" + isUserExist.Picture)

		file, err := lib.Upload(c, "picture", "users")
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		data.Picture = file
	}

	user, err := models.UpdateUser(data)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &service.Response{
		Success: true,
		Message: "User updated successfully",
		Results: user,
	})
}

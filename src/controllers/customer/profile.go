package controllers_customer

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/KEINOS/go-argonize"
	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/gin-gonic/gin"
	"github.com/rayfarandi/fwg17-go-backend/src/lib"
	"github.com/rayfarandi/fwg17-go-backend/src/models"
	"github.com/rayfarandi/fwg17-go-backend/src/services"
)

func GetProfile(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	id := int(claims["id"].(float64))
	// data := c.MustGet("id").(*services.User)
	// id := data.Id
	// fmt.Print(id)

	user, err := models.FindOneUser(id)
	if err != nil {
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
				Success: false,
				Message: "User not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Detail user",
		Results: user,
	})
}

func UpdateProfile(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	id := int(claims["id"].(float64))
	// costomer : c.MustGet("id").(*services.UserForm)
	// id := costomer.Id

	isUserExist, error := models.FindOneUser(id)
	if error != nil {
		fmt.Println(isUserExist, error)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "no data found",
		})
		return
	}

	data := services.UserForm{}
	err := c.ShouldBind(&data)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	plain := []byte(data.Password)
	hash, err := argonize.Hash(plain)
	if err != nil {
		fmt.Println(err)
		return
	}
	data.Password = hash.String()
	data.Id = id

	// upload
	// _, err = c.FormFile("picture")
	// if err == nil {
	// 	err = os.Remove("./" + isUserExist.Picture)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	file, err := lib.Upload(c, "picture", "users")
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
	// 			Success: false,
	// 			Message: err.Error(),
	// 		})
	// 		return
	// 	}

	// 	data.Picture = file
	// } else {
	// 	fmt.Println(err)
	// 	data.Picture = ""
	// }
	_, err = c.FormFile("picture")
	if err == nil {
		err := os.Remove("./" + isUserExist.Picture)

		file, err := lib.Upload(c, "picture", "users")
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
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
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "User updated successfully",
		Results: user,
	})
}

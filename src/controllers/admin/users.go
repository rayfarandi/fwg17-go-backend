package controllers_admin

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/KEINOS/go-argonize"
	"github.com/gin-gonic/gin"
	"github.com/rayfarandi/fwg17-go-backend/src/lib"
	"github.com/rayfarandi/fwg17-go-backend/src/models"
	"github.com/rayfarandi/fwg17-go-backend/src/services"
)

func ListAllUsers(c *gin.Context) {
	searchKey := c.DefaultQuery("searchKey", "")
	sortBy := c.DefaultQuery("sortBy", "id")
	order := c.DefaultQuery("order", "ASC")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	offset := (page - 1) * limit

	result, err := models.FindAllUsers(searchKey, sortBy, order, limit, offset)

	totalPage := int(math.Ceil(float64(result.Count) / float64(limit)))
	nextPage := page + 1
	if nextPage > totalPage {
		nextPage = 0
	}
	prevPage := page - 1
	if prevPage < 1 {
		prevPage = 0
	}

	pageInfo := services.PageInfo{
		Page:      page,
		Limit:     limit,
		NextPage:  nextPage,
		PrevPage:  prevPage,
		TotalPage: totalPage,
		TotalData: result.Count,
	}
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, &services.ResponseList{
		Success:  true,
		Message:  "List All Users",
		PageInfo: pageInfo,
		Results:  result.Data,
	})
}

func DetailUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := models.FindOneUser(id)
	if err != nil {
		// log.Println(err)
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Detail User",
		Results: user,
	})
}

func CreateUser(c *gin.Context) {
	data := services.UserForm{}

	emailInput := c.PostForm("email")
	passwordInput := c.PostForm("password")

	if emailInput == "" || passwordInput == "" {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Email or Password not be empty",
		})
		return
	}
	exitingEmail, _ := models.FindOneUserEmail(emailInput)
	if exitingEmail.Email == emailInput {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Email already use from data",
		})
		return
	}
	c.ShouldBind(&data)

	//upload
	_, err := c.FormFile("picture")
	if err == nil {
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
	} else {
		data.Picture = ""
	}
	//upload

	plain := []byte(data.Password)
	hash, err := argonize.Hash(plain)
	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Generate password error",
		})
		return
	}
	data.Password = hash.String()

	user, err := models.CreateUser(data)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "User Created successfully",
		Results: user,
	})
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := services.UserForm{}

	data.Id = id

	isExist, err := models.FindOneUser(id)
	if err != nil {
		fmt.Println(isExist, err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "User not found",
		})
		return
	}

	c.ShouldBind(&data)

	//upload
	_, err = c.FormFile("picture")
	if err == nil {
		err := os.Remove("./" + isExist.Picture)
		if err != nil {
		}

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
	//upload
	plain := []byte(data.Password)
	hash, err := argonize.Hash(plain)
	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Failid generate hash",
		})
		return
	}
	data.Password = hash.String()

	user, err := models.UpdateUser(data)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "User update successfully",
		Results: user,
	})
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := models.FindOneUser(id)

	if err != nil {
		fmt.Println(data, err)
		if strings.HasPrefix(err.Error(), "sql: no rows in result set") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "No data",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}
	user, err := models.DeleteUser(id)

	if user.Picture != "" {
		err := os.Remove("./" + user.Picture)
		if err != nil {
			fmt.Println("Error deleting file:", err)
		}
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Delete User",
		Results: user,
	})
}

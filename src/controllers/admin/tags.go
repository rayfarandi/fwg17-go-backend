package controllers_admin

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rayfarandi/fwg17-go-backend/src/models"
	"github.com/rayfarandi/fwg17-go-backend/src/services"
)

func ListAllTags(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	offset := (page - 1) * limit

	searchKey := c.DefaultQuery("searchKey", "")
	result, err := models.FindAllTags(searchKey, limit, offset)
	pageInfo := services.PageInfo{
		Page:      page,
		Limit:     limit,
		TotalPage: int(math.Ceil(float64(result.Count) / float64(limit))),
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
		Message:  "List All Tagss",
		PageInfo: pageInfo,
		Results:  result.Data,
	})
}

func DetailTags(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tags, err := models.FindOneTags(id)
	if err != nil {
		// log.Println(err)
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "Tags not found",
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
		Message: "Detail Tags",
		Results: tags,
	})
}

func CreateTags(c *gin.Context) {
	data := models.Tags{}

	//upload
	c.ShouldBind(&data)

	// data.Image = lib.Upload(c, "image", "product")
	// //upload
	nameInput := c.PostForm("name")

	if nameInput == "" {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Name Must Not Be Empty",
		})
		return
	}

	checkTags, _ := models.FindOneTagsByName(nameInput)
	checkTagsName := checkTags.Name

	if *checkTagsName == nameInput {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Tags Name Already Exist",
		})
		return
	}

	// c.Bind(&data)
	Tags, err := models.CreateTags(data)
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
		Message: "Tags Created successfully",
		Results: Tags,
	})
}

func UpdateTags(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	//check Tags
	nameInput := c.PostForm("name")
	checkTags, err := models.FindOneTags(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "Tags Variant not found",
			})
			return
		}
		fmt.Println(checkTags)
	}
	if nameInput != "" {
		checkTags, _ := models.FindOneTagsByName(nameInput)
		if nameInput == *checkTags.Name {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "Name already used",
			})
			return
		}
	}

	//

	data := models.Tags{}

	c.ShouldBind(&data)
	// // c.Bind(&data)
	// //upload

	// data.Image = lib.Upload(c, "image", "product")
	// //upload
	data.Id = id

	tags, err := models.UpdateTags(data)
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
		Message: "Tags update successfully",
		Results: tags,
	})
}

func DeleteTags(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tags, err := models.DeleteTags(id)
	if err != nil {
		// log.Fatalln(err)
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

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Delete Tags",
		Results: tags,
	})
}

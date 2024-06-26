package controllers

import (
	"fmt"
	"math"
	"strings"

	"github.com/rayfarandi/fwg17-go-backend/src/models"
	"github.com/rayfarandi/fwg17-go-backend/src/service"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListAllCategories(c *gin.Context) {
	searchKey := c.DefaultQuery("searchKey", "")
	sortBy := c.DefaultQuery("sortBy", "id")
	order := c.DefaultQuery("order", "ASC")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit

	result, err := models.FindAllCategories(searchKey, sortBy, order, limit, offset)
	if len(result.Data) == 0 {
		c.JSON(http.StatusNotFound, &service.ResponseOnly{
			Success: false,
			Message: "data not found",
		})
		return
	}

	totalPage := int(math.Ceil(float64(result.Count) / float64(limit)))
	nextPage := page + 1
	if nextPage > totalPage {
		nextPage = 0
	}
	prevPage := page - 1
	if prevPage < 1 {
		prevPage = 0
	}

	PageInfo := service.PageInfo{
		CurrentPage: page,
		NextPage:    nextPage,
		PrevPage:    prevPage,
		Limit:       limit,
		TotalPage:   totalPage,
		TotalData:   result.Count,
	}

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &service.ResponseList{
		Success:  true,
		Message:  "List all categories",
		PageInfo: PageInfo,
		Results:  result.Data,
	})
}

func DetailCategories(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := models.FindOneCategories(id)
	if err != nil {
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "sql: no rows in result set") {
			c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
				Success: false,
				Message: "Category not found",
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
		Message: "Detail categories",
		Results: category,
	})
}

func CreateCategories(c *gin.Context) {
	data := models.Categories{}
	c.ShouldBind(&data)

	category, err := models.CreateCategories(data)
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
		Message: "Categories created successfully",
		Results: category,
	})
}

func UpdateCategories(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := models.Categories{}

	c.ShouldBind(&data)
	data.Id = id

	isExist, err := models.FindOneCategories(id)
	if err != nil {
		fmt.Println(isExist, err)
		c.JSON(http.StatusNotFound, &service.ResponseOnly{
			Success: false,
			Message: "Category not found",
		})
		return
	}

	category, err := models.UpdateCategories(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &service.Response{
		Success: true,
		Message: "Categories updated successfully",
		Results: category,
	})
}

func DeleteCategories(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	isExist, err := models.FindOneCategories(id)
	if err != nil {
		fmt.Println(isExist, err)
		c.JSON(http.StatusNotFound, &service.ResponseOnly{
			Success: false,
			Message: "Category not found",
		})
		return
	}

	category, err := models.DeleteCategories(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &service.Response{
		Success: true,
		Message: "Delete category successfully",
		Results: category,
	})
}

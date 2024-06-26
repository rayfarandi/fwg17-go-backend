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

func ListAllVariants(c *gin.Context) {
	searchKey := c.DefaultQuery("searchKey", "")
	sortBy := c.DefaultQuery("sortBy", "id")
	order := c.DefaultQuery("order", "ASC")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit

	result, err := models.FindAllVariants(searchKey, sortBy, order, limit, offset)
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
		Message:  "List all variants",
		PageInfo: PageInfo,
		Results:  result.Data,
	})
}

func DetailVariants(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	variants, err := models.FindOneVariants(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
				Success: false,
				Message: "Variants not found",
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
		Message: "Detail variants",
		Results: variants,
	})
}

func CreateVariants(c *gin.Context) {
	data := models.Variants{}
	c.ShouldBind(&data)

	variants, err := models.CreateVariants(data)
	if err != nil {
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "pq: duplicate key") {
			c.JSON(http.StatusBadRequest, &service.ResponseOnly{
				Success: false,
				Message: "duplicate variant name",
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
		Message: "Variants created successfully",
		Results: variants,
	})
}

func UpdateVariants(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := models.Variants{}

	c.ShouldBind(&data)
	data.Id = id

	isExist, err := models.FindOneVariants(id)
	if err != nil {
		fmt.Println(isExist, err)
		c.JSON(http.StatusNotFound, &service.ResponseOnly{
			Success: false,
			Message: "Variants not found",
		})
		return
	}

	variants, err := models.UpdateVariants(data)
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
		Message: "Variants updated successfully",
		Results: variants,
	})
}

func DeleteVariants(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	isExist, err := models.FindOneVariants(id)
	if err != nil {
		fmt.Println(isExist, err)
		c.JSON(http.StatusNotFound, &service.ResponseOnly{
			Success: false,
			Message: "Variants not found",
		})
		return
	}

	variants, err := models.DeleteVariants(id)
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
		Message: "Delete variants successfully",
		Results: variants,
	})
}

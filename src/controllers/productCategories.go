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

func ListAllProductCategories(c *gin.Context) {
	sortBy := c.DefaultQuery("sortBy", "id")
	order := c.DefaultQuery("order", "ASC")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit

	result, err := models.FindAllProductCategories(sortBy, order, limit, offset)
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
		Message:  "List all product categories",
		PageInfo: PageInfo,
		Results:  result.Data,
	})
}

func DetailProductCategories(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pc, err := models.FindOneProductCategories(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
				Success: false,
				Message: "Product categories not found",
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
		Message: "Detail product categories",
		Results: pc,
	})
}

func CreateProductCategories(c *gin.Context) {
	data := models.ProductCategories{}
	c.ShouldBind(&data)

	_, err := models.FindOneProducts(data.ProductId)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, &service.ResponseOnly{
			Success: false,
			Message: "product id not found",
		})
		return
	}

	_, err = models.FindOneCategories(data.CategoryId)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, &service.ResponseOnly{
			Success: false,
			Message: "category id not found",
		})
		return
	}

	pc, err := models.CreateProductCategories(data)
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
		Message: "Product categories created successfully",
		Results: pc,
	})
}

func UpdateProductCategories(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := models.ProductCategories{}

	c.ShouldBind(&data)

	_, err := models.FindOneProducts(data.ProductId)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, &service.ResponseOnly{
			Success: false,
			Message: "product id not found",
		})
		return
	}

	_, err = models.FindOneCategories(data.CategoryId)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, &service.ResponseOnly{
			Success: false,
			Message: "category id not found",
		})
		return
	}

	data.Id = id

	isExist, err := models.FindOneProductCategories(id)
	if err != nil {
		fmt.Println(isExist, err)
		c.JSON(http.StatusNotFound, &service.ResponseOnly{
			Success: false,
			Message: "Product categories not found",
		})
		return
	}

	pc, err := models.UpdateProductCategories(data)
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
		Message: "Product categories updated successfully",
		Results: pc,
	})
}

func DeleteProductCategories(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	isExist, err := models.FindOneProductCategories(id)
	if err != nil {
		fmt.Println(isExist, err)
		c.JSON(http.StatusNotFound, &service.ResponseOnly{
			Success: false,
			Message: "Product categories not found",
		})
		return
	}

	pc, err := models.DeleteProductCategories(id)
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
		Message: "Delete product categories successfully",
		Results: pc,
	})
}

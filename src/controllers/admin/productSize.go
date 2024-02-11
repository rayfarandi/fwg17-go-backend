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

func ListAllProductSize(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	offset := (page - 1) * limit

	result, err := models.FindAllProductSize(limit, offset)
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
		Message:  "List All Products size",
		PageInfo: pageInfo,
		Results:  result.Data,
	})
}

func DetailProductSize(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	productSize, err := models.FindOneProductSize(id)
	if err != nil {
		log.Println(err)
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "Product size not found",
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
		Message: "Detail Product size",
		Results: productSize,
	})
}

func CreateProductSize(c *gin.Context) {
	data := models.ProductSize{}

	c.Bind(&data)
	productSize, err := models.CreateProductSize(data)
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
		Message: "product Created successfully",
		Results: productSize,
	})
}

func UpdateProductSize(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data := models.ProductSize{}

	c.ShouldBind(&data)
	// c.Bind(&data)
	// upload

	// data.Image = lib.Upload(c, "image", "product")
	// upload
	data.Id = id
	isExist, err := models.FindOneProductSize(id)
	if err != nil {
		fmt.Println(isExist, err)
		c.JSON(http.StatusNotFound, &services.ResponseOnly{
			Success: false,
			Message: "Size not found",
		})
		return
	}

	productSize, err := models.UpdateProductSize(data)
	if err != nil {
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
				Success: false,
				Message: "Sizes not found",
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
		Message: "Sizes updated successfully",
		Results: productSize,
	})

	// productSize, err := models.UpdateProductSize(data)
	// log.Println(data)
	// if err != nil {
	// 	log.Println(err)
	// 	c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
	// 		Success: false,
	// 		Message: "Internal server error",
	// 	})
	// 	return
	// }

	// c.JSON(http.StatusOK, &services.Response{
	// 	Success: true,
	// 	Message: "Product update successfully",
	// 	Results: productSize,
	// })
}

func DeleteProductSize(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := models.DeleteProductSize(id)
	if err != nil {
		log.Fatalln(err)
		if strings.HasPrefix(err.Error(), "sql:no rows") {
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
		Message: "Delete product",
		Results: product,
	})
}

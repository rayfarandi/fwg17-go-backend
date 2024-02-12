package controllers_admin

import (
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rayfarandi/fwg17-go-backend/src/models"
	"github.com/rayfarandi/fwg17-go-backend/src/services"
)

func ListAllProductCategories(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	offset := (page - 1) * limit

	result, err := models.FindAllProductCategories(limit, offset)
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
		Message:  "List All Categoriess",
		PageInfo: pageInfo,
		Results:  result.Data,
	})
}

func DetailProductCategories(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	ProductCategories, err := models.FindOneProductCategories(id)
	if err != nil {
		log.Println(err)
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "Product Categories not found",
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
		Message: "Detail Product Categories",
		Results: ProductCategories,
	})
}

func CreateProductCategories(c *gin.Context) {
	data := models.ProductCategories{}

	//upload
	c.ShouldBind(&data)

	// data.Image = lib.Upload(c, "image", "product")
	// //upload

	// c.Bind(&data)
	productCategories, err := models.CreateProductCategories(data)
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
		Message: "Product Categories Created successfully",
		Results: productCategories,
	})
}

func UpdateProductCategories(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data := models.ProductCategories{}

	c.ShouldBind(&data)

	data.Id = id

	productCategories, err := models.UpdateProductCategories(data)
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
		Message: "Product Categories update successfully",
		Results: productCategories,
	})
}

func DeleteProductCategories(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	productCategories, err := models.DeleteProductCategories(id)
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
		Message: "Delete Product Categories",
		Results: productCategories,
	})
}

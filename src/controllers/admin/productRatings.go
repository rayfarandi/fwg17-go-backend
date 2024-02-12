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

func ListAllProductRatings(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	offset := (page - 1) * limit

	searchKey := c.DefaultQuery("searchKey", "")
	result, err := models.FindAllProductRatings(searchKey, limit, offset)
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
		Message:  "List All Products Ratings",
		PageInfo: pageInfo,
		Results:  result.Data,
	})
}

func DetailProductRatings(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	productRatings, err := models.FindOneProductRatings(id)
	if err != nil {
		// log.Println(err)
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "Product Ratings not found",
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
		Message: "Detail Product Ratings",
		Results: productRatings,
	})
}

func CreateProductRatings(c *gin.Context) {
	data := models.ProductRatings{}

	//upload
	c.ShouldBind(&data)

	// data.Image = lib.Upload(c, "image", "product")
	// //upload

	// c.Bind(&data)
	productRatings, err := models.CreateProductRatings(data)
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
		Message: "product Ratings Created successfully",
		Results: productRatings,
	})
}

func UpdateProductRatings(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	//check product

	//

	data := models.ProductRatings{}

	c.ShouldBind(&data)
	// // c.Bind(&data)
	// //upload

	// data.Image = lib.Upload(c, "image", "product")
	// //upload
	data.Id = id

	productRatings, err := models.UpdateProductRatings(data)
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
		Message: "Product Ratings update successfully",
		Results: productRatings,
	})
}

func DeleteProductRatings(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	productRatings, err := models.DeleteProductRatings(id)
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
		Message: "Delete product ratings",
		Results: productRatings,
	})
}

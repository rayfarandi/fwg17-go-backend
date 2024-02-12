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

func ListAllProductTags(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	offset := (page - 1) * limit

	result, err := models.FindAllProductTags(limit, offset)
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

func DetailProductTags(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	Producttags, err := models.FindOneProductTags(id)
	if err != nil {
		log.Println(err)
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "Product Tags not found",
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
		Message: "Detail Product Tags",
		Results: Producttags,
	})
}

func CreateProductTags(c *gin.Context) {
	data := models.ProductTags{}

	//upload
	c.ShouldBind(&data)

	// data.Image = lib.Upload(c, "image", "product")
	// //upload

	// c.Bind(&data)
	productTags, err := models.CreateProductTags(data)
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
		Message: "Product Tags Created successfully",
		Results: productTags,
	})
}

func UpdateProductTags(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data := models.ProductTags{}

	c.ShouldBind(&data)

	data.Id = id

	productTags, err := models.UpdateProductTags(data)
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
		Message: "Product Tags update successfully",
		Results: productTags,
	})
}

func DeleteProductTags(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	productTags, err := models.DeleteProductTags(id)
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
		Message: "Delete Product Tags",
		Results: productTags,
	})
}

package controllers_admin

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rayfarandi/fwg17-go-backend/src/lib"
	"github.com/rayfarandi/fwg17-go-backend/src/models"
	"github.com/rayfarandi/fwg17-go-backend/src/services"
)

func ListAllProduct(c *gin.Context) {
	searchKey := c.DefaultQuery("searchKey", "")
	// category := c.DefaultQuery("category", "")
	sortBy := c.DefaultQuery("sortBy", "id")
	order := c.DefaultQuery("order", "ASC")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	offset := (page - 1) * limit

	// result, err := models.FindAllProduct(searchKey, category, sortBy, order, limit, offset)
	result, err := models.FindAllProduct(searchKey, sortBy, order, limit, offset)

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
		Message:  "List All Products",
		PageInfo: pageInfo,
		Results:  result.Data,
	})
}

func DetailProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := models.FindOneProduct(id)
	if err != nil {
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "Product not found",
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
		Message: "Detail Product",
		Results: product,
	})
}

func CreateProduct(c *gin.Context) {
	data := services.ProductForm{}

	errBind := c.ShouldBind(&data)
	//upload

	_, err := c.FormFile("image")
	if err == nil {
		file, err := lib.Upload(c, "image", "product")
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		data.Image = file
	} else {
		data.Image = ""
	}
	//upload
	if errBind != nil {
		fmt.Println(errBind)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	product, err := models.CreateProduct(data)
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
		Results: product,
	})
}

func UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := services.ProductForm{}

	data.Id = id

	checkProduct, err := models.FindOneProduct(id)
	if err != nil {
		fmt.Println(checkProduct, err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Product not found",
		})
		return
	}

	c.ShouldBind(&data)

	//

	// c.Bind(&data)
	//upload

	_, err = c.FormFile("image")
	if err == nil {
		err := os.Remove("./" + checkProduct.Image)
		if err != nil {
		}

		file, err := lib.Upload(c, "image", "product")
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		data.Image = file
	}
	//upload
	data.Id = id

	product, err := models.UpdateProduct(data)
	if err != nil {
		fmt.Println(err, product)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Product update successfully",
		Results: product,
	})
}

func DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := models.FindOneProduct(id)
	if err != nil {
		fmt.Println(data, err)
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
	product, err := models.DeleteProduct(id)

	if product.Image != "" {
		err := os.Remove("./" + product.Image)
		if err != nil {
			fmt.Println("Error deleting file:", err)
		}
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Delete product",
		Results: product,
	})
}

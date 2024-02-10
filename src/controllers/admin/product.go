package controllers_admin

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rayfarandi/fwg17-go-backend/src/lib"
	"github.com/rayfarandi/fwg17-go-backend/src/models"
	"github.com/rayfarandi/fwg17-go-backend/src/services"
)

func ListAllProduct(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	offset := (page - 1) * limit

	searchKey := c.DefaultQuery("searchKey", "")
	result, err := models.FindAllProduct(searchKey, limit, offset)
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
		Message:  "List All Products",
		PageInfo: pageInfo,
		Results:  result.Data,
	})
}

func DetailProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := models.FindOneProduct(id)
	if err != nil {
		// log.Println(err)
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
	data := models.Product{}

	//upload
	c.ShouldBind(&data)

	data.Image = lib.Upload(c, "image", "product")
	//upload
	nameInput := c.PostForm("name")

	if nameInput == "" {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Name Must Not Be Empty",
		})
		return
	}

	checkProduct, _ := models.FindOneProductByName(nameInput)
	checkProductName := checkProduct.Name

	if *checkProductName == nameInput {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Product Name Already Exist",
		})
		return
	}

	// c.Bind(&data)
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

	//check product
	nameInput := c.PostForm("name")
	checkProduct, err := models.FindOneProduct(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "Product not found",
			})
			return
		}
		fmt.Println(checkProduct)
	}
	if nameInput != "" {
		checkProduct, _ := models.FindOneProductByName(nameInput)
		if nameInput == *checkProduct.Name {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "Name already used",
			})
			return
		}
	}

	//

	data := models.Product{}

	c.ShouldBind(&data)
	// c.Bind(&data)
	//upload

	data.Image = lib.Upload(c, "image", "product")
	//upload
	data.Id = id

	product, err := models.UpdateProduct(data)
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
		Message: "Product update successfully",
		Results: product,
	})
}

func DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := models.DeleteProduct(id)
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

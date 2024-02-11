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

func ListAllProductVariant(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	offset := (page - 1) * limit

	searchKey := c.DefaultQuery("searchKey", "")
	result, err := models.FindAllProductVariant(searchKey, limit, offset)
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

func DetailProductVariant(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	productVariant, err := models.FindOneProductVariant(id)
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
		Results: productVariant,
	})
}

func CreateProductVariant(c *gin.Context) {
	data := models.ProductVariant{}

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

	checkProductVariant, _ := models.FindOneProductVariantByName(nameInput)
	checkProductVariantName := checkProductVariant.Name

	if *checkProductVariantName == nameInput {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Product Name Already Exist",
		})
		return
	}

	// c.Bind(&data)
	productVariant, err := models.CreateProductVariant(data)
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
		Results: productVariant,
	})
}

func UpdateProductVariant(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	//check product
	nameInput := c.PostForm("name")
	checkProduct, err := models.FindOneProductVariant(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "Product Variant not found",
			})
			return
		}
		fmt.Println(checkProduct)
	}
	if nameInput != "" {
		checkProduct, _ := models.FindOneProductVariantByName(nameInput)
		if nameInput == *checkProduct.Name {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "Name already used",
			})
			return
		}
	}

	//

	data := models.ProductVariant{}

	c.ShouldBind(&data)
	// // c.Bind(&data)
	// //upload

	// data.Image = lib.Upload(c, "image", "product")
	// //upload
	data.Id = id

	productVariant, err := models.UpdateProductVariant(data)
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
		Results: productVariant,
	})
}

func DeleteProductVariant(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	productVariant, err := models.DeleteProductVariant(id)
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
		Results: productVariant,
	})
}

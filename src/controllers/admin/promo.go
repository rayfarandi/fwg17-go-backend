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

func ListAllPromo(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	offset := (page - 1) * limit

	result, err := models.FindAllPromo(limit, offset)
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

func DetailPromo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	Promo, err := models.FindOnePromo(id)
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
		Results: Promo,
	})
}

func CreatePromo(c *gin.Context) {
	data := models.Promo{}

	//upload
	c.ShouldBind(&data)

	// data.Image = lib.Upload(c, "image", "product")
	// //upload

	// c.Bind(&data)
	Promo, err := models.CreatePromo(data)
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
		Results: Promo,
	})
}

func UpdatePromo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data := models.Promo{}

	c.ShouldBind(&data)

	data.Id = id

	Promo, err := models.UpdatePromo(data)
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
		Results: Promo,
	})
}

func DeletePromo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	Promo, err := models.DeletePromo(id)
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
		Message: "Delete Promo",
		Results: Promo,
	})
}

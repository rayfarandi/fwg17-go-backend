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

func ListAllOrders(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	offset := (page - 1) * limit

	searchKey := c.DefaultQuery("searchKey", "")
	result, err := models.FindAllOrders(searchKey, limit, offset)
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
		Message:  "List All Products Categories",
		PageInfo: pageInfo,
		Results:  result.Data,
	})
}

func DetailOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	order, err := models.FindOneOrder(id)
	if err != nil {
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "Order Not Found",
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
		Message: "Detail Order",
		Results: order,
	})
}

func CreateOrder(c *gin.Context) {
	data := models.Order{}
	fullNameInput := c.PostForm("fullName")

	if fullNameInput == "" {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Name Must Not Be Empty",
		})
		return
	}
	existingOrder, err := models.FindOneOrderByOrderNumber(*data.OrderNumber)

	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Order Number is Already Exist",
		})
		return
	} else {
		fmt.Println(existingOrder)
	}

	c.ShouldBind(&data)

	order, err := models.CreateOrder(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Order Created Successfully",
		Results: order,
	})
}

func UpdateOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	existingOrder, err := models.FindOneOrder(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "Order Not Found",
			})
			return
		}
		fmt.Println(existingOrder)
	}

	data := models.Order{}

	c.ShouldBind(&data)

	data.Id = id

	order, err := models.UpdateOrder(data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Update Order Succesfully",
		Results: order,
	})
}

func DeleteOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	order, err := models.DeleteOrder(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "No Data",
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
		Message: "Delete Order Succesfully",
		Results: order,
	})
}

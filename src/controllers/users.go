package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rayfarandi/fwg17-go-backend/src/models"
)

type pageInfo struct {
	Page int `json:"page"`
}

type responseList struct {
	Success  bool        `json:"success"`
	Message  string      `json:"message"`
	PageInfo pageInfo    `json:"pageInfo"`
	Results  interface{} `json:"results"`
}
type response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Results interface{} `json:"results"`
}

type User struct {
	Id       int    `json:"id" form:"id"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
}
type responseOnly struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func ListAllUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	users, err := models.FindAllUsers()
	if err != nil {
		// log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, &responseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, &responseList{
		Success: true,
		Message: "List All Users",
		PageInfo: pageInfo{
			Page: page,
		},
		Results: users,
	})
}

func DetailUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, &response{
		Success: true,
		Message: "List All Users",
		Results: User{
			Id:       id,
			Email:    "admin@mail.com",
			Password: "1234",
		},
	})
}

func CreateUser(c *gin.Context) {
	user := User{}

	c.ShouldBind(&user)

	c.JSON(http.StatusOK, &response{
		Success: true,
		Message: "User Created successfully",
		Results: user,
	})
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := User{}

	c.Bind(&user)
	user.Id = id
	c.JSON(http.StatusOK, &response{
		Success: true,
		Message: "User update successfully",
		Results: user,
	})
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, &response{
		Success: true,
		Message: "User Successfully deleted",
		Results: User{
			Id:       id,
			Email:    "fazztrack@mail.com",
			Password: "1234",
		},
	})
}

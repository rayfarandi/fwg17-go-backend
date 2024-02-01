package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
}

func ListAllUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	c.JSON(http.StatusOK, &responseList{
		Success: true,
		Message: "List All Users",
		PageInfo: pageInfo{
			Page: page,
		},
		Results: []User{
			{
				Id:       1,
				Email:    "admin@mail.com",
				Password: "1234",
			},
			{
				Id:       2,
				Email:    "fazztrack@mail.com",
				Password: "1234",
			},
		},
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

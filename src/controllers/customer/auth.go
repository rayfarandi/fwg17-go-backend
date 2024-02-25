package controllers_customer

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/rayfarandi/fwg17-go-backend/src/lib"
	"github.com/rayfarandi/fwg17-go-backend/src/models"
	"github.com/rayfarandi/fwg17-go-backend/src/service"

	"github.com/KEINOS/go-argonize"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	form := service.UserForm{}
	err := c.ShouldBind(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest, &service.ResponseOnly{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	defaultRole := "customer"
	form.Role = &defaultRole

	plain := []byte(form.Password)
	hash, _ := argonize.Hash(plain)
	form.Password = hash.String()

	result, err := models.CreateUser(form)

	if err != nil {
		fmt.Println(err)
		if strings.HasSuffix(err.Error(), `unique constraint "users_email_key"`) {
			c.JSON(http.StatusBadRequest, &service.ResponseOnly{
				Success: false,
				Message: "email already registered. . . please login",
			})
			return
		}
		c.JSON(http.StatusBadRequest, &service.ResponseOnly{
			Success: false,
			Message: "Register failed",
		})
		return
	}

	c.JSON(http.StatusOK, &service.Response{
		Success: true,
		Message: "register success. welcome Home",
		Results: result,
	})
}

func ForgotPassword(c *gin.Context) {
	form := service.FormReset{}
	c.ShouldBind(&form)

	if form.Email != "" {
		found, _ := models.FindOneUsersByEmail(form.Email)

		if found.Id == 0 {
			c.JSON(http.StatusBadRequest, &service.ResponseOnly{
				Success: false,
				Message: "email not registered. use other email",
			})
			return
		}

		FormReset := service.ForgotPassword{
			Otp:   lib.RandomNumberStr(6),
			Email: form.Email,
		}
		models.CreateForgotPassword(FormReset)
		// start send email
		fmt.Println(FormReset.Otp)
		// end send email
		c.JSON(http.StatusOK, &service.ResponseOnly{
			Success: true,
			Message: "OTP has sent to your email",
		})
		return
	}

	if form.Otp != "" {
		found, _ := models.FindOneByOtp(form.Otp)
		if found.Id == 0 {
			c.JSON(http.StatusBadRequest, &service.ResponseOnly{
				Success: false,
				Message: "invalid OTP code. please enter correct code",
			})
			return
		}

		foundUser, _ := models.FindOneUsersByEmail(found.Email)
		data := service.UserForm{
			Id: foundUser.Id,
		}

		hash, _ := argonize.Hash([]byte(form.Password))
		data.Password = hash.String()

		updated, err := models.UpdateUser(data)
		if err != nil {
			fmt.Println(updated, err)
			c.JSON(http.StatusBadRequest, &service.ResponseOnly{
				Success: false,
				Message: err.Error(),
			})
		}

		models.DeleteForgotPassword(found.Id)
		message := fmt.Sprintf("Reset password for %v success", *updated.Email)
		c.JSON(http.StatusOK, &service.ResponseOnly{
			Success: true,
			Message: message,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
		Success: false,
		Message: "Internal server error",
	})
}

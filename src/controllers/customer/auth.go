package controllers_customer

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/KEINOS/go-argonize"
	"github.com/gin-gonic/gin"

	"github.com/rayfarandi/fwg17-go-backend/src/lib"
	"github.com/rayfarandi/fwg17-go-backend/src/models"
	"github.com/rayfarandi/fwg17-go-backend/src/service"
)

type FormReset struct {
	Email           string `form:"email"`
	Otp             string `form:"otp"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirmPassword" binding:"eqfield=Password"`
}

func Register(c *gin.Context) {
	form := models.UserForm{}
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
		Message: "register success. . . welcome aboard!",
		Results: result,
	})
}

func ForgotPassword(c *gin.Context) {
	form := FormReset{}
	c.ShouldBind(&form)

	if form.Email != "" {
		found, _ := models.FindOneUsersByEmail(form.Email)

		if found.Id == 0 {
			c.JSON(http.StatusBadRequest, &service.ResponseOnly{
				Success: false,
				Message: "email not registered. . . . please use another email",
			})
			return
		}

		FormReset := models.ForgotPassword{
			Otp:   lib.RandomNumberStr(6),
			Email: form.Email,
		}
		models.CreateForgotPassword(FormReset)
		// START SEND EMAIL
		fmt.Println(FormReset.Otp)
		// END SEND EMAIL
		c.JSON(http.StatusOK, &service.ResponseOnly{
			Success: true,
			Message: "OTP has been sent to your email",
		})
		return
	}

	if form.Otp != "" {
		found, _ := models.FindOneByOtp(form.Otp)
		if found.Id == 0 {
			c.JSON(http.StatusBadRequest, &service.ResponseOnly{
				Success: false,
				Message: "invalid OTP code. . . please enter the correct code",
			})
			return
		}

		// if form.Password != form.ConfirmPassword{
		// 	c.JSON(http.StatusBadRequest, &service.ResponseOnly{
		// 		Success: false,
		// 		Message: "Confirm password does not match",
		// 	})
		// 	return
		// }

		foundUser, _ := models.FindOneUsersByEmail(found.Email)
		data := models.UserForm{
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

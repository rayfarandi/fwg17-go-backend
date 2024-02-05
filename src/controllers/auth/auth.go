package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/KEINOS/go-argonize"
	"github.com/gin-gonic/gin"
	"github.com/rayfarandi/fwg17-go-backend/src/lib"
	"github.com/rayfarandi/fwg17-go-backend/src/models"
	"github.com/rayfarandi/fwg17-go-backend/src/services"
)

type FormReset struct {
	// Email           string `json:"email" form:"email" binding:"email"`
	Id              int    `db:"id"`
	Email           string `form:"email" `
	Otp             string `form:"otp"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirmPassword"`
}

func Login(c *gin.Context) {
	form := models.User{}
	err := c.ShouldBind(&form)

	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "invalid",
		})
		return
	}
	found, err := models.FindOneUserEmail(form.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, services.ResponseOnly{
			Success: false,
			Message: "wrong email or password",
		})
		return
	}
	decoded, err := argonize.DecodeHashStr(found.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, services.ResponseOnly{
			Success: false,
			Message: "wrong email or password",
		})
		return
	}
	plain := []byte(form.Password)
	if decoded.IsValidPassword(plain) {
		c.JSON(http.StatusOK, &services.ResponseOnly{
			Success: true,
			Message: "Login success",
		})
		return
	} else {
		c.JSON(http.StatusUnauthorized, services.ResponseOnly{
			Success: false,
			Message: "wrong email or password",
		})
		return
	}

}

func Register(c *gin.Context) {
	form := models.User{}

	err := c.ShouldBind(&form)

	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "invalid",
		})
		return
	}

	plain := []byte(form.Password)
	hash, err := argonize.Hash(plain)
	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Failed hash password",
		})
		return
	}
	form.Password = hash.String()

	_, err = models.CreateUser(form)
	if err != nil {
		if strings.HasSuffix(err.Error(), `unique constraint "users_email_key"`) {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "Email already register",
			})
			return
		}
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Failed register",
		})
		return
	}
	c.JSON(http.StatusOK, &services.ResponseOnly{
		Success: true,
		Message: "Register success",
	})
}

func ForgotPassword(c *gin.Context) {
	form := FormReset{}
	c.ShouldBind(&form)
	if form.Email != "" {
		found, _ := models.FindOneUserEmail(form.Email)
		if found.Id != 0 {
			formReset := models.FormReset{
				Otp:   lib.RandomNumberStr(6),
				Email: found.Email,
			}
			models.CreateResetPassword(formReset)
			//start send email
			fmt.Println(formReset.Otp)
			//end send email
			c.JSON(http.StatusOK, &services.ResponseOnly{
				Success: true,
				Message: "OTP has been sent to your email",
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "Failid to reset",
			})
		}
	}

	if form.Otp != "" {
		// log.Fatalln(form.Otp)
		found, _ := models.FindOneRPByOtp(form.Otp)
		log.Println(found.Id)

		if found.Id == 0 {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "Failed to reset password",
			})
			return
		}
		if form.Password != form.ConfirmPassword {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "confirm password does not match",
			})
			return
		}

		//mainflow
		foundUser, _ := models.FindOneUserByEmail((found.Email))
		data := models.User{
			Id: foundUser.Id,
		}
		hash, _ := argonize.Hash([]byte(form.Password))
		data.Password = hash.String()

		updated, _ := models.UpdateUser(data)
		message := fmt.Sprintf("Reset password for %v succes", updated.Email)
		c.JSON(http.StatusOK, &services.ResponseOnly{
			Success: true,
			Message: message,
		})
		models.DeleteResetPassword(found.Id)
		return
	}
	c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
		Success: false,
		Message: "Internal server error",
	})
}

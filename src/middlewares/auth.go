package middlewares

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/KEINOS/go-argonize"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rayfarandi/fwg17-go-backend/src/models"
	"github.com/rayfarandi/fwg17-go-backend/src/services"
)

func Auth() (*jwt.GinJWTMiddleware, error) {
	godotenv.Load()
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "go-backend",
		Key:         []byte(os.Getenv("APP_SECRET")),
		IdentityKey: "id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			user := data.(*services.User)
			return jwt.MapClaims{
				"id":   user.Id,
				"role": user.Role,
			}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &services.User{
				Id:   int(claims["id"].(float64)),
				Role: claims["role"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			form := services.User{}
			err := c.ShouldBind(&form)

			if err != nil {

				return nil, err
			}
			found, err := models.FindOneUserEmail(form.Email)
			if err != nil {

				return nil, err
			}
			decoded, err := argonize.DecodeHashStr(found.Password)
			if err != nil {

				return nil, err
			}
			plain := []byte(form.Password)
			if decoded.IsValidPassword(plain) {

				return &services.User{
					Id:   found.Id,
					Role: found.Role,
				}, nil
			} else {
				return nil, errors.New("invalid_password")
			}
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			user := data.(*services.User)
			if strings.HasPrefix(c.Request.URL.Path, "/admin") {
				if user.Role != "admin" {
					return false
				}
			}
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {

			//menampilkan pesan error jika salah memasukan email or pass
			// if strings.HasPrefix(c.Request.URL.Path, "/login") {
			if strings.HasPrefix(c.Request.URL.Path, "/auth/login") {
				c.JSON(http.StatusUnauthorized, &services.ResponseOnly{
					Success: false,
					Message: "wrong Email or password",
				})
				return
			}
			// if strings.HasPrefix(c.Request.URL.Path, "/forgot-password") {
			if strings.HasPrefix(c.Request.URL.Path, "/auth/forgot-password") {
				c.JSON(http.StatusUnauthorized, &services.ResponseOnly{
					Success: false,
					Message: "email not registered",
				})
				return
			}

			c.JSON(http.StatusUnauthorized, &services.ResponseOnly{
				Success: false,
				Message: "Unauthorized",
			})
		}, LoginResponse: func(c *gin.Context, code int, token string, time time.Time) {
			c.JSON(http.StatusOK, &services.Response{
				Success: true,
				Message: "Login success",
				Results: struct {
					Token string `json:"token"`
				}{
					Token: token,
				},
			})
		},
	})

	if err != nil {
		return nil, err
	}
	return authMiddleware, nil
}

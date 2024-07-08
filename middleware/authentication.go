package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/AswinJoseOpen/Login-Auth/config"
	"github.com/AswinJoseOpen/Login-Auth/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func AuthMiddleWare(db *gorm.DB, config config.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("testing middleware")

		//Get Header
		tokenString := c.GetHeader("Authorization")[7:]
		fmt.Println("token:", tokenString)
		//Validate Token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.Secret), nil
		})
		if err != nil {
			fmt.Println("1;;;", err.Error())
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			fmt.Println(claims["sub"], claims["exp"])
			//check expiry
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				fmt.Println("2")
				c.AbortWithStatus(http.StatusUnauthorized)
			}
			//find user from token
			var user *model.Users

			db.First(&user, claims["sub"])
			if user.ID == 0 {
				fmt.Println("3")
				c.AbortWithStatus(http.StatusUnauthorized)
			}
			//continue
			c.Set("user", user)
			c.Next()

		} else {
			fmt.Println("4")
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}

}

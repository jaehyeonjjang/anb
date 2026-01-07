package router

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"repair/models"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dgrijalva/jwt-go/v4"
)

type AuthTokenClaims struct {
	User               models.User `json:"user"`
	jwt.StandardClaims             // 표준 토큰 Claims
}

var _secretCode string = "SecretCode"

func JwtAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string

		/*
			if (c.Request.Method == "POST" || c.Request.Method == "PUT") && c.Request.URL.String() == "/api/user" {
				c.Next()
				return
			}
		*/

		u, _ := url.Parse(c.Request.URL.String())
		path := u.Path

		if path == "/api/jwt" {
			c.Next()
			return
		}

		if c.Request.Method == "GET" && path == "/api/program" {
			c.Next()
			return
		}

		if values, _ := c.Request.Header["Authorization"]; len(values) > 0 {
			str := values[0]

			if len(str) > 7 && str[:7] == "Bearer " {
				token = str[7:]

				claims := AuthTokenClaims{}
				key := func(token *jwt.Token) (any, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, errors.New("Unexpected Signing Method")
					}
					return []byte(_secretCode), nil
				}

				_, err := jwt.ParseWithClaims(token, &claims, key)
				if err == nil {
					c.Set("user", &(claims.User))
					c.Next()
					return
				}
			} else {
				log.Println("Jwt header is broken")
			}
		} else {
			log.Println("Jwt header not found")
		}

		c.Writer.WriteHeader(http.StatusUnauthorized)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Write([]byte(`{"code":"error","message":"not auth"}`))
		c.Abort()

	}
}

func JwtAuth(c *gin.Context, loginid string, passwd string) gin.H {
	log.Println(loginid, passwd)
	conn := models.NewConnection()

	manager := models.NewUserManager(conn)
	user := manager.GetByLoginid(loginid)

	if user == nil {
		return gin.H{
			"code":    "error",
			"message": "user not found",
		}
	}

	if user.Passwd != passwd {
		return gin.H{
			"code":    "error",
			"message": "wrong password",
		}
	}

	at := AuthTokenClaims{
		User: *user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * 24 * 365 * 10)),
		},
	}

	atoken := jwt.NewWithClaims(jwt.SigningMethodHS256, &at)
	signedAuthToken, _ := atoken.SignedString([]byte(_secretCode))

	user.Passwd = ""
	return gin.H{
		"code":  "ok",
		"token": signedAuthToken,
		"user":  user,
	}
}

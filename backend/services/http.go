package services

import (
	"anb/config"
	"anb/global"
	"anb/router"
	"net/http"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Refresh(c *gin.Context, url string) {
	str := "<script>location.href = '" + url + "';</script>"
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(str))
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Http() {
	r := gin.Default()
	r.Use(CORSMiddleware())

	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	/*
		r.Static("/assets", "./assets")
		r.Static("/webdata", "./webdata")
		r.Static("/_images", "./assets/_images")
		r.Static("/_common", "./assets/_common")
		r.Static("/img", "./assets/img")

		r.Static("/.well-known/pki-validation", "./")

		r.GET("/", func(c *gin.Context) {
			auth, _ := c.Cookie("auth")
			c.SetCookie("auth", "on", 0, "/", "", false, true)

			session := sessions.Default(c)
			user := session.Get("user")

			if auth != "on" {
				if user != nil {
					session.Delete("user")
					session.Save()
					user = nil
				}
			}

			Refresh(c, "/admin/user")
		})
	*/

	r.Static("/assets", "../front/www/assets")
	r.Static("/js", "../front/www/js")
	r.Static("/webdata", "./webdata")
	r.GET("/", func(c *gin.Context) {
		content := global.ReadFile("../front/www/index.html")

		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, content)
	})

	r.GET("/make", func(c *gin.Context) {
		content := global.ReadFile("../front/www/make.html")

		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, content)
	})

	router.SetRouter(r)

	s := &http.Server{
		Addr:           ":" + config.Port,
		Handler:        r,
		ReadTimeout:    10 * time.Minute,
		WriteTimeout:   10 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

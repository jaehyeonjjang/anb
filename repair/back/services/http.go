package services

import (
	"net/http"
	"repair/chat"
	"repair/global/config"
	"repair/router"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Header("Access-Control-Allow-Credentials", "true")
		for _, v := range config.Cors {
			c.Header("Access-Control-Allow-Origin", v)
		}
		c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Http() {
	if config.Mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(CORSMiddleware())

	c := chat.NewChat()

	r.GET("/socket.io/", gin.WrapH(c.Server))
	r.POST("/socket.io/", gin.WrapH(c.Server))

	r.Handle("WS", "/socket.io/", gin.WrapH(c.Server))
	r.Handle("WSS", "/socket.io/", gin.WrapH(c.Server))

	r.Static("/assets", "./dist/assets")
	r.Static("/webdata", "./webdata")
	r.NoRoute(func(c *gin.Context) {
		c.File("./dist/index.html")
	})

	router.SetRouter(r)

	if config.Tls.Use {
		secureFunc := func() gin.HandlerFunc {
			return func(c *gin.Context) {
				secureMiddleware := secure.New(secure.Options{
					SSLRedirect: true,
					SSLHost:     ":" + config.Port,
				})
				err := secureMiddleware.Process(c.Writer, c.Request)

				// If there was an error, do not continue.
				if err != nil {
					return
				}

				c.Next()
			}
		}()

		router := gin.Default()
		router.Use(secureFunc)

		router.GET("/", func(c *gin.Context) {
			c.String(200, "X-Frame-Options header is now `DENY`.")
		})

		go router.Run(":" + config.Port)

		s := &http.Server{
			Addr:           ":" + config.Port,
			Handler:        r,
			ReadTimeout:    10 * time.Minute,
			WriteTimeout:   10 * time.Minute,
			MaxHeaderBytes: 1 << 20,
		}
		s.ListenAndServeTLS(config.Tls.Cert, config.Tls.Key)
	} else {
		s := &http.Server{
			Addr:           ":" + config.Port,
			Handler:        r,
			ReadTimeout:    10 * time.Minute,
			WriteTimeout:   10 * time.Minute,
			MaxHeaderBytes: 1 << 20,
		}
		s.ListenAndServe()
	}
}

package router

import (
	"anb/global"
	"anb/models"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthRequired(isMobile bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		u, _ := url.Parse(c.Request.URL.String())
		path := u.Path

		if strings.Contains(path, "/api/report") {
			c.Next()
			return
		}

		session := sessions.Default(c)
		user := session.Get("user")

		if user != nil {
			c.Next()
			return
		}

		var str string
		str = "<script>location.href = '" + "/login/login" + "';</script>"
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write([]byte(str))
		c.Abort()
	}
}

func AuthAdminRequired(isMobile bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		var str string
		if user != nil {
			item := user.(*models.User)
			if item.Level >= int(global.Manager) {
				c.Next()
				return
			}
		}

		str = "<script>location.href = '" + "/login/login" + "';</script>"
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write([]byte(str))
		c.Abort()
	}
}

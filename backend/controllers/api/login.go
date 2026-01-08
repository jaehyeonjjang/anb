package api

import (
	"anb/config"
	"anb/controllers"
	"anb/global"
	"anb/models"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/contrib/sessions"
)

type LoginController struct {
	controllers.Controller
}

func (c *LoginController) AjaxLogin() {
	loginid := c.Get("loginid")
	passwd := c.Get("passwd")

	var user *models.User

	if config.LocalMode == "true" {
		items := models.GetUsers()

		for _, item := range items {
			log.Println(item)
			log.Println(loginid)
			if item.Loginid == loginid {
				user = &item

				log.Println("find")
				log.Println(user)

				break
			}
		}
	} else {
		conn := c.NewConnection()

		manager := models.NewUserManager(conn)

		user = manager.GetByLoginid(loginid)
	}

	if user == nil {
		c.Set("code", "fail")
		return
	}

	hashedPasswd := global.GetSha256(passwd)
	
	if user.Passwd != hashedPasswd {
		c.Set("code", "fail")
		return
	}

	session := sessions.Default(c.Context)

	if session.Get("user") != nil {
		session.Delete("user")
	}

	user.Passwd = ""

	session.Set("user", user)
	session.Save()

	// Generate a simple token
	token := global.GetSha256(fmt.Sprintf("%d_%s_%d", user.Id, user.Loginid, time.Now().Unix()))

	c.Set("user", user)
	c.Set("level", user.Level)
	c.Set("token", token)
}

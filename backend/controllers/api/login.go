package api

import (
	"anb/config"
	"anb/controllers"
	"anb/global"
	"anb/models"
	"log"

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
		c.Set("code", "user not found")
		log.Println("user not found")
	} else if user.Passwd != passwd && user.Passwd != global.GetSha256(passwd) {
		c.Set("code", "wrong password")
		log.Println("wrong password")
	} else if user.Status != int(global.Default) {
		c.Set("code", "not permit")
		log.Println("not permit")
	} else {
		session := sessions.Default(c.Context)

		if session.Get("user") != nil {
			session.Delete("user")
		}

		user.Passwd = ""

		session.Set("user", user)
		session.Save()

		c.Set("user", user)
		c.Set("level", user.Level)
	}
}

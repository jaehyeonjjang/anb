package api

import (
	"anb/config"
	"anb/controllers"
	"anb/models"
)

type AptController struct {
	controllers.Controller
}

func (c *AptController) AjaxList() {
	if config.LocalMode == "true" {
		items := make([]models.Aptuserlist, 0)

		aptusers := models.GetAptusers()

		for _, item := range aptusers {
			if c.Session.Id == item.User {
				var aptuserlist models.Aptuserlist

				aptuserlist.User = item.User
				aptuserlist.Apt = item.Apt
				aptuserlist.Level = item.Level

				items = append(items, aptuserlist)
			}
		}

		users := models.GetUsers()

		for _, user := range users {
			if user.Id == c.Session.Id {
				for i, _ := range items {
					items[i].Id = user.Id
					items[i].Loginid = user.Loginid
					items[i].Passwd = user.Passwd
					items[i].Name = user.Name
					items[i].Level = user.Level
					items[i].Hp = user.Hp
					items[i].Email = user.Email
					items[i].Date = user.Date
				}
				break
			}
		}

		apts := models.GetApts()

		for i, aptuser := range items {
			for _, apt := range apts {
				if aptuser.Apt == apt.Id {
					items[i].Aptname = apt.Name
					break
				}
			}
		}

		c.Set("items", items)

		return
	}

	conn := c.NewConnection()

	user := c.Geti64("user")

	/*
		userManager := models.NewUserManager(conn)
		userItem := userManager.Get(user)

		if userItem == nil {
			return
		}

		if userItem.Level == global.Admin {
			var items []models.Aptuserlist

			aptManager := models.NewAptManager(conn)
			apts := aptManager.GetListByCompanyStatus(userItem.Company, 1, 0, 0, "")

			c.Set("items", &items)
		} else {
			aptUserManager := models.NewAptuserlistManager(conn)
			items := aptUserManager.GetListByUserAptstatus(user, 1, 0, 0, "")

			c.Set("items", items)
		}
	*/

	aptUserManager := models.NewAptuserlistManager(conn)
	items := aptUserManager.GetListByUserAptstatus(user, 1, 0, 0, "")

	c.Set("items", items)
}

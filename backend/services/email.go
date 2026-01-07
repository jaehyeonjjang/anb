package services

import (
	"anb/global"
	"anb/models"

	log "github.com/sirupsen/logrus"
)

func SendEmail() {
	log.Println("Cron : SendEmail Start")

	conn := models.NewConnection()
	defer conn.Close()

	manager := models.NewSendmailManager(conn)
	items := manager.GetListByStatus(int(global.Use), 0, 0, "")

	if items == nil {
		return
	}

	userManager := models.NewUserManager(conn)

	for _, item := range *items {
		item.Status = int(global.NotUse)
		manager.Update(&item)

		users := userManager.GetListByLevel(item.Level, 0, 0, "")

		if users == nil {
			continue
		}

		for _, user := range *users {
			global.SendEmail(user.Email, item.Title, item.Content)
		}
	}
}

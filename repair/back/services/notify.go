package services

import (
	"repair/controllers/api/document/detail/image"
	"repair/controllers/api/document/periodic"
	"repair/global"
	"repair/models"
	pr "repair/periodic"
	"time"

	log "github.com/sirupsen/logrus"
)

func Notify() {
	log.Println("Notify Service Start")

	ch := global.GetChannel()

	go func() {
		time.Sleep(1 * time.Second)

		for {
			conn := models.NewConnection()
			periodicblueprintzoomManager := models.NewPeriodicblueprintzoomManager(conn)
			periodicblueprintzoomManager.SelectLog = false

			items := periodicblueprintzoomManager.Find([]interface{}{
				models.Where{Column: "status", Value: 1, Compare: "="},
				models.Ordering("pb_id"),
			})

			conn.Close()

			for _, v := range items {
				global.SendNotify(global.Notify{Type: global.NotifyBlueprint, Periodic: v.Periodic, Blueprint: v.Blueprint})
			}

			time.Sleep(10 * time.Second)
		}
	}()

	for {
		msg := <-ch

		switch msg.Type {
		case global.NotifyBlueprint:
			conn := models.NewConnection()
			periodicblueprintzoomManager := models.NewPeriodicblueprintzoomManager(conn)
			blueprintManager := models.NewBlueprintManager(conn)
			periodicdataManager := models.NewPeriodicdataManager(conn)

			periodicblueprintzoom := periodicblueprintzoomManager.GetByPeriodicBlueprint(msg.Periodic, msg.Blueprint)

			blueprint := blueprintManager.Get(msg.Blueprint)
			periodicdatas := periodicdataManager.Find([]interface{}{
				models.Where{Column: "periodic", Value: msg.Periodic, Compare: "="},
				models.Where{Column: "blueprint", Value: msg.Blueprint, Compare: "="},
				models.Ordering("pd_order"),
			})

			if periodicblueprintzoom != nil {
				if blueprint != nil {
					log.Println("Make image", msg.Periodic)
					pr.MakeImage(msg.Periodic, *blueprint, periodicdatas, float64(periodicblueprintzoom.Iconzoom), float64(periodicblueprintzoom.Numberzoom), float64(periodicblueprintzoom.Crackzoom))

					inclination := false
					fiber := false
					meterial := false
					for _, v2 := range periodicdatas {
						if v2.Type >= 200 && v2.Type < 300 {
							inclination = true
						} else if v2.Type >= 300 && v2.Type < 400 {
							fiber = true
						} else if v2.Type >= 400 && v2.Type < 500 {
							meterial = true
						}
					}

					if inclination == true {
						image.MakeInclinationImage(msg.Periodic, *blueprint, periodicdatas, float64(periodicblueprintzoom.Iconzoom))
					}

					if fiber == true {
						image.MakeFiberImage(msg.Periodic, *blueprint, periodicdatas, float64(periodicblueprintzoom.Iconzoom))
					}

					if meterial == true {
						image.MakeMeterialImage(msg.Periodic, *blueprint, periodicdatas, float64(periodicblueprintzoom.Iconzoom))
					}
				}

				periodicblueprintzoomManager.UpdateStatusByPeriodicBlueprint(2, msg.Periodic, msg.Blueprint)
			} else {
				log.Println("periodicblueprintzoom is nil")
			}

			conn.Close()
		case global.NotifyImage:
			log.Println("notify image")
			periodic.GetThumbnail(msg.Filename)
		}
	}

}

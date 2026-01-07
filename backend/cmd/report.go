package main

import (
	"anb/config"
	"anb/models"
	"anb/services"
	"encoding/gob"
	"os"
	"runtime"

	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	gob.Register(&models.User{})

	rand.Seed(time.Now().UnixNano())
}
func main() {
	/*
		width, height := global.GetImageSize("webdata/8943095c39a34dc7b9ef0aa57c2dce29.png")

		fullFilename := fmt.Sprintf("webdata/%v", "86-13007.png")
		targetFilename := fmt.Sprintf("webdata/%v", "86-13007_.png")
		global.MakeThumbnailFault(1600, 1193, width, height, fullFilename, targetFilename)
	*/

	var filename string = "logfile.log"
	// Create the log file if doesn't exist. And append to it if it already exists.
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	if err == nil {
		log.SetOutput(f)
	}

	log.Printf("ANB Version=" + config.Version + " Build=" + config.Build)
	log.Info("Server Start")

	models.InitCache()

	services.MakeHwp(54)
}

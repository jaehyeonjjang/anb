package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var (
	Database         string
	Owner            string
	ConnectionString string
	Port             string
	TempPath         string
	Convert          string

	ServiceUrl string

	UploadPath string
	ImagePath  string

	SmsUser   string
	SmsKey    string
	SmsSender string

	AdminEmail string

	Version string
	Build   string
	DEBUG   uint64

	MailSender string
	MailHost   string
	MailPort   int
	MailUser   string
	MailPasswd string

	LocalMode string
)

func init() {
	UploadPath = "webdata"
	Database = "mysql"
	Port = "3001"
	LocalMode = "true"
	ServiceUrl = "http://anb.netb.co.kr"
	Convert = "convert"

	DEBUG = 0
	if os.Getenv("GIN_MODE") == "release" {
		DEBUG = 0
	}
	if DEBUG > 0 {
		fmt.Printf("Debug: MODE=true, flag=%+v \n", DEBUG)
	}

	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if value := viper.Get("connectionString"); value != nil {
		ConnectionString = value.(string)
	}

	if value := viper.Get("uploadPath"); value != nil {
		UploadPath = value.(string)
	}

	if value := viper.Get("port"); value != nil {
		Port = value.(string)
	}

	if value := viper.Get("mailSender"); value != nil {
		MailSender = value.(string)
	}

	if value := viper.Get("mailHost"); value != nil {
		MailHost = value.(string)
	}

	if value := viper.Get("mailPort"); value != nil {
		i, _ := strconv.Atoi(value.(string))
		MailPort = i
	}

	if value := viper.Get("mailUser"); value != nil {
		MailUser = value.(string)
	}

	if value := viper.Get("mailPasswd"); value != nil {
		MailPasswd = value.(string)
	}

	if value := viper.Get("smsUser"); value != nil {
		SmsUser = value.(string)
	}

	if value := viper.Get("smsKey"); value != nil {
		SmsKey = value.(string)
	}

	if value := viper.Get("smsSender"); value != nil {
		SmsSender = value.(string)
	}

	if value := viper.Get("adminEmail"); value != nil {
		AdminEmail = value.(string)
	}

	if value := viper.Get("convert"); value != nil {
		Convert = value.(string)
	}

	if value := viper.Get("imagePath"); value != nil {
		ImagePath = value.(string)
	}

	if value := viper.Get("localMode"); value != nil {
		LocalMode = value.(string)
	}

	if value := viper.Get("serviceUrl"); value != nil {
		ServiceUrl = value.(string)
	}

	/*
	log.Printf("config.Database = %v\n", Database)
	log.Printf("config.Owner = %v\n", Owner)
	log.Printf("config.ConnectionString = %v\n", ConnectionString)
	log.Printf("config.Port = %v\n", Port)
	log.Printf("config.TempPath = %v\n", TempPath)
	log.Printf("config.MailSender = %v\n", MailSender)
	log.Printf("config.ServiceUrl = %v\n", ServiceUrl)

	log.Printf("config.UploadPath = %v\n", UploadPath)

	log.Printf("config.ImagePath = %v\n", ImagePath)
	log.Printf("config.LocalMode = %v\n", LocalMode)

	log.Printf("config.Version = %v\n", Version)
	log.Printf("config.Build = %v\n", Build)
	log.Printf("config.DEBUG = %v\n", DEBUG)
	*/
}

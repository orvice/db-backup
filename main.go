package main

import (
	"time"

	"github.com/lukerodham/barkup"
	"github.com/orvice/kit/log"
	tk "github.com/orvice/kit/time"
)

var (
	logger log.Logger
)

func backup() {

	// Configure a MySQL exporter
	mysql := &barkup.MySQL{
		Host:     DB_Host,
		Port:     DB_Port,
		DB:       DB_DB,
		User:     DB_User,
		Password: DB_Password,
	}

	// Configure a S3 storer
	s3 := &barkup.S3{
		Region:       S3_Region,
		Bucket:       S3_Bucket,
		AccessKey:    S3_Key,
		ClientSecret: S3_Secret,
	}

	// Export the database, and send it to the
	// bucket in the `db_backups` folder
	result := mysql.Export(false)
	if result.Error != nil {
		logger.Errorf("back error : %v", result.Error.Error())
		return
	}

	err := result.To("db_backups/", s3)
	if err != nil {
		logger.Error(err)
	}
}

func main() {
	logger = log.NewDefaultLogger()
	InitEnv()
	for {
		logger.Infof("Start backup time %v", time.Now().Format(tk.DefaultTimeFormat))
		backup()
		time.Sleep(time.Minute * 10)
	}
}

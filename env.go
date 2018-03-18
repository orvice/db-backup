package main

import "github.com/orvice/utils/env"

var (
	S3_Key, S3_Secret, S3_Region, S3_Bucket string

	DB_Host, DB_User, DB_Password, DB_Port, DB_DB string

	Telegram_Token  string
	Telegram_ChatID int
)

func InitEnv() {
	S3_Key = env.Get("S3_KEY")
	S3_Secret = env.Get("S3_SECRET")
	S3_Region = env.Get("S3_REGION")
	S3_Bucket = env.Get("S3_BUCKET")

	DB_Host = env.Get("DB_HOST")
	DB_User = env.Get("DB_USER")
	DB_Password = env.Get("DB_PASSWORD")
	DB_Port = env.Get("DB_PORT")
	DB_DB = env.Get("DB_DB")
}

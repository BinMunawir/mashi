package configs

import "os"

var (
	HOST         = os.Getenv("HOST")
	PORT         = os.Getenv("PORT")
	DB_HOST      = os.Getenv("DB_HOST")
	DB_PORT      = os.Getenv("DB_PORT")
	DB_USER      = os.Getenv("DB_USER")
	DB_PASSWORRD = os.Getenv("DB_PASSWORRD")
	DB_NAME      = os.Getenv("DB_NAME")
	DNS          = "postgres://" + DB_USER + ":" + DB_PASSWORRD + "@" + DB_HOST + "/" + DB_NAME + "?sslmode=disable"
	KAFKA_HOST   = os.Getenv("KAFKA_HOST")
)

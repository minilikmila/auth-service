package main

import (
	"embed"
	"os"

	"auth.service/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	// "github.com/sirupsen/logrus"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
	// "gorm.io/gorm/logger"
	// "gorm.io/gorm/schema"
	// // "github.com/golang-migrate/migrate/v4"
	// "github.com/golang-migrate/migrate/v4"
	// "github.com/johejo/golang-migrate-extra/source/iofs"
	// // "github.com/google/martian/log"
)

var fs embed.FS

func main() {

	godotenv.Load()

	router := gin.Default()

	logrus.SetLevel(logrus.DebugLevel)

	port := os.Getenv("PORT")

	err := db.Init(fs)
	if err != nil {
		logrus.Info("error")
		logrus.Fatalln(err)
	}

	logrus.Fatalln(router.Run(":"+port))
}

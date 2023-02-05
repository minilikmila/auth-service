package db

import (
	"embed"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/johejo/golang-migrate-extra/source/iofs"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	// "gorm.io/driver/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func Init(fs embed.FS) error {
	godotenv.Load()

	logrus.SetLevel(logrus.DebugLevel)
	
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.SetOutput(os.Stdout)
	
	url := os.Getenv("URL")

	d, err := iofs.New(fs, "/migrations")
	if err != nil {
		logrus.Fatal(err)
	}

	dbb, err := gorm.Open(postgres.Open(url), &gorm.Config{
		NamingStrategy: schema.NamingStrategy {
			TablePrefix: "trust",
			SingularTable: false,
		},
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		logrus.Fatalln(err)
		return err
	}

	sql, err := dbb.DB()

	if err != nil {
		logrus.Fatalln(err)
		return err
	}

	sql.SetMaxIdleConns(20)



	m, err := migrate.NewWithSourceInstance("iofs", d, url)

	if err != nil {
		logrus.Fatalln(err)
		return err
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		logrus.Fatalln(err)
		return err
	}
 return nil
}
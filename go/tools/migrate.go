package main

import (
	"github.com/sirupsen/logrus"
	"github.com/tanu8282/hacking-to-the-YouTube-API/go/databases"
	"github.com/tanu8282/hacking-to-the-YouTube-API/go/models"
)

func main() {
	db, err := databases.Connect()
	defer db.Close()

	if err != nil {
		logrus.Fatal(err)
	}

	db.Debug().AutoMigrate(&models.User{})
	db.Debug().AutoMigrate(&models.Favorite{})
}

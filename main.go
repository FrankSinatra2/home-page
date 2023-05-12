package main

import (
	"github.com/FrankSinatra2/home-page/internal/controllers"
	"github.com/FrankSinatra2/home-page/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize Database
	db, err := gorm.Open(sqlite.Open("file:db/test.sqlite?_fk=on"), &gorm.Config{})

	if err != nil {
		panic("Failed to open database!")
	}

	if res := db.Exec("PRAGMA foreign_keys = ON", nil); res.Error != nil {
		panic("Faild to turn on foreign keys")
	}

	db.AutoMigrate(&models.ApplicationLauncher{})
	db.AutoMigrate(&models.Group{})

	db.FirstOrCreate(&models.Group{
		Title: "default",
		Icon:  "dataset",
	})

	// Create Controllers
	appLauncherCtrl := controllers.ApplicationLauncherController{
		DB: db,
	}

	groupCtrl := controllers.GroupController{
		DB: db,
	}

	// Initialize Enpoints
	r := gin.Default()
	r.POST("/v1/applicationLaunchers", appLauncherCtrl.CreateApplicationLauncher)
	r.GET("/v1/applicationLaunchers/:id", appLauncherCtrl.GetApplicationLauncher)

	r.PATCH("/v1/applicationLaunchers/:id", nil)

	r.POST("/v1/groups", groupCtrl.CreateGroup)
	r.GET("/v1/groups", groupCtrl.GetGroups)

	// Bind website
	r.Static("/home", "dist")

	r.Run(":3000")
}

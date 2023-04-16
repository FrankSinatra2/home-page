package main

import (
	"github.com/FrankSinatra2/home-page/internal/controllers"
	"github.com/FrankSinatra2/home-page/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateGroup(c *gin.Context) {

}

func main() {
	// Initialize Database
	db, err := gorm.Open(sqlite.Open("db/test.sqlite"), &gorm.Config{})

	if err != nil {
		panic("Failed to open database!")
	}

	db.AutoMigrate(&models.ApplicationLauncher{})
	db.AutoMigrate(&models.Group{})

	db.FirstOrCreate(&models.Group{
		Title:                "default",
		Icon:                 "dataset",
		ApplicationLaunchers: []models.ApplicationLauncher{},
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

	r.PATCH("/v1/applicationLaunchers/:id/move", nil)

	r.POST("/v1/groups", groupCtrl.CreateGroup)
	r.GET("/v1/groups", groupCtrl.GetGroups)

	// Bind website
	r.Static("/home", "dist")

	r.Run(":3000")
}

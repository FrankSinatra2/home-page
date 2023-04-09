package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/FrankSinatra2/home-page/pkg/contracts"
	"github.com/FrankSinatra2/home-page/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func CreateApplicationLauncherHandler(c *gin.Context) {
	body := contracts.CreateApplicationLauncher{}

	err := c.BindJSON(&body)

	if err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	model := models.ApplicationLauncher{
		Title:          body.Title,
		ApplicationUrl: body.ApplicationUrl,
		Icon:           body.Icon,
		GroupId:        body.GroupId,
	}

	result := db.Create(&model)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	res := contracts.GetApplicationLauncher{
		ID:             model.ID,
		Title:          model.Title,
		Icon:           model.Icon,
		ApplicationUrl: model.ApplicationUrl,
		GroupId:        model.GroupId,
		CreatedAt:      model.CreatedAt,
		UpdatedAt:      model.UpdatedAt,
	}

	c.JSON(http.StatusCreated, res)
}

func GetApplicationLauncherHandler(c *gin.Context) {
	id := c.Param("id")

	dbId, err := strconv.Atoi(id)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	model := models.ApplicationLauncher{}
	if err = db.Find(&model, dbId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound)
		} else {
			c.Status(http.StatusInternalServerError)
		}

		return
	}

	res := contracts.GetApplicationLauncher{
		ID:             model.ID,
		Title:          model.Title,
		Icon:           model.Icon,
		ApplicationUrl: model.ApplicationUrl,
		GroupId:        model.GroupId,
		CreatedAt:      model.CreatedAt,
		UpdatedAt:      model.UpdatedAt,
	}

	c.JSON(http.StatusCreated, res)
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

	// Initialize Enpoints
	r := gin.Default()
	r.POST("/v1/applicationLauncher", CreateApplicationLauncherHandler)
	r.POST("/v1/applicationLauncher/:id", GetApplicationLauncherHandler)

	// Bind website
	r.Static("/", "dist")

	r.Run(":3000")
}

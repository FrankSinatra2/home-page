package controllers

import (
	"fmt"
	"net/http"

	"github.com/FrankSinatra2/home-page/pkg/contracts"
	"github.com/FrankSinatra2/home-page/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GroupController struct {
	DB *gorm.DB
}

func (ctrl *GroupController) CreateGroup(c *gin.Context) {
	body := contracts.CreateGroup{}
	err := c.BindJSON(&body)

	if err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	model := models.Group{
		Title: body.Title,
		Icon:  body.Icon,
	}

	result := ctrl.DB.Create(&model)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		fmt.Printf("%v", result.Error.Error())
		return
	}

	res := contracts.GetGroup{
		ID:                   model.ID,
		Title:                model.Title,
		Icon:                 model.Icon,
		ApplicationLaunchers: []contracts.GetApplicationLauncher{},
		CreatedAt:            model.CreatedAt,
		UpdatedAt:            model.UpdatedAt,
	}

	c.JSON(http.StatusCreated, res)
}

func (ctrl *GroupController) GetGroups(c *gin.Context) {
	models := []models.Group{}

	if err := ctrl.DB.Find(&models).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	res := []contracts.GetGroup{}

	for _, model := range models {
		appLaunchers := []contracts.GetApplicationLauncher{}

		for _, appLauncher := range model.ApplicationLaunchers {
			appLaunchers = append(appLaunchers, contracts.GetApplicationLauncher{
				ID:             appLauncher.ID,
				Title:          appLauncher.Title,
				Icon:           appLauncher.Icon,
				ApplicationUrl: appLauncher.ApplicationUrl,
				GroupId:        appLauncher.GroupId,
				CreatedAt:      appLauncher.CreatedAt,
				UpdatedAt:      appLauncher.UpdatedAt,
			})
		}

		res = append(res, contracts.GetGroup{
			ID:                   model.ID,
			Title:                model.Title,
			ApplicationLaunchers: appLaunchers,
			CreatedAt:            model.CreatedAt,
			UpdatedAt:            model.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, res)
}

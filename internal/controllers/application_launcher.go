package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/FrankSinatra2/home-page/pkg/contracts"
	"github.com/FrankSinatra2/home-page/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ApplicationLauncherController struct {
	DB *gorm.DB
}

func (ctrl *ApplicationLauncherController) CreateApplicationLauncher(c *gin.Context) {
	body := contracts.CreateApplicationLauncher{}

	err := c.BindJSON(&body)

	if err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	fmt.Printf("body: %v", body)

	model := models.ApplicationLauncher{
		Title:          body.Title,
		ApplicationUrl: body.ApplicationUrl,
		Icon:           body.Icon,
		GroupId:        body.GroupId,
	}

	result := ctrl.DB.Create(&model)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		fmt.Printf("%v", result.Error.Error())
		return
	}

	fmt.Println("here")

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

func (ctrl *ApplicationLauncherController) GetApplicationLauncher(c *gin.Context) {
	id := c.Param("id")

	dbId, err := strconv.Atoi(id)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	model := models.ApplicationLauncher{}
	model.ID = uint(dbId)
	if err = ctrl.DB.First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound)
			fmt.Printf("t: %s", err.Error())
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

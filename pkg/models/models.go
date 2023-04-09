package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model

	Title string
	Icon  string

	ApplicationLaunchers []ApplicationLauncher
}

type ApplicationLauncher struct {
	gorm.Model

	Title          string
	Icon           string
	ApplicationUrl string

	GroupId uint
}

package contracts

import "time"

type CreateApplicationLauncher struct {
	Title          string `json:"title"`
	Icon           string `json:"icon"`
	ApplicationUrl string `json:"application_url"`
	GroupId        uint   `json:"group_id"`
}

type GetApplicationLauncher struct {
	ID             uint      `json:"id"`
	Title          string    `json:"title"`
	Icon           string    `json:"icon"`
	ApplicationUrl string    `json:"application_url"`
	GroupId        uint      `json:"group_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

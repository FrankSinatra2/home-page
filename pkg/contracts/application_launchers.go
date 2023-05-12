package contracts

import "time"

type CreateApplicationLauncher struct {
	Title          string `json:"title"`
	Icon           string `json:"icon"`
	ApplicationUrl string `json:"application_url"`
	GroupID        uint   `json:"group_id"`
}

type GetApplicationLauncher struct {
	ID             uint      `json:"id"`
	Title          string    `json:"title"`
	Icon           string    `json:"icon"`
	ApplicationUrl string    `json:"application_url"`
	GroupID        uint      `json:"group_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type PatchApplicationLauncher struct {
	Title          string `json:"title,omitempty"`
	Icon           string `json:"icon,omitempty"`
	ApplicationUrl string `json:"application_url,omitempty"`
	GroupID        uint   `json:"group_id,omitempty"`
}

package contracts

import "time"

type CreateGroup struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
}

type GetGroup struct {
	ID                   uint                     `json:"id"`
	Title                string                   `json:"title"`
	Icon                 string                   `json:"icon"`
	ApplicationLaunchers []GetApplicationLauncher `json:"application_launchers"`
	CreatedAt            time.Time                `json:"created_at"`
	UpdatedAt            time.Time                `json:"updated_at"`
}

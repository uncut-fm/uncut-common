package model

type User struct {
	Id                  int    `json:"user_id"`
	Name                string `json:"name"`
	Email               string `json:"email"`
	ProfileImageUrl     string `json:"profile_image_url"`
	HasAdminPanelAccess bool   `json:"has_admin_panel_access,omitempty"`
}

package model

type User struct {
	Id                  int    `json:"user_id"`
	Name                string `json:"name,omitempty"`
	Email               string `json:"email"`
	ProfileImageUrl     string `json:"profile_image_url,omitempty"`
	HasAdminPanelAccess bool   `json:"has_admin_panel_access"`
}

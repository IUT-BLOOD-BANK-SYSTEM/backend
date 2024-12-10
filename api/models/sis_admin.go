package models

type SisAdminPrimaryKey struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

type CreateSisAdmin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SisAdmin struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateSisAdmin struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SisAdminGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type SisAdminGetListResponse struct {
	Count     int         `json:"count"`
	SisAdmins []*SisAdmin `json:"sis_admins"`
}

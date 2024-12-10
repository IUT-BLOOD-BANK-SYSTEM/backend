package models

type BloodTypePrimaryKey struct {
	Id string `json:"id"`
}

type CreateBloodType struct {
	BloodsType string `json:"bloods_type"`
}

type BloodType struct {
	Id         string `json:"id"`
	BloodsType string `json:"bloods_type"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type UpdateBloodType struct {
	Id         string `json:"id"`
	BloodsType string `json:"bloods_type"`
}

type BloodTypeGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type BloodTypeGetListResponse struct {
	Count      int          `json:"count"`
	BloodTypes []*BloodType `json:"blood_types"`
}

package models

type BloodInventoryPrimaryKey struct {
	Id string `json:"id"`
}

type CreateBloodInventory struct {
	HospitalID  string  `json:"hospital_id"`
	BloodTypeID string  `json:"blood_type_id"`
	Quantity    float64 `json:"quantity"`
}

type BloodInventory struct {
	Id        string     `json:"id"`
	Hospital  *Hospital  `json:"hospital"`
	BloodType *BloodType `json:"blood_type"`
	Quantity  float64    `json:"quantity"`
	CreatedAt string     `json:"created_at"`
	UpdatedAt string     `json:"updated_at"`
}

type UpdateBloodInventory struct {
	Id          string  `json:"id"`
	HospitalID  string  `json:"hospital_id"`
	BloodTypeID string  `json:"blood_type_id"`
	Quantity    float64 `json:"quantity"`
}

type BloodInventoryGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type BloodInventoryGetListResponse struct {
	Count            int               `json:"count"`
	BloodInventories []*BloodInventory `json:"blood_inventories"`
}

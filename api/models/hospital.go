package models

type HospitalPrimaryKey struct {
	Id string `json:"id"`
	// Email string `json:"email"`
}

type CreateHospital struct {
	Name        string `json:"name"`
	RegionID    string `json:"region_id"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	HeadNurseID string `json:"head_nurse_id"`
}

type Hospital struct {
	Id          string     `json:"id"`
	Name        string     `json:"name"`
	Region      *Region    `json:"region"`
	Address     string     `json:"address"`
	PhoneNumber string     `json:"phone_number"`
	HeadNurse   *HeadNurse `json:"head_nurse"`
	CreatedAt   string     `json:"created_at"`
	UpdatedAt   string     `json:"updated_at"`
}

type UpdateHospital struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	RegionID    string `json:"region_id"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	HeadNurseID string `json:"head_nurse_id"`
}

type HospitalGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type HospitalGetListResponse struct {
	Count     int         `json:"count"`
	Hospitals []*Hospital `json:"hospitals"`
}

package models

type BloodRequestPrimaryKey struct {
	Id string `json:"id"`
}

type CreateBloodRequest struct {
	DoctorID    string  `json:"doctor_id"`
	HospitalID  string  `json:"hospital_id"`
	BloodTypeID string  `json:"blood_type_id"`
	Quantity    float64 `json:"quantity"`
	Status      string  `json:"status"`
}

type BloodRequest struct {
	Id        string     `json:"id"`
	Doctor    *Doctor    `json:"doctor"`
	Hospital  *Hospital  `json:"hospital"`
	BloodType *BloodType `json:"blood_type"`
	Quantity  float64    `json:"quantity"`
	Status    string     `json:"status"`
	CreatedAt string     `json:"created_at"`
	UpdatedAt string     `json:"updated_at"`
}

type UpdateBloodRequest struct {
	Id          string  `json:"id"`
	DoctorID    string  `json:"doctor_id"`
	HospitalID  string  `json:"hospital_id"`
	BloodTypeID string  `json:"blood_type_id"`
	Quantity    float64 `json:"quantity"`
	Status      string  `json:"status"`
}

type BloodRequestGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type BloodRequestGetListResponse struct {
	Count         int             `json:"count"`
	BloodRequests []*BloodRequest `json:"BloodRequests"`
}

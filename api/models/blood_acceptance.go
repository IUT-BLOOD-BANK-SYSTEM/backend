package models

type BloodAcceptancePrimaryKey struct {
	Id string `json:"id"`
}

type CreateBloodAcceptance struct {
	DoctorID       string `json:"doctor_id"`
	HospitalID     string `json:"hospital_id"`
	PatientID      string `json:"patient_id"`
	AcceptanceDate string `json:"acceptance_date"`
	Status         string `json:"status"`
}

type BloodAcceptance struct {
	Id             string    `json:"id"`
	Doctor         *Doctor   `json:"doctor"`
	Hospital       *Hospital `json:"hospital"`
	Patient        *User     `json:"patient"`
	AcceptanceDate string    `json:"acceptance_date"`
	Status         string    `json:"status"`
	CreatedAt      string    `json:"created_at"`
	UpdatedAt      string    `json:"updated_at"`
}

type UpdateBloodAcceptance struct {
	Id             string `json:"id"`
	DoctorID       string `json:"doctor_id"`
	HospitalID     string `json:"hospital_id"`
	PatientID      string `json:"patient_id"`
	AcceptanceDate string `json:"acceptance_date"`
	Status         string `json:"status"`
}

type BloodAcceptanceGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type BloodAcceptanceGetListResponse struct {
	Count            int                `json:"count"`
	BloodAcceptances []*BloodAcceptance `json:"blood_acceptances"`
}

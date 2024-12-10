package models

type BloodDonationPrimaryKey struct {
	Id string `json:"id"`
}

type CreateBloodDonation struct {
	DonorID      string  `json:"donor_id"`
	HospitalID   string  `json:"hospital_id"`
	BloodTypeID  string  `json:"blood_type_id"`
	Quantity     float64 `json:"quantity"`
	DonationDate string  `json:"donation_date"`
	Status       string  `json:"status"`
}

type BloodDonation struct {
	Id           string     `json:"id"`
	Donor        *User      `json:"donor"`
	Hospital     *Hospital  `json:"hospital"`
	BloodType    *BloodType `json:"blood_type"`
	Quantity     float64    `json:"quantity"`
	DonationDate string     `json:"donation_date"`
	Status       string     `json:"status"`
	CreatedAt    string     `json:"created_at"`
	UpdatedAt    string     `json:"updated_at"`
}

type UpdateBloodDonation struct {
	Id           string  `json:"id"`
	DonorID      string  `json:"donor_id"`
	HospitalID   string  `json:"hospital_id"`
	BloodTypeID  string  `json:"blood_type_id"`
	Quantity     float64 `json:"quantity"`
	DonationDate string  `json:"donation_date"`
	Status       string  `json:"status"`
}

type BloodDonationGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type BloodDonationGetListResponse struct {
	Count          int              `json:"count"`
	BloodDonations []*BloodDonation `json:"BloodDonations"`
}

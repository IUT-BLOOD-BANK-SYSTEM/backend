package models

type AppointmentPrimaryKey struct {
	Id string `json:"id"`
}

type CreateAppointment struct {
	DonorID         string `json:"donor_id"`
	HospitalID      string `json:"hospital_id"`
	AppointmentDate string `json:"appointment_date"`
	Status          string `json:"status"`
}

type Appointment struct {
	Id              string    `json:"id"`
	Donor           *User     `json:"donor"`
	Hospital        *Hospital `json:"hospital"`
	AppointmentDate string    `json:"appointment_date"`
	Status          string    `json:"status"`
	CreatedAt       string    `json:"created_at"`
	UpdatedAt       string    `json:"updated_at"`
}

type UpdateAppointment struct {
	Id              string `json:"id"`
	DonorID         string `json:"donor_id"`
	HospitalID      string `json:"hospital_id"`
	AppointmentDate string `json:"appointment_date"`
	Status          string `json:"status"`
}

type AppointmentGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type AppointmentGetListResponse struct {
	Count        int            `json:"count"`
	Appointments []*Appointment `json:"appointments"`
}

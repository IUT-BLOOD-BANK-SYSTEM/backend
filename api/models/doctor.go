package models

type DoctorPrimaryKey struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	//Doctorname string `json:"Doctorname"`
}

type CreateDoctor struct {
	FirstName      string `json:"first_name"`
	SecondName     string `json:"second_name"`
	HospitalID     string `json:"hospital_id"`
	Specialization string `json:"specialization"`
	PhoneNumber    string `json:"phone_number"`
	Address        string `json:"address"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	//Doctorname string `json:"Doctorname"`

}

type Doctor struct {
	Id             string    `json:"id"`
	FirstName      string    `json:"first_name"`
	SecondName     string    `json:"second_name"`
	Hospital       *Hospital `json:"hospital"`
	Specialization string    `json:"specialization"`
	PhoneNumber    string    `json:"phone_number"`
	Address        string    `json:"address"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	CreatedAt      string    `json:"created_at"`
	UpdatedAt      string    `json:"updated_at"`
}

type UpdateDoctor struct {
	Id             string `json:"id"`
	FirstName      string `json:"first_name"`
	SecondName     string `json:"second_name"`
	HospitalID     string `json:"hospital_id"`
	Specialization string `json:"specialization"`
	PhoneNumber    string `json:"phone_number"`
	Address        string `json:"address"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	//Doctorname string `json:"Doctorname"`

}

type DoctorGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type DoctorGetListResponse struct {
	Count   int       `json:"count"`
	Doctors []*Doctor `json:"doctors"`
}

package models

type HeadNursePrimaryKey struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	//HeadNursename string `json:"HeadNursename"`
}

type CreateHeadNurse struct {
	FirstName   string `json:"first_name"`
	SecondName  string `json:"second_name"`
	DateOfBirth string `json:"date_of_birth"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	HospitalID  string `json:"hospital"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	//HeadNursename string `json:"HeadNursename"`

}

type HeadNurse struct {
	Id          string    `json:"id"`
	FirstName   string    `json:"first_name"`
	SecondName  string    `json:"second_name"`
	DateOfBirth string    `json:"date_of_birth"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	Hospital    *Hospital `json:"hospital"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}

type UpdateHeadNurse struct {
	Id          string `json:"id"`
	FirstName   string `json:"first_name"`
	SecondName  string `json:"second_name"`
	DateOfBirth string `json:"date_of_birth"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	HospitalID  string `json:"hospital"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	//HeadNursename string `json:"HeadNursename"`

}

type HeadNurseGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type HeadNurseGetListResponse struct {
	Count      int          `json:"count"`
	HeadNurses []*HeadNurse `json:"head_nurses"`
}

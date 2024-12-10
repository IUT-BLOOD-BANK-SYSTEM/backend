package models

type UserPrimaryKey struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	//Username string `json:"username"`
}

type CreateUser struct {
	FirstName      string `json:"first_name"`
	SecondName     string `json:"second_name"`
	DateOfBirth    string `json:"date_of_birth"`
	PhoneNumber    string `json:"phone_number"`
	Address        string `json:"address"`
	BloodTypeID    string `json:"blood_type_id"`
	Diseases       string `json:"diseases"`
	PassportNumber string `json:"passport_number"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	//Username string `json:"username"`

}

type User struct {
	Id             string     `json:"id"`
	FirstName      string     `json:"first_name"`
	SecondName     string     `json:"second_name"`
	DateOfBirth    string     `json:"date_of_birth"`
	PhoneNumber    string     `json:"phone_number"`
	Address        string     `json:"address"`
	BloodType      *BloodType `json:"blood_type"`
	Diseases       string     `json:"diseases"`
	PassportNumber string     `json:"passport_number"`
	Email          string     `json:"email"`
	Password       string     `json:"password"`
	CreatedAt      string     `json:"created_at"`
	UpdatedAt      string     `json:"updated_at"`
}

type UpdateUser struct {
	Id             string `json:"id"`
	FirstName      string `json:"first_name"`
	SecondName     string `json:"second_name"`
	DateOfBirth    string `json:"date_of_birth"`
	PhoneNumber    string `json:"phone_number"`
	Address        string `json:"address"`
	BloodTypeID    string `json:"blood_type_id"`
	Diseases       string `json:"diseases"`
	PassportNumber string `json:"passport_number"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	//Username string `json:"username"`

}

type UserGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type UserGetListResponse struct {
	Count int     `json:"count"`
	Users []*User `json:"users"`
}

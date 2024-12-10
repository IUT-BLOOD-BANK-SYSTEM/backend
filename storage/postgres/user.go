package postgres

import (
	"context"
	"database/sql"
	"fmt"

	uuid "github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"app/api/models"
	"app/pkg/helper"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(ctx context.Context, req *models.CreateUser) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO users(id, first_name, second_name,date_of_birth, phone_number,address,blood_type_id,diseases,passport_number,email,password,updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.FirstName,
		req.SecondName,
		req.DateOfBirth,
		req.PhoneNumber,
		req.Address,
		req.BloodTypeID,
		req.Diseases,
		req.PassportNumber,
		req.Email,
		req.Password,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *userRepo) GetByID(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error) {

	var whereField = "u.id"
	if len(req.Email) > 0 {
		whereField = "email"
		req.Id = req.Email
	}

	var (
		query string

		id              sql.NullString
		first_name      sql.NullString
		second_name     sql.NullString
		date_of_birth   sql.NullString
		phone_number    sql.NullString
		address         sql.NullString
		blood_type_id   sql.NullString
		bloods_type     sql.NullString
		diseases        sql.NullString
		passport_number sql.NullString
		email           sql.NullString
		password        sql.NullString
		createdAt       sql.NullString
		updatedAt       sql.NullString
	)

	query = `
		SELECT
			u.id,
			u.first_name,
			u.second_name,
			u.date_of_birth,
			u.phone_number,
			u.address,
			u.blood_type_id,
			bt.bloods_type,
			u.diseases,
			u.passport_number,
			u.email,
			u.password,
			u.created_at,
			u.updated_at
		FROM users u
		LEFT JOIN blood_types bt ON u.blood_type_id = bt.id
		WHERE ` + whereField + ` = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&first_name,
		&second_name,
		&date_of_birth,
		&phone_number,
		&address,
		&blood_type_id,
		&bloods_type,
		&diseases,
		&passport_number,
		&email,
		&password,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.User{
		Id:          id.String,
		FirstName:   first_name.String,
		SecondName:  second_name.String,
		DateOfBirth: date_of_birth.String,
		PhoneNumber: phone_number.String,
		Address:     address.String,
		BloodType: &models.BloodType{
			Id:         blood_type_id.String,
			BloodsType: bloods_type.String,
		},
		Diseases:       diseases.String,
		PassportNumber: passport_number.String,
		Email:          email.String,
		Password:       password.String,
		CreatedAt:      createdAt.String,
		UpdatedAt:      updatedAt.String,
	}, nil
}

func (r *userRepo) GetList(ctx context.Context, req *models.UserGetListRequest) (*models.UserGetListResponse, error) {

	var (
		resp   = &models.UserGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
			SELECT
			COUNT(*) OVER(),
			u.id,
			u.first_name,
			u.second_name,
			u.date_of_birth,
			u.phone_number,
			u.address,
			u.blood_type_id,
			bt.bloods_type,
			u.diseases,
			u.passport_number,
			u.email,
			u.password,
			u.created_at,
			u.updated_at
		FROM users u
		LEFT JOIN blood_types bt ON u.blood_type_id = bt.id
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND first_name ILIKE '%' || '` + req.Search + `' || '%'`
	}

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id              sql.NullString
			first_name      sql.NullString
			second_name     sql.NullString
			date_of_birth   sql.NullString
			phone_number    sql.NullString
			address         sql.NullString
			blood_type_id   sql.NullString
			bloods_type     sql.NullString
			diseases        sql.NullString
			passport_number sql.NullString
			email           sql.NullString
			password        sql.NullString
			createdAt       sql.NullString
			updatedAt       sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&first_name,
			&second_name,
			&date_of_birth,
			&phone_number,
			&address,
			&blood_type_id,
			&bloods_type,
			&diseases,
			&passport_number,
			&email,
			&password,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users, &models.User{
			Id:          id.String,
			FirstName:   first_name.String,
			SecondName:  second_name.String,
			DateOfBirth: date_of_birth.String,
			PhoneNumber: phone_number.String,
			Address:     address.String,
			BloodType: &models.BloodType{
				Id:         blood_type_id.String,
				BloodsType: bloods_type.String,
			},
			Diseases:       diseases.String,
			PassportNumber: passport_number.String,
			Email:          email.String,
			Password:       password.String,
			CreatedAt:      createdAt.String,
			UpdatedAt:      updatedAt.String,
		})
	}

	return resp, nil
}

func (r *userRepo) Update(ctx context.Context, req *models.UpdateUser) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			users
		SET
		id = :id,
		first_name = :first_name,
		second_name = :second_name,
		date_of_birth = :date_of_birth,
		phone_number = :phone_number,
		address = :address,
		blood_type_id = :blood_type_id,
		diseases = :diseases,
		passport_number = :passport_number,
		email = :email,
		password = :password,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":              req.Id,
		"first_name":      req.FirstName,
		"second_name":     req.SecondName,
		"date_of_birth":   req.DateOfBirth,
		"phone_number":    req.PhoneNumber,
		"address":         req.Address,
		"blood_type_id":   req.BloodTypeID,
		"diseases":        req.Diseases,
		"passport_number": req.PassportNumber,
		"email":           req.Email,
		"password":        req.Password,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *userRepo) Delete(ctx context.Context, req *models.UserPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM users WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}

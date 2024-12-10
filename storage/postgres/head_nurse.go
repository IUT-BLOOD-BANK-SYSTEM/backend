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

type head_nurseRepo struct {
	db *pgxpool.Pool
}

func NewHeadNurseRepo(db *pgxpool.Pool) *head_nurseRepo {
	return &head_nurseRepo{
		db: db,
	}
}

func (r *head_nurseRepo) Create(ctx context.Context, req *models.CreateHeadNurse) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO head_nurses(id, first_name, second_name,date_of_birth, phone_number,address,hospital_id,email,password,updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.FirstName,
		req.SecondName,
		req.DateOfBirth,
		req.PhoneNumber,
		req.Address,
		req.HospitalID,
		req.Email,
		req.Password,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *head_nurseRepo) GetByID(ctx context.Context, req *models.HeadNursePrimaryKey) (*models.HeadNurse, error) {

	var whereField = "hn.id"
	if len(req.Email) > 0 {
		whereField = "email"
		req.Id = req.Email
	}

	var (
		query string

		id               sql.NullString
		first_name       sql.NullString
		second_name      sql.NullString
		date_of_birth    sql.NullString
		phone_number     sql.NullString
		address          sql.NullString
		hospital_id      sql.NullString
		hospital_name    sql.NullString
		region_id        sql.NullString
		region_name      sql.NullString
		hospital_address sql.NullString
		hospital_phone   sql.NullString
		email            sql.NullString
		password         sql.NullString
		createdAt        sql.NullString
		updatedAt        sql.NullString
	)

	query = `
		SELECT
		hn.id,
		hn.first_name,
		hn.second_name,
		hn.date_of_birth,
		hn.phone_number,
		hn.address,
		hn.email,
		hn.hospital_id,
		h.name AS hospital_name,
		h.region_id,
		r.name AS region_name,
		h.address AS hospital_address,
		h.phone_number AS hospital_phone
	FROM head_nurses hn
	LEFT JOIN hospitals h ON hn.hospital_id = h.id
	LEFT JOIN regions r ON h.region_id = r.id
	WHERE ` + whereField + ` = $1

	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&first_name,
		&second_name,
		&date_of_birth,
		&phone_number,
		&address,
		&email,
		&hospital_id,
		&hospital_name,
		&region_id,
		&region_name,
		&hospital_address,
		&hospital_phone,
	)

	if err != nil {
		return nil, err
	}

	return &models.HeadNurse{
		Id:          id.String,
		FirstName:   first_name.String,
		SecondName:  second_name.String,
		DateOfBirth: date_of_birth.String,
		PhoneNumber: phone_number.String,
		Address:     address.String,
		Hospital: &models.Hospital{
			Id:   hospital_id.String,
			Name: hospital_name.String,
			Region: &models.Region{
				Id:   region_id.String,
				Name: region_name.String,
			},
			Address:     hospital_address.String,
			PhoneNumber: hospital_phone.String,
		},
		Email:     email.String,
		Password:  password.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}, nil
}

func (r *head_nurseRepo) GetList(ctx context.Context, req *models.HeadNurseGetListRequest) (*models.HeadNurseGetListResponse, error) {

	var (
		resp   = &models.HeadNurseGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
    SELECT
        hn.id,
        hn.first_name,
        hn.second_name,
        hn.date_of_birth,
        hn.phone_number,
        hn.address,
        hn.email,
        hn.hospital_id,
        h.name AS hospital_name,
        h.region_id,
        r.name AS region_name,
        h.address AS hospital_address,
        h.phone_number AS hospital_phone,
        h.created_at AS hospital_created_at,
        h.updated_at AS hospital_updated_at
    FROM head_nurses hn
    LEFT JOIN hospitals h ON hn.hospital_id = h.id
    LEFT JOIN regions r ON h.region_id = r.id
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
			id               sql.NullString
			first_name       sql.NullString
			second_name      sql.NullString
			date_of_birth    sql.NullString
			phone_number     sql.NullString
			address          sql.NullString
			email            sql.NullString
			hospital_id      sql.NullString
			hospital_name    sql.NullString
			region_id        sql.NullString
			region_name      sql.NullString
			hospital_address sql.NullString
			hospital_phone   sql.NullString
			createdAt        sql.NullString
			updatedAt        sql.NullString
		)

		err := rows.Scan(
			&id,
			&first_name,
			&second_name,
			&date_of_birth,
			&phone_number,
			&address,
			&email,
			&hospital_id,
			&hospital_name,
			&region_id,
			&region_name,
			&hospital_address,
			&hospital_phone,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.HeadNurses = append(resp.HeadNurses, &models.HeadNurse{
			Id:          id.String,
			FirstName:   first_name.String,
			SecondName:  second_name.String,
			DateOfBirth: date_of_birth.String,
			PhoneNumber: phone_number.String,
			Address:     address.String,
			Hospital: &models.Hospital{
				Id:   hospital_id.String,
				Name: hospital_name.String,
				Region: &models.Region{
					Id:   region_id.String,
					Name: region_name.String,
				},
				Address:     hospital_address.String,
				PhoneNumber: hospital_phone.String,
			},
			Email:     email.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return resp, nil
}

func (r *head_nurseRepo) Update(ctx context.Context, req *models.UpdateHeadNurse) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			head_nurses
		SET
		id = :id,
		first_name = :first_name,
		second_name = :second_name,
		date_of_birth = :date_of_birth,
		phone_number = :phone_number,
		address = :address,
		hospital_id = :hospital_id,
		email = :email,
		password = :password,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":            req.Id,
		"first_name":    req.FirstName,
		"second_name":   req.SecondName,
		"date_of_birth": req.DateOfBirth,
		"phone_number":  req.PhoneNumber,
		"address":       req.Address,
		"hospital_id":   req.HospitalID,
		"email":         req.Email,
		"password":      req.Password,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *head_nurseRepo) Delete(ctx context.Context, req *models.HeadNursePrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM head_nurses WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}

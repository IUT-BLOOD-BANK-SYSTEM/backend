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

type doctorRepo struct {
	db *pgxpool.Pool
}

func NewDoctorRepo(db *pgxpool.Pool) *doctorRepo {
	return &doctorRepo{
		db: db,
	}
}

func (r *doctorRepo) Create(ctx context.Context, req *models.CreateDoctor) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO doctors(id, first_name, second_name,hospital_id,specialization, phone_number,address,email,password,updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8,$9, NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.FirstName,
		req.SecondName,
		req.HospitalID,
		req.Specialization,
		req.PhoneNumber,
		req.Address,
		req.Email,
		req.Password,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *doctorRepo) GetByID(ctx context.Context, req *models.DoctorPrimaryKey) (*models.Doctor, error) {

	var whereField = "d.id"
	if len(req.Email) > 0 {
		whereField = "email"
		req.Id = req.Email
	}

	var (
		query string

		id                       sql.NullString
		first_name               sql.NullString
		second_name              sql.NullString
		hospital_id              sql.NullString
		hospital_name            sql.NullString
		hospital_address         sql.NullString
		hospital_phone           sql.NullString
		region_id                sql.NullString
		region_name              sql.NullString
		specialization           sql.NullString
		phone_number             sql.NullString
		address                  sql.NullString
		email                    sql.NullString
		password                 sql.NullString
		created_at               sql.NullString
		updated_at               sql.NullString
		head_nurse_id            sql.NullString
		head_nurse_first_name    sql.NullString
		head_nurse_second_name   sql.NullString
		head_nurse_date_of_birth sql.NullString
		head_nurse_phone         sql.NullString
		head_nurse_email         sql.NullString
	)

	query = `
		SELECT
			d.id,
			d.first_name,
			d.second_name,
			h.id AS hospital_id,
			h.name AS hospital_name,
			h.address AS hospital_address,
			h.phone_number AS hospital_phone,
			r.id AS region_id,
			r.name AS region_name,
			d.specialization,
			d.phone_number,
			d.address,
			d.email,
			d.password,
			d.created_at,
			d.updated_at,
			hn.id AS head_nurse_id,
		   hn.first_name AS head_nurse_first_name,
		   hn.second_name AS head_nurse_second_name,
		   hn.date_of_birth AS head_nurse_date_of_birth,
		   hn.phone_number AS head_nurse_phone,
		   hn.email AS head_nurse_email
		FROM doctors d
		LEFT JOIN hospitals h ON d.hospital_id = h.id
	   LEFT JOIN regions r ON h.region_id = r.id
	   LEFT JOIN head_nurses hn ON h.head_nurse_id = hn.id
		WHERE ` + whereField + ` = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&first_name,
		&second_name,
		&hospital_id,
		&hospital_name,
		&hospital_address,
		&hospital_phone,
		&region_id,
		&region_name,
		&specialization,
		&phone_number,
		&address,
		&email,
		&password,
		&created_at,
		&updated_at,
		&head_nurse_id,
		&head_nurse_first_name,
		&head_nurse_second_name,
		&head_nurse_date_of_birth,
		&head_nurse_phone,
		&head_nurse_email,
	)

	if err != nil {
		return nil, err
	}

	return &models.Doctor{
		Id:         id.String,
		FirstName:  first_name.String,
		SecondName: second_name.String,
		Hospital: &models.Hospital{
			Id:          hospital_id.String,
			Name:        hospital_name.String,
			Address:     hospital_address.String,
			PhoneNumber: hospital_phone.String,
			Region: &models.Region{
				Id:   region_id.String,
				Name: region_name.String,
			},
			HeadNurse: &models.HeadNurse{
				Id:         head_nurse_id.String,
				FirstName:  head_nurse_first_name.String,
				SecondName: head_nurse_second_name.String,
				Hospital: &models.Hospital{
					Id:          hospital_id.String,
					Name:        hospital_name.String,
					PhoneNumber: hospital_phone.String,
					Address:     hospital_address.String,
					Region: &models.Region{
						Id:   region_id.String,
						Name: region_name.String,
					},
				},
				DateOfBirth: head_nurse_date_of_birth.String,
				PhoneNumber: head_nurse_phone.String,
				Email:       head_nurse_email.String,
			},
		},
		Specialization: specialization.String,
		PhoneNumber:    phone_number.String,
		Address:        address.String,
		Email:          email.String,
		Password:       password.String,
		CreatedAt:      created_at.String,
		UpdatedAt:      updated_at.String,
	}, nil
}

func (r *doctorRepo) GetList(ctx context.Context, req *models.DoctorGetListRequest) (*models.DoctorGetListResponse, error) {

	var (
		resp   = &models.DoctorGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			d.id,
			d.first_name,
			d.second_name,
			h.id AS hospital_id,
			h.name AS hospital_name,
			h.address AS hospital_address,
			h.phone_number AS hospital_phone,
			r.id AS region_id,
			r.name AS region_name,
			d.specialization,
			d.phone_number,
			d.address,
			d.email,
			d.password,
			d.created_at,
			d.updated_at,
			hn.id AS head_nurse_id,
		   hn.first_name AS head_nurse_first_name,
		   hn.second_name AS head_nurse_second_name,
		   hn.date_of_birth AS head_nurse_date_of_birth,
		   hn.phone_number AS head_nurse_phone,
		   hn.email AS head_nurse_email
		FROM doctors d
		LEFT JOIN hospitals h ON d.hospital_id = h.id
	   LEFT JOIN regions r ON h.region_id = r.id
	   LEFT JOIN head_nurses hn ON h.head_nurse_id = hn.id
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
			id                       sql.NullString
			first_name               sql.NullString
			second_name              sql.NullString
			hospital_id              sql.NullString
			hospital_name            sql.NullString
			hospital_address         sql.NullString
			hospital_phone           sql.NullString
			region_id                sql.NullString
			region_name              sql.NullString
			specialization           sql.NullString
			phone_number             sql.NullString
			address                  sql.NullString
			email                    sql.NullString
			password                 sql.NullString
			created_at               sql.NullString
			updated_at               sql.NullString
			head_nurse_id            sql.NullString
			head_nurse_first_name    sql.NullString
			head_nurse_second_name   sql.NullString
			head_nurse_date_of_birth sql.NullString
			head_nurse_phone         sql.NullString
			head_nurse_email         sql.NullString
		)

		err := rows.Scan(
			&id,
			&first_name,
			&second_name,
			&hospital_id,
			&hospital_name,
			&hospital_address,
			&hospital_phone,
			&region_id,
			&region_name,
			&specialization,
			&phone_number,
			&address,
			&email,
			&password,
			&created_at,
			&updated_at,
			&head_nurse_id,
			&head_nurse_first_name,
			&head_nurse_second_name,
			&head_nurse_date_of_birth,
			&head_nurse_phone,
			&head_nurse_email,
		)

		if err != nil {
			return nil, err
		}

		resp.Doctors = append(resp.Doctors, &models.Doctor{
			Id:         id.String,
			FirstName:  first_name.String,
			SecondName: second_name.String,
			Hospital: &models.Hospital{
				Id:          hospital_id.String,
				Name:        hospital_name.String,
				Address:     hospital_address.String,
				PhoneNumber: hospital_phone.String,
				Region: &models.Region{
					Id:   region_id.String,
					Name: region_name.String,
				},
				HeadNurse: &models.HeadNurse{
					Id:         head_nurse_id.String,
					FirstName:  head_nurse_first_name.String,
					SecondName: head_nurse_second_name.String,
					Hospital: &models.Hospital{
						Id:          hospital_id.String,
						Name:        hospital_name.String,
						PhoneNumber: hospital_phone.String,
						Address:     hospital_address.String,
						Region: &models.Region{
							Id:   region_id.String,
							Name: region_name.String,
						},
					},
					DateOfBirth: head_nurse_date_of_birth.String,
					PhoneNumber: head_nurse_phone.String,
					Email:       head_nurse_email.String,
				},
			},
			Specialization: specialization.String,
			PhoneNumber:    phone_number.String,
			Address:        address.String,
			Email:          email.String,
			Password:       password.String,
			CreatedAt:      created_at.String,
			UpdatedAt:      updated_at.String,
		})
	}

	return resp, nil
}

func (r *doctorRepo) Update(ctx context.Context, req *models.UpdateDoctor) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			doctors
		SET
		id = :id,
		first_name = :first_name,
		second_name = :second_name,
		hospital_id = :hospital_id,
		specialization = :specialization,
		phone_number = :phone_number,
		address = :address,
		email = :email,
		password = :password,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":             req.Id,
		"first_name":     req.FirstName,
		"second_name":    req.SecondName,
		"hospital_id":    req.HospitalID,
		"specialization": req.Specialization,
		"phone_number":   req.PhoneNumber,
		"address":        req.Address,
		"email":          req.Email,
		"password":       req.Password,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *doctorRepo) Delete(ctx context.Context, req *models.DoctorPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM doctors WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}

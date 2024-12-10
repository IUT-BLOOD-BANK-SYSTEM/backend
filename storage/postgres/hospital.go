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

type hospitalRepo struct {
	db *pgxpool.Pool
}

func NewHospitalRepo(db *pgxpool.Pool) *hospitalRepo {
	return &hospitalRepo{
		db: db,
	}
}

func (r *hospitalRepo) Create(ctx context.Context, req *models.CreateHospital) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	var headNurseID interface{}
	if req.HeadNurseID == "" {
		headNurseID = nil // This will set head_nurse_id to NULL
	} else {
		headNurseID = req.HeadNurseID // Use the provided HeadNurseID
	}

	query = `
		INSERT INTO hospitals(id, name, region_id,address, phone_number,head_nurse_id,updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Name,
		req.RegionID,
		req.Address,
		req.PhoneNumber,
		headNurseID,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *hospitalRepo) GetByID(ctx context.Context, req *models.HospitalPrimaryKey) (*models.Hospital, error) {

	// var whereField = "id"
	// if len(req.Email) > 0 {
	// 	whereField = "email"
	// 	req.Id = req.Email
	// }

	// var (
	// 	query string

	// 	id            sql.NullString
	// 	name          sql.NullString
	// 	region_id     sql.NullString
	// 	address       sql.NullString
	// 	phone_number  sql.NullString
	// 	head_nurse_id sql.NullString
	// 	createdAt     sql.NullString
	// 	updatedAt     sql.NullString
	// )

	var (
		query string

		id                 sql.NullString
		name               sql.NullString
		region_id          sql.NullString
		region_name        sql.NullString
		address            sql.NullString
		phone_number       sql.NullString
		head_nurse_id      sql.NullString
		first_name         sql.NullString
		second_name        sql.NullString
		date_of_birth      sql.NullString
		head_nurse_phone   sql.NullString
		head_nurse_address sql.NullString
		email              sql.NullString
		// password           sql.NullString
		created_at sql.NullString
		updated_at sql.NullString
	)

	// query = `
	// 	SELECT
	// 		id,
	// 		name,
	// 		region_id,
	// 		address,
	// 		phone_number,
	// 		head_nurse_id,
	// 		created_at,
	// 		updated_at
	// 	FROM hospitals
	// 	WHERE ` + whereField + ` = $1
	// `

	query = `
		SELECT
	    h.id,
	    h.name,
	    r.id AS region_id,
		r.name AS region_name,
	    h.address,
	    h.phone_number,
	    hn.id AS head_nurse_id,
	    hn.first_name,
	    hn.second_name,
	    hn.date_of_birth,
	    hn.phone_number AS head_nurse_phone,
	    hn.address AS head_nurse_address,
	    hn.email,
	    h.created_at,
	    h.updated_at
	FROM hospitals h
	LEFT JOIN head_nurses hn ON h.head_nurse_id = hn.id
	LEFT JOIN regions r ON h.region_id = r.id

	WHERE h.id = $1

	` //    hn.password,

	// err := r.db.QueryRow(ctx, query, req.Id).Scan(
	// 	&id,
	// 	&name,
	// 	&region_id,
	// 	&address,
	// 	&phone_number,
	// 	&head_nurse_id,
	// 	&createdAt,
	// 	&updatedAt,
	// )

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&region_id,
		&region_name,
		&address,
		&phone_number,
		&head_nurse_id,
		&first_name,
		&second_name,
		&date_of_birth,
		&head_nurse_phone,
		&head_nurse_address,
		&email,
		// &password,
		&created_at,
		&updated_at,
	)

	if err != nil {
		return nil, err
	}

	return &models.Hospital{
		Id:   id.String,
		Name: name.String,
		Region: &models.Region{
			Id:   region_id.String,
			Name: region_name.String,
		},
		Address:     address.String,
		PhoneNumber: phone_number.String,
		HeadNurse: &models.HeadNurse{
			Id:          head_nurse_id.String,
			FirstName:   first_name.String,
			SecondName:  second_name.String,
			DateOfBirth: date_of_birth.String,
			PhoneNumber: phone_number.String,
			Address:     address.String,
			Email:       email.String,
			// Password:    password.String,
		},
		CreatedAt: created_at.String,
		UpdatedAt: updated_at.String,
	}, nil
}

func (r *hospitalRepo) GetList(ctx context.Context, req *models.HospitalGetListRequest) (*models.HospitalGetListResponse, error) {
	var (
		resp   = &models.HospitalGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		err    error
	)

	query = `
        SELECT 
            h.id,
            h.name,
            r.id AS region_id,
		    r.name AS region_name,
            h.address,
            h.phone_number,
            hn.id AS head_nurse_id,
            hn.first_name AS head_nurse_first_name,
            hn.second_name AS head_nurse_second_name,
			hn.date_of_birth,
            hn.phone_number AS head_nurse_phone,
            hn.address AS head_nurse_address,
            hn.email AS head_nurse_email,
            h.created_at,
            h.updated_at
        FROM hospitals h
        LEFT JOIN head_nurses hn ON h.head_nurse_id = hn.id
		LEFT JOIN regions r ON h.region_id = r.id
    `
	//hn.password AS password,
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

	defer rows.Close() // Ensure rows are closed properly

	for rows.Next() {
		var (
			id                     sql.NullString
			name                   sql.NullString
			region_id              sql.NullString
			region_name            sql.NullString
			address                sql.NullString
			phone_number           sql.NullString
			head_nurse_id          sql.NullString
			head_nurse_first_name  sql.NullString
			head_nurse_second_name sql.NullString
			date_of_birth          sql.NullString
			head_nurse_phone       sql.NullString
			head_nurse_address     sql.NullString
			head_nurse_email       sql.NullString
			// password               sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&id,
			&name,
			&region_id,
			&region_name,
			&address,
			&phone_number,
			&head_nurse_id,
			&head_nurse_first_name,
			&head_nurse_second_name,
			&date_of_birth,
			&head_nurse_phone,
			&head_nurse_address,
			&head_nurse_email,
			// &password,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Hospitals = append(resp.Hospitals, &models.Hospital{
			Id:   id.String,
			Name: name.String,
			Region: &models.Region{
				Id:   region_id.String,
				Name: region_id.String,
			},
			Address:     address.String,
			PhoneNumber: phone_number.String,
			HeadNurse: &models.HeadNurse{
				Id:          head_nurse_id.String,
				FirstName:   head_nurse_first_name.String,
				SecondName:  head_nurse_second_name.String,
				DateOfBirth: date_of_birth.String,
				PhoneNumber: head_nurse_phone.String,
				Address:     head_nurse_address.String,
				Email:       head_nurse_email.String,
				// Password:    password.String,
			},
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return resp, nil
}

func (r *hospitalRepo) Update(ctx context.Context, req *models.UpdateHospital) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			hospitals
		SET
		id = :id,
		name = :name,
		region_id = :region_id,
		address = :address,
		phone_number = :phone_number,
		head_nurse_id = :head_nurse_id,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":            req.Id,
		"name":          req.Name,
		"region_id":     req.RegionID,
		"address":       req.Address,
		"phone_number":  req.PhoneNumber,
		"head_nurse_id": req.HeadNurseID,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *hospitalRepo) Delete(ctx context.Context, req *models.HospitalPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM hospitals WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}

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

type blood_requestRepo struct {
	db *pgxpool.Pool
}

func NewBloodRequestRepo(db *pgxpool.Pool) *blood_requestRepo {
	return &blood_requestRepo{
		db: db,
	}
}

func (r *blood_requestRepo) Create(ctx context.Context, req *models.CreateBloodRequest) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO blood_requests(id, doctor_id,hospital_id,blood_type_id, quantity, status, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6,NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.DoctorID,
		req.HospitalID,
		req.BloodTypeID,
		req.Quantity,
		req.Status,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *blood_requestRepo) GetByID(ctx context.Context, req *models.BloodRequestPrimaryKey) (*models.BloodRequest, error) {

	var whereField = "br.id"
	var (
		query string

		id                       sql.NullString
		doctor_id                sql.NullString
		doctor_first_name        sql.NullString
		doctor_second_name       sql.NullString
		doctor_specialization    sql.NullString
		doctor_address           sql.NullString
		doctor_phone_number      sql.NullString
		doctor_email             sql.NullString
		hospital_id              sql.NullString
		hospital_name            sql.NullString
		hospital_address         sql.NullString
		hospital_phone           sql.NullString
		blood_type_id            sql.NullString
		requests_blood_type      sql.NullString
		region_id                sql.NullString
		region_name              sql.NullString
		quantity                 sql.NullFloat64
		status                   sql.NullString
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
		br.id,
		d.id AS doctor_id,
	    d.first_name AS doctor_first_name,
		d.second_name AS doctor_second_name,
		d.specialization AS doctor_specialization,
		d.address AS doctor_address,
		d.email AS doctor_email,
		d.phone_number AS doctor_phone_number,
		h.id AS hospital_id,
		h.name AS hospital_name,
		h.address AS hospital_address,
		h.phone_number AS hospital_phone,
		r.id AS region_id,
		r.name AS region_name,
		bt.id AS blood_type_id,
		bt.bloods_type AS requests_blood_type,
		br.quantity,
		br.status,
        br.created_at,
		br.updated_at,
		hn.id AS head_nurse_id,
		   hn.first_name AS head_nurse_first_name,
		   hn.second_name AS head_nurse_second_name,
		   hn.date_of_birth AS head_nurse_date_of_birth,
		   hn.phone_number AS head_nurse_phone,
		   hn.email AS head_nurse_email
		FROM blood_requests br
		LEFT JOIN hospitals h ON br.hospital_id = h.id
		LEFT JOIN regions r ON h.region_id = r.id
		LEFT JOIN doctors d ON br.doctor_id = d.id
		LEFT JOIN blood_types bt ON br.blood_type_id = bt.id
	   LEFT JOIN head_nurses hn ON h.head_nurse_id = hn.id
		WHERE ` + whereField + ` = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&doctor_id,
		&doctor_first_name,
		&doctor_second_name,
		&doctor_specialization,
		&doctor_address,
		&doctor_phone_number,
		&doctor_email,
		&hospital_id,
		&hospital_name,
		&hospital_address,
		&hospital_phone,
		&region_id,
		&region_name,
		&blood_type_id,
		&requests_blood_type,
		&quantity,
		&status,
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

	return &models.BloodRequest{
		Id: id.String,
		Doctor: &models.Doctor{
			Id:         doctor_id.String,
			FirstName:  doctor_first_name.String,
			SecondName: doctor_second_name.String,
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
						Address:     hospital_address.String,
						PhoneNumber: hospital_phone.String,
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
			Specialization: doctor_specialization.String,
			Address:        doctor_address.String,
			PhoneNumber:    doctor_phone_number.String,
			Email:          doctor_email.String,
		},
		Hospital: &models.Hospital{
			Id:          hospital_id.String,
			Name:        hospital_name.String,
			Address:     hospital_address.String,
			PhoneNumber: hospital_phone.String,
			Region: &models.Region{
				Id:   region_id.String,
				Name: region_name.String,
			},
		},
		BloodType: &models.BloodType{
			Id:         blood_type_id.String,
			BloodsType: requests_blood_type.String,
		},
		Quantity:  quantity.Float64,
		Status:    status.String,
		CreatedAt: created_at.String,
		UpdatedAt: updated_at.String,
	}, nil
}

func (r *blood_requestRepo) GetList(ctx context.Context, req *models.BloodRequestGetListRequest) (*models.BloodRequestGetListResponse, error) {

	var (
		resp   = &models.BloodRequestGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
		br.id,
		d.id AS doctor_id,
	    d.first_name AS doctor_first_name,
		d.second_name AS doctor_second_name,
		d.specialization AS doctor_specialization,
		d.address AS doctor_address,
		d.email AS doctor_email,
		d.phone_number AS doctor_phone_number,
		h.id AS hospital_id,
		h.name AS hospital_name,
		h.address AS hospital_address,
		h.phone_number AS hospital_phone,
		r.id AS region_id,
		r.name AS region_name,
		bt.id AS blood_type_id,
		bt.bloods_type AS requests_blood_type,
		br.quantity,
		br.status,
        br.created_at,
		br.updated_at,
		hn.id AS head_nurse_id,
		   hn.first_name AS head_nurse_first_name,
		   hn.second_name AS head_nurse_second_name,
		   hn.date_of_birth AS head_nurse_date_of_birth,
		   hn.phone_number AS head_nurse_phone,
		   hn.email AS head_nurse_email
		FROM blood_requests br
		LEFT JOIN hospitals h ON br.hospital_id = h.id
		LEFT JOIN regions r ON h.region_id = r.id
		LEFT JOIN doctors d ON br.doctor_id = d.id
		LEFT JOIN blood_types bt ON br.blood_type_id = bt.id
	   LEFT JOIN head_nurses hn ON h.head_nurse_id = hn.id
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	// if req.Search != "" {
	// 	where += ` AND Request_date ILIKE '%' || '` + req.Search + `' || '%'`
	// }

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id                       sql.NullString
			doctor_id                sql.NullString
			doctor_first_name        sql.NullString
			doctor_second_name       sql.NullString
			doctor_specialization    sql.NullString
			doctor_address           sql.NullString
			doctor_phone_number      sql.NullString
			doctor_email             sql.NullString
			hospital_id              sql.NullString
			hospital_name            sql.NullString
			hospital_address         sql.NullString
			hospital_phone           sql.NullString
			blood_type_id            sql.NullString
			requests_blood_type      sql.NullString
			region_id                sql.NullString
			region_name              sql.NullString
			quantity                 sql.NullFloat64
			status                   sql.NullString
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
			&doctor_id,
			&doctor_first_name,
			&doctor_second_name,
			&doctor_specialization,
			&doctor_address,
			&doctor_phone_number,
			&doctor_email,
			&hospital_id,
			&hospital_name,
			&hospital_address,
			&hospital_phone,
			&region_id,
			&region_name,
			&blood_type_id,
			&requests_blood_type,
			&quantity,
			&status,
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

		resp.BloodRequests = append(resp.BloodRequests, &models.BloodRequest{
			Id: id.String,
			Doctor: &models.Doctor{
				Id:         doctor_id.String,
				FirstName:  doctor_first_name.String,
				SecondName: doctor_second_name.String,
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
							Address:     hospital_address.String,
							PhoneNumber: hospital_phone.String,
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
				Specialization: doctor_specialization.String,
				Address:        doctor_address.String,
				PhoneNumber:    doctor_phone_number.String,
				Email:          doctor_email.String,
			},
			Hospital: &models.Hospital{
				Id:          hospital_id.String,
				Name:        hospital_name.String,
				Address:     hospital_address.String,
				PhoneNumber: hospital_phone.String,
				Region: &models.Region{
					Id:   region_id.String,
					Name: region_name.String,
				},
			},
			BloodType: &models.BloodType{
				Id:         blood_type_id.String,
				BloodsType: requests_blood_type.String,
			},
			Quantity:  quantity.Float64,
			Status:    status.String,
			CreatedAt: created_at.String,
			UpdatedAt: updated_at.String,
		})
	}

	return resp, nil
}

func (r *blood_requestRepo) Update(ctx context.Context, req *models.UpdateBloodRequest) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			blood_Requests
		SET
		id = :id,
		doctor_id = :doctor_id,
		hospital_id = :hospital_id,
		blood_type_id = :blood_type_id,
		quantity = :quantity,
		status = :status,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":            req.Id,
		"doctor_id":     req.DoctorID,
		"hospital_id":   req.HospitalID,
		"blood_type_id": req.BloodTypeID,
		"quantity":      req.Quantity,
		"status":        req.Status,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *blood_requestRepo) Delete(ctx context.Context, req *models.BloodRequestPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM blood_Requests WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}

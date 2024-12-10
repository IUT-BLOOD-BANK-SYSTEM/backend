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

type blood_acceptanceRepo struct {
	db *pgxpool.Pool
}

func NewBloodAcceptanceRepo(db *pgxpool.Pool) *blood_acceptanceRepo {
	return &blood_acceptanceRepo{
		db: db,
	}
}

func (r *blood_acceptanceRepo) Create(ctx context.Context, req *models.CreateBloodAcceptance) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO blood_acceptances(id, patient_id,hospital_id,doctor_id,acceptance_date, status, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.PatientID,
		req.HospitalID,
		req.DoctorID,
		req.AcceptanceDate,
		req.Status,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *blood_acceptanceRepo) GetByID(ctx context.Context, req *models.BloodAcceptancePrimaryKey) (*models.BloodAcceptance, error) {

	var whereField = "id"
	var (
		query string

		id                       sql.NullString
		acceptance_date          sql.NullString
		status                   sql.NullString
		createdAt                sql.NullString
		updatedAt                sql.NullString
		patient_id               sql.NullString
		patient_first_name       sql.NullString
		patient_second_name      sql.NullString
		patient_phone_number     sql.NullString
		patient_address          sql.NullString
		patient_email            sql.NullString
		patient_disease          sql.NullString
		patient_passport_number  sql.NullString
		patient_date_of_birth    sql.NullString
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
		region_id                sql.NullString
		region_name              sql.NullString
		head_nurse_id            sql.NullString
		head_nurse_first_name    sql.NullString
		head_nurse_second_name   sql.NullString
		head_nurse_date_of_birth sql.NullString
		head_nurse_phone         sql.NullString
		head_nurse_email         sql.NullString
		patient_blood_type_id    sql.NullString
		patient_bloods_type      sql.NullString
	)

	query = `
   SELECT
			ba.id,
			ba.acceptance_date,
			ba.status,
			ba.created_at,
			ba.updated_at,
			u.id AS patient_id,
			u.first_name AS patient_first_name,
			u.second_name AS patient_second_name,
			u.date_of_birth AS patient_date_of_birth,
			u.phone_number AS patient_phone_number,
			u.address AS patient_address,
			u.passport_number AS patient_passport_number,
			u.email AS patient_email,
			u.diseases AS patient_disease,
			h.id AS hospital_id,
			h.name AS hospital_name,
			h.address AS hospital_address,
			h.phone_number AS hospital_phone,
			r.id AS region_id,
			r.name AS region_name,
			hn.id AS head_nurse_id,
			hn.first_name AS head_nurse_first_name,
			hn.second_name AS head_nurse_second_name,
			hn.date_of_birth AS head_nurse_date_of_birth,
			hn.phone_number AS head_nurse_phone,
			hn.email AS head_nurse_email,
			d.id AS doctor_id,
			d.first_name AS doctor_first_name,
			d.second_name AS doctor_second_name,
			d.specialization AS doctor_specialization,
			d.address AS doctor_address,
			d.email AS doctor_email,
			d.phone_number AS doctor_phone_number,
			 bt.id AS patient_blood_type_id,
	        bt.bloods_type AS patient_bloods_type
		FROM blood_acceptances ba
		LEFT JOIN users u ON ba.patient_id = u.id
		LEFT JOIN hospitals h ON ba.hospital_id = h.id
		LEFT JOIN regions r ON h.region_id = r.id
		LEFT JOIN head_nurses hn ON h.head_nurse_id = hn.id
		LEFT JOIN doctors d ON ba.doctor_id = d.id
		LEFT JOIN blood_types bt ON u.blood_type_id = bt.id
		WHERE ` + whereField + ` = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&acceptance_date,
		&status,
		&createdAt,
		&updatedAt,
		&patient_id,
		&patient_first_name,
		&patient_second_name,
		&patient_date_of_birth,
		&patient_phone_number,
		&patient_address,
		&patient_passport_number,
		&patient_email,
		&patient_disease,
		&hospital_id,
		&hospital_name,
		&hospital_address,
		&hospital_phone,
		&region_id,
		&region_name,
		&head_nurse_id,
		&head_nurse_first_name,
		&head_nurse_second_name,
		&head_nurse_date_of_birth,
		&head_nurse_phone,
		&head_nurse_email,
		&doctor_id,
		&doctor_first_name,
		&doctor_second_name,
		&doctor_specialization,
		&doctor_address,
		&doctor_email,
		&doctor_phone_number,
		&patient_blood_type_id,
		&patient_bloods_type,
	)

	if err != nil {
		return nil, err
	}

	return &models.BloodAcceptance{
		Id:             id.String,
		AcceptanceDate: acceptance_date.String,
		Status:         status.String,
		CreatedAt:      createdAt.String,
		UpdatedAt:      updatedAt.String,
		Patient: &models.User{
			Id:          patient_id.String,
			FirstName:   patient_first_name.String,
			SecondName:  patient_second_name.String,
			DateOfBirth: patient_date_of_birth.String,
			BloodType: &models.BloodType{
				Id:         patient_blood_type_id.String,
				BloodsType: patient_bloods_type.String,
			},
			PhoneNumber:    patient_phone_number.String,
			Address:        patient_address.String,
			PassportNumber: patient_passport_number.String,
			Email:          patient_email.String,
			Diseases:       patient_disease.String,
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
		Doctor: &models.Doctor{
			Id:             doctor_id.String,
			FirstName:      doctor_first_name.String,
			SecondName:     doctor_second_name.String,
			Specialization: doctor_specialization.String,
			Address:        doctor_address.String,
			Email:          doctor_email.String,
			PhoneNumber:    doctor_phone_number.String,
		},
	}, nil

}

func (r *blood_acceptanceRepo) GetList(ctx context.Context, req *models.BloodAcceptanceGetListRequest) (*models.BloodAcceptanceGetListResponse, error) {

	var (
		resp   = &models.BloodAcceptanceGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
	        ba.id,
			ba.acceptance_date,
			ba.status,
			ba.created_at,
			ba.updated_at,
			u.id AS patient_id,
			u.first_name AS patient_first_name,
			u.second_name AS patient_second_name,
			u.date_of_birth AS patient_date_of_birth,
			u.phone_number AS patient_phone_number,
			u.address AS patient_address,
			u.passport_number AS patient_passport_number,
			u.email AS patient_email,
			u.diseases AS patient_disease,
			h.id AS hospital_id,
			h.name AS hospital_name,
			h.address AS hospital_address,
			h.phone_number AS hospital_phone,
			r.id AS region_id,
			r.name AS region_name,
			hn.id AS head_nurse_id,
			hn.first_name AS head_nurse_first_name,
			hn.second_name AS head_nurse_second_name,
			hn.date_of_birth AS head_nurse_date_of_birth,
			hn.phone_number AS head_nurse_phone,
			hn.email AS head_nurse_email,
			d.id AS doctor_id,
			d.first_name AS doctor_first_name,
			d.second_name AS doctor_second_name,
			d.specialization AS doctor_specialization,
			d.address AS doctor_address,
			d.email AS doctor_email,
			d.phone_number AS doctor_phone_number,
			bt.id AS patient_blood_type_id,
			bt.bloods_type AS patient_bloods_type
		FROM blood_acceptances ba
		LEFT JOIN users u ON ba.patient_id = u.id
		LEFT JOIN hospitals h ON ba.hospital_id = h.id
		LEFT JOIN regions r ON h.region_id = r.id
		LEFT JOIN head_nurses hn ON h.head_nurse_id = hn.id
		LEFT JOIN doctors d ON ba.doctor_id = d.id
	    LEFT JOIN blood_types bt ON u.blood_type_id = bt.id
	 `

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	// if req.Search != "" {
	// 	where += ` AND Acceptance_date ILIKE '%' || '` + req.Search + `' || '%'`
	// }

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id                       sql.NullString
			acceptance_date          sql.NullString
			status                   sql.NullString
			createdAt                sql.NullString
			updatedAt                sql.NullString
			patient_id               sql.NullString
			patient_first_name       sql.NullString
			patient_second_name      sql.NullString
			patient_phone_number     sql.NullString
			patient_address          sql.NullString
			patient_email            sql.NullString
			patient_disease          sql.NullString
			patient_passport_number  sql.NullString
			patient_date_of_birth    sql.NullString
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
			region_id                sql.NullString
			region_name              sql.NullString
			head_nurse_id            sql.NullString
			head_nurse_first_name    sql.NullString
			head_nurse_second_name   sql.NullString
			head_nurse_date_of_birth sql.NullString
			head_nurse_phone         sql.NullString
			head_nurse_email         sql.NullString
			patient_blood_type_id    sql.NullString
			patient_bloods_type      sql.NullString
		)

		err := rows.Scan(
			&id,
			&acceptance_date,
			&status,
			&createdAt,
			&updatedAt,
			&patient_id,
			&patient_first_name,
			&patient_second_name,
			&patient_date_of_birth,
			&patient_phone_number,
			&patient_address,
			&patient_passport_number,
			&patient_email,
			&patient_disease,
			&hospital_id,
			&hospital_name,
			&hospital_address,
			&hospital_phone,
			&region_id,
			&region_name,
			&head_nurse_id,
			&head_nurse_first_name,
			&head_nurse_second_name,
			&head_nurse_date_of_birth,
			&head_nurse_phone,
			&head_nurse_email,
			&doctor_id,
			&doctor_first_name,
			&doctor_second_name,
			&doctor_specialization,
			&doctor_address,
			&doctor_email,
			&doctor_phone_number,
			&patient_blood_type_id,
			&patient_bloods_type,
		)

		if err != nil {
			return nil, err
		}

		resp.BloodAcceptances = append(resp.BloodAcceptances, &models.BloodAcceptance{
			Id:             id.String,
			AcceptanceDate: acceptance_date.String,
			Status:         status.String,
			CreatedAt:      createdAt.String,
			UpdatedAt:      updatedAt.String,
			Patient: &models.User{
				Id:          patient_id.String,
				FirstName:   patient_first_name.String,
				SecondName:  patient_second_name.String,
				DateOfBirth: patient_date_of_birth.String,
				BloodType: &models.BloodType{
					Id:         patient_blood_type_id.String,
					BloodsType: patient_bloods_type.String,
				},
				PhoneNumber:    patient_phone_number.String,
				Address:        patient_address.String,
				PassportNumber: patient_passport_number.String,
				Email:          patient_email.String,
				Diseases:       patient_disease.String,
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
			Doctor: &models.Doctor{
				Id:             doctor_id.String,
				FirstName:      doctor_first_name.String,
				SecondName:     doctor_second_name.String,
				Specialization: doctor_specialization.String,
				Address:        doctor_address.String,
				Email:          doctor_email.String,
				PhoneNumber:    doctor_phone_number.String,
			},
		})
	}

	return resp, nil
}

func (r *blood_acceptanceRepo) Update(ctx context.Context, req *models.UpdateBloodAcceptance) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			blood_acceptances
		SET
		id = :id,
		patient_id = :patient_id,
		hospital_id = :hospital_id,
		doctor_id = :doctor_id,
		acceptance_date = :acceptance_date,
		status = :status,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":              req.Id,
		"patient_id":      req.PatientID,
		"hospital_id":     req.HospitalID,
		"doctor_id":       req.DoctorID,
		"acceptance_date": req.AcceptanceDate,
		"status":          req.Status,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *blood_acceptanceRepo) Delete(ctx context.Context, req *models.BloodAcceptancePrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM blood_acceptances WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}

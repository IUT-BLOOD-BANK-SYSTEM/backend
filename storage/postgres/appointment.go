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

type appointmentRepo struct {
	db *pgxpool.Pool
}

func NewAppointmentRepo(db *pgxpool.Pool) *appointmentRepo {
	return &appointmentRepo{
		db: db,
	}
}

func (r *appointmentRepo) Create(ctx context.Context, req *models.CreateAppointment) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO appointments(id, donor_id,hospital_id,appointment_date, status, updated_at)
		VALUES ($1, $2, $3, $4, $5,NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.DonorID,
		req.HospitalID,
		req.AppointmentDate,
		req.Status,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *appointmentRepo) GetByID(ctx context.Context, req *models.AppointmentPrimaryKey) (*models.Appointment, error) {

	var whereField = "a.id"
	var (
		query string

		id                     sql.NullString
		donor_id               sql.NullString
		user_first_name        sql.NullString
		user_second_name       sql.NullString
		user_phone_number      sql.NullString
		user_address           sql.NullString
		user_email             sql.NullString
		user_disease           sql.NullString
		user_passport_number   sql.NullString
		user_date_of_birth     sql.NullString
		hospital_id            sql.NullString
		hospital_name          sql.NullString
		hospital_address       sql.NullString
		hospital_phone         sql.NullString
		region_id              sql.NullString
		region_name            sql.NullString
		blood_type_id          sql.NullString
		appointment_blood_type sql.NullString
		appointment_date       sql.NullString
		status                 sql.NullString
		created_at             sql.NullString
		updated_at             sql.NullString
	)

	query = `
		SELECT
			a.id,
			u.id AS donor_id,
		u.first_name AS user_first_name,
		u.second_name AS user_second_name,
		u.date_of_birth AS user_date_of_birth,
		u.phone_number AS user_phone_number,
		u.address AS user_address,
		u.passport_number AS user_passport_number,
		u.email AS user_email,
		u.diseases AS user_disease,
		h.id AS hospital_id,
		h.name AS hospital_name,
		h.address AS hospital_address,
		h.phone_number AS hospital_phone,
		r.id AS region_id,
		r.name AS region_name,
		bt.id AS blood_type_id,
		bt.bloods_type AS appointment_blood_type,
			a.appointment_date,
			a.status,
			a.created_at,
			a.updated_at
		FROM appointments AS a
		LEFT JOIN users u ON a.donor_id = u.id
	LEFT JOIN hospitals h ON a.hospital_id = h.id
	LEFT JOIN regions r ON h.region_id = r.id
	LEFT JOIN blood_types bt ON u.blood_type_id = bt.id
		WHERE ` + whereField + ` = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&donor_id,
		&user_first_name,
		&user_second_name,
		&user_date_of_birth,
		&user_phone_number,
		&user_address,
		&user_passport_number,
		&user_email,
		&user_disease,
		&hospital_id,
		&hospital_name,
		&hospital_address,
		&hospital_phone,
		&region_id,
		&region_name,
		&blood_type_id,
		&appointment_blood_type,
		&appointment_date,
		&status,
		&created_at,
		&updated_at,
	)

	if err != nil {
		return nil, err
	}

	return &models.Appointment{
		Id: id.String,
		Donor: &models.User{
			Id:          donor_id.String,
			FirstName:   user_first_name.String,
			SecondName:  user_second_name.String,
			PhoneNumber: user_phone_number.String,
			Address:     user_address.String,
			Diseases:    user_disease.String,
			DateOfBirth: user_date_of_birth.String,
			BloodType: &models.BloodType{
				Id:         blood_type_id.String,
				BloodsType: appointment_blood_type.String,
			},
			Email:          user_email.String,
			PassportNumber: user_passport_number.String,
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

		AppointmentDate: appointment_date.String,
		Status:          status.String,
		CreatedAt:       created_at.String,
		UpdatedAt:       updated_at.String,
	}, nil
}

func (r *appointmentRepo) GetList(ctx context.Context, req *models.AppointmentGetListRequest) (*models.AppointmentGetListResponse, error) {

	var (
		resp   = &models.AppointmentGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			a.id,
			u.id AS donor_id,
		u.first_name AS user_first_name,
		u.second_name AS user_second_name,
		u.date_of_birth AS user_date_of_birth,
		u.phone_number AS user_phone_number,
		u.address AS user_address,
		u.passport_number AS user_passport_number,
		u.email AS user_email,
		u.diseases AS user_disease,
		h.id AS hospital_id,
		h.name AS hospital_name,
		h.address AS hospital_address,
		h.phone_number AS hospital_phone,
		r.id AS region_id,
		r.name AS region_name,
		bt.id AS blood_type_id,
		bt.bloods_type AS appointment_blood_type,
			a.appointment_date,
			a.status,
			a.created_at,
			a.updated_at
		FROM appointments AS a
		LEFT JOIN users u ON a.donor_id = u.id
	LEFT JOIN hospitals h ON a.hospital_id = h.id
	LEFT JOIN regions r ON h.region_id = r.id
	LEFT JOIN blood_types bt ON u.blood_type_id = bt.id
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	// if req.Search != "" {
	// 	where += ` AND donation_date ILIKE '%' || '` + req.Search + `' || '%'`
	// }

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id                     sql.NullString
			donor_id               sql.NullString
			user_first_name        sql.NullString
			user_second_name       sql.NullString
			user_phone_number      sql.NullString
			user_address           sql.NullString
			user_email             sql.NullString
			user_disease           sql.NullString
			user_passport_number   sql.NullString
			user_date_of_birth     sql.NullString
			hospital_id            sql.NullString
			hospital_name          sql.NullString
			hospital_address       sql.NullString
			hospital_phone         sql.NullString
			region_id              sql.NullString
			region_name            sql.NullString
			blood_type_id          sql.NullString
			appointment_blood_type sql.NullString
			appointment_date       sql.NullString
			status                 sql.NullString
			created_at             sql.NullString
			updated_at             sql.NullString
		)

		err := rows.Scan(
			&id,
			&donor_id,
			&user_first_name,
			&user_second_name,
			&user_date_of_birth,
			&user_phone_number,
			&user_address,
			&user_passport_number,
			&user_email,
			&user_disease,
			&hospital_id,
			&hospital_name,
			&hospital_address,
			&hospital_phone,
			&region_id,
			&region_name,
			&blood_type_id,
			&appointment_blood_type,
			&appointment_date,
			&status,
			&created_at,
			&updated_at,
		)

		if err != nil {
			return nil, err
		}

		resp.Appointments = append(resp.Appointments, &models.Appointment{
			Id: id.String,
			Donor: &models.User{
				Id:          donor_id.String,
				FirstName:   user_first_name.String,
				SecondName:  user_second_name.String,
				PhoneNumber: user_phone_number.String,
				Address:     user_address.String,
				Diseases:    user_disease.String,
				DateOfBirth: user_date_of_birth.String,
				BloodType: &models.BloodType{
					Id:         blood_type_id.String,
					BloodsType: appointment_blood_type.String,
				},
				Email:          user_email.String,
				PassportNumber: user_passport_number.String,
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

			AppointmentDate: appointment_date.String,
			Status:          status.String,
			CreatedAt:       created_at.String,
			UpdatedAt:       updated_at.String,
		})
	}

	return resp, nil
}

func (r *appointmentRepo) Update(ctx context.Context, req *models.UpdateAppointment) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			appointments
		SET
		id = :id,
		donor_id = :donor_id,
		hospital_id = :hospital_id,
		appointment_date = :appointment_date,
		status = :status,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":               req.Id,
		"donor_id":         req.DonorID,
		"hospital_id":      req.HospitalID,
		"appointment_date": req.AppointmentDate,
		"status":           req.Status,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *appointmentRepo) Delete(ctx context.Context, req *models.AppointmentPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM appointments WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}

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

type blood_donationRepo struct {
	db *pgxpool.Pool
}

func NewBloodDonationRepo(db *pgxpool.Pool) *blood_donationRepo {
	return &blood_donationRepo{
		db: db,
	}
}

func (r *blood_donationRepo) Create(ctx context.Context, req *models.CreateBloodDonation) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO blood_donations(id, donor_id,hospital_id,blood_type_id, quantity,donation_date, status, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7 ,NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.DonorID,
		req.HospitalID,
		req.BloodTypeID,
		req.Quantity,
		req.DonationDate,
		req.Status,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *blood_donationRepo) GetByID(ctx context.Context, req *models.BloodDonationPrimaryKey) (*models.BloodDonation, error) {

	var whereField = "bd.id"
	var (
		query string

		id                       sql.NullString
		donor_id                 sql.NullString
		user_first_name          sql.NullString
		user_second_name         sql.NullString
		user_phone_number        sql.NullString
		user_address             sql.NullString
		user_email               sql.NullString
		user_disease             sql.NullString
		user_passport_number     sql.NullString
		user_date_of_birth       sql.NullString
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
		blood_type_id            sql.NullString
		donation_blood_type      sql.NullString
		quantity                 sql.NullFloat64
		donation_date            sql.NullString
		status                   sql.NullString
		createdAt                sql.NullString
		updatedAt                sql.NullString
	)

	query = `
		SELECT
         bd.id,
		bd.quantity,
		bd.donation_date,
		bd.status,
		bd.created_at,
		bd.updated_at,
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
		hn.id AS head_nurse_id,
		hn.first_name AS head_nurse_first_name,
		hn.second_name AS head_nurse_second_name,
		hn.date_of_birth AS head_nurse_date_of_birth,
		hn.phone_number AS head_nurse_phone,
		hn.email AS head_nurse_email,
		bt.id AS blood_type_id,
		bt.bloods_type AS donation_blood_type
	FROM blood_donations bd
	LEFT JOIN users u ON bd.donor_id = u.id
	LEFT JOIN hospitals h ON bd.hospital_id = h.id
	LEFT JOIN regions r ON h.region_id = r.id
	LEFT JOIN head_nurses hn ON h.head_nurse_id = hn.id
	LEFT JOIN blood_types bt ON bd.blood_type_id = bt.id
    WHERE ` + whereField + ` = $1
	`

	// bd.id AS blood_donation_id,

	err := r.db.QueryRow(ctx, query, req.Id).Scan(

		&id,
		&quantity,
		&donation_date,
		&status,
		&createdAt,
		&updatedAt,
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
		&head_nurse_id,
		&head_nurse_first_name,
		&head_nurse_second_name,
		&head_nurse_date_of_birth,
		&head_nurse_phone,
		&head_nurse_email,
		&blood_type_id,
		&donation_blood_type,
	)

	if err != nil {
		return nil, err
	}

	return &models.BloodDonation{
		Id:           id.String,
		Quantity:     quantity.Float64,
		DonationDate: donation_date.String,
		Status:       status.String,
		Donor: &models.User{
			Id:             donor_id.String,
			FirstName:      user_first_name.String,
			SecondName:     user_second_name.String,
			PhoneNumber:    user_phone_number.String,
			Address:        user_address.String,
			Diseases:       user_disease.String,
			DateOfBirth:    user_date_of_birth.String,
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
		BloodType: &models.BloodType{
			Id:         blood_type_id.String,
			BloodsType: donation_blood_type.String,
		},
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}, nil

}

func (r *blood_donationRepo) GetList(ctx context.Context, req *models.BloodDonationGetListRequest) (*models.BloodDonationGetListResponse, error) {

	var (
		resp   = &models.BloodDonationGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
	SELECT
		COUNT(*) OVER(),
		bd.id,
		bd.quantity,
		bd.donation_date,
		bd.status,
		bd.created_at,
		bd.updated_at,
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
		hn.id AS head_nurse_id,
		hn.first_name AS head_nurse_first_name,
		hn.second_name AS head_nurse_second_name,
		hn.date_of_birth AS head_nurse_date_of_birth,
		hn.phone_number AS head_nurse_phone,
		hn.email AS head_nurse_email,
		bt.id AS blood_type_id,
		bt.bloods_type AS donation_blood_type
	FROM blood_donations bd
	LEFT JOIN users u ON bd.donor_id = u.id
	LEFT JOIN hospitals h ON bd.hospital_id = h.id
	LEFT JOIN regions r ON h.region_id = r.id
	LEFT JOIN head_nurses hn ON h.head_nurse_id = hn.id
	LEFT JOIN blood_types bt ON bd.blood_type_id = bt.id
	`
	// bd.id AS blood_donation_id,

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
			id                       sql.NullString
			quantity                 sql.NullFloat64
			donation_date            sql.NullString
			status                   sql.NullString
			createdAt                sql.NullString
			updatedAt                sql.NullString
			donor_id                 sql.NullString
			user_first_name          sql.NullString
			user_second_name         sql.NullString
			user_date_of_birth       sql.NullString
			user_phone_number        sql.NullString
			user_address             sql.NullString
			user_passport_number     sql.NullString
			user_email               sql.NullString
			user_disease             sql.NullString
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
			blood_type_id            sql.NullString
			donation_blood_type      sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&quantity,
			&donation_date,
			&status,
			&createdAt,
			&updatedAt,
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
			&head_nurse_id,
			&head_nurse_first_name,
			&head_nurse_second_name,
			&head_nurse_date_of_birth,
			&head_nurse_phone,
			&head_nurse_email,
			&blood_type_id,
			&donation_blood_type,
		)

		if err != nil {
			return nil, err
		}

		resp.BloodDonations = append(resp.BloodDonations, &models.BloodDonation{
			Id: id.String,
			Donor: &models.User{
				Id:          donor_id.String,
				FirstName:   user_first_name.String,
				SecondName:  user_second_name.String,
				PhoneNumber: user_phone_number.String,
				BloodType: &models.BloodType{
					Id:         blood_type_id.String,
					BloodsType: donation_blood_type.String,
				},
				Address:        user_address.String,
				Email:          user_email.String,
				Diseases:       user_disease.String,
				DateOfBirth:    user_date_of_birth.String,
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
			BloodType: &models.BloodType{
				Id:         blood_type_id.String,
				BloodsType: donation_blood_type.String,
			},
			Quantity:     quantity.Float64,
			DonationDate: donation_date.String,
			Status:       status.String,
			CreatedAt:    createdAt.String,
			UpdatedAt:    updatedAt.String,
		})
	}

	return resp, nil
}

func (r *blood_donationRepo) Update(ctx context.Context, req *models.UpdateBloodDonation) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			blood_donations
		SET
		id = :id,
		donor_id = :donor_id,
		hospital_id = :hospital_id,
		blood_type_id = :blood_type_id,
		quantity = :quantity,
		donation_date = :donation_date,
		status = :status,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":            req.Id,
		"donor_id":      req.DonorID,
		"hospital_id":   req.HospitalID,
		"blood_type_id": req.BloodTypeID,
		"quantity":      req.Quantity,
		"donation_date": req.DonationDate,
		"status":        req.Status,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *blood_donationRepo) Delete(ctx context.Context, req *models.BloodDonationPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM blood_donations WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}

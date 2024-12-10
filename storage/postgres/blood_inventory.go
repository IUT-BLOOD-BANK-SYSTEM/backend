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

type blood_inventoryRepo struct {
	db *pgxpool.Pool
}

func NewBloodInventoryRepo(db *pgxpool.Pool) *blood_inventoryRepo {
	return &blood_inventoryRepo{
		db: db,
	}
}

func (r *blood_inventoryRepo) Create(ctx context.Context, req *models.CreateBloodInventory) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO blood_inventories(id,hospital_id,blood_type_id, quantity, updated_at)
		VALUES ($1, $2, $3, $4,NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.HospitalID,
		req.BloodTypeID,
		req.Quantity,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *blood_inventoryRepo) GetByID(ctx context.Context, req *models.BloodInventoryPrimaryKey) (*models.BloodInventory, error) {

	var whereField = "id"
	var (
		query string

		id                       sql.NullString
		hospital_id              sql.NullString
		hospital_name            sql.NullString
		hospital_phone           sql.NullString
		hospital_address         sql.NullString
		region_id                sql.NullString
		region_name              sql.NullString
		head_nurse_id            sql.NullString
		head_nurse_first_name    sql.NullString
		head_nurse_second_name   sql.NullString
		head_nurse_date_of_birth sql.NullString
		head_nurse_phone         sql.NullString
		head_nurse_email         sql.NullString
		blood_type_id            sql.NullString
		blood_type               sql.NullString
		quantity                 sql.NullFloat64
		createdAt                sql.NullString
		updatedAt                sql.NullString
	)

	query = `
		SELECT
			bi.id,
			bi.hospital_id,
			h.name AS hospital_name,
			h.phone_number AS hospital_phone,
			h.address AS hospital_address,
			r.id AS region_id,
		   r.name AS region_name,
		   hn.id AS head_nurse_id,
		   hn.first_name AS head_nurse_first_name,
		   hn.second_name AS head_nurse_second_name,
		   hn.date_of_birth AS head_nurse_date_of_birth,
		   hn.phone_number AS head_nurse_phone,
		   hn.email AS head_nurse_email,
			bi.blood_type_id,
			bt.bloods_type AS blood_type,
			bi.quantity,
			bi.created_at,
			bi.updated_at
		FROM blood_inventories bi
		LEFT JOIN hospitals h ON bi.hospital_id = h.id
		LEFT JOIN blood_types bt ON bi.blood_type_id = bt.id
		LEFT JOIN regions r ON h.region_id = r.id
	   LEFT JOIN head_nurses hn ON h.head_nurse_id = hn.id
		WHERE bi.` + whereField + ` = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&hospital_id,
		&hospital_name,
		&hospital_phone,
		&hospital_address,
		&region_id,
		&region_name,
		&head_nurse_id,
		&head_nurse_first_name,
		&head_nurse_second_name,
		&head_nurse_date_of_birth,
		&head_nurse_phone,
		&head_nurse_email,
		&blood_type_id,
		&blood_type,
		&quantity,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.BloodInventory{
		Id: id.String,
		Hospital: &models.Hospital{
			Id:          hospital_id.String,
			Name:        hospital_name.String,
			PhoneNumber: hospital_phone.String,
			Address:     hospital_address.String,
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
		BloodType: &models.BloodType{
			Id:         blood_type_id.String,
			BloodsType: blood_type.String,
		},
		Quantity:  quantity.Float64,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}, nil
}

func (r *blood_inventoryRepo) GetList(ctx context.Context, req *models.BloodInventoryGetListRequest) (*models.BloodInventoryGetListResponse, error) {

	var (
		resp   = &models.BloodInventoryGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			bi.id,
			bi.hospital_id,
			h.name AS hospital_name,
			h.phone_number AS hospital_phone,
			h.address AS hospital_address,
			r.id AS region_id,
		   r.name AS region_name,
		   hn.id AS head_nurse_id,
		   hn.first_name AS head_nurse_first_name,
		   hn.second_name AS head_nurse_second_name,
		   hn.date_of_birth AS head_nurse_date_of_birth,
		   hn.phone_number AS head_nurse_phone,
		   hn.email AS head_nurse_email,
			bi.blood_type_id,
			bt.bloods_type AS blood_type,
			bi.quantity,
			bi.created_at,
			bi.updated_at
		FROM blood_inventories bi
		LEFT JOIN hospitals h ON bi.hospital_id = h.id
		LEFT JOIN blood_types bt ON bi.blood_type_id = bt.id
		LEFT JOIN regions r ON h.region_id = r.id
	   LEFT JOIN head_nurses hn ON h.head_nurse_id = hn.id
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	// if req.Search != "" {
	// 	where += ` AND Inventory_date ILIKE '%' || '` + req.Search + `' || '%'`
	// }

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id                       sql.NullString
			hospital_id              sql.NullString
			hospital_name            sql.NullString
			hospital_phone           sql.NullString
			hospital_address         sql.NullString
			region_id                sql.NullString
			region_name              sql.NullString
			head_nurse_id            sql.NullString
			head_nurse_first_name    sql.NullString
			head_nurse_second_name   sql.NullString
			head_nurse_date_of_birth sql.NullString
			head_nurse_phone         sql.NullString
			head_nurse_email         sql.NullString
			blood_type_id            sql.NullString
			blood_type               sql.NullString
			quantity                 sql.NullFloat64
			createdAt                sql.NullString
			updatedAt                sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&hospital_id,
			&hospital_name,
			&hospital_phone,
			&hospital_address,
			&region_id,
			&region_name,
			&head_nurse_id,
			&head_nurse_first_name,
			&head_nurse_second_name,
			&head_nurse_date_of_birth,
			&head_nurse_phone,
			&head_nurse_email,
			&blood_type_id,
			&blood_type,
			&quantity,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.BloodInventories = append(resp.BloodInventories, &models.BloodInventory{
			Id: id.String,
			Hospital: &models.Hospital{
				Id:          hospital_id.String,
				Name:        hospital_name.String,
				PhoneNumber: hospital_phone.String,
				Address:     hospital_address.String,
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
			BloodType: &models.BloodType{
				Id:         blood_type_id.String,
				BloodsType: blood_type.String,
			},
			Quantity:  quantity.Float64,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return resp, nil
}

func (r *blood_inventoryRepo) Update(ctx context.Context, req *models.UpdateBloodInventory) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			blood_inventories
		SET
		id = :id,
		hospital_id = :hospital_id,
		blood_type_id = :blood_type_id,
		quantity = :quantity,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":            req.Id,
		"hospital_id":   req.HospitalID,
		"blood_type_id": req.BloodTypeID,
		"quantity":      req.Quantity,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *blood_inventoryRepo) Delete(ctx context.Context, req *models.BloodInventoryPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM blood_inventories WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}

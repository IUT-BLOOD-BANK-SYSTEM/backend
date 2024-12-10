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

type blood_typeRepo struct {
	db *pgxpool.Pool
}

func NewBloodTypeRepo(db *pgxpool.Pool) *blood_typeRepo {
	return &blood_typeRepo{
		db: db,
	}
}

func (r *blood_typeRepo) Create(ctx context.Context, req *models.CreateBloodType) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO blood_types(id, bloods_type, updated_at)
		VALUES ($1, $2, NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.BloodsType,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *blood_typeRepo) GetByID(ctx context.Context, req *models.BloodTypePrimaryKey) (*models.BloodType, error) {

	var whereField = "id"
	var (
		query string

		id          sql.NullString
		bloods_type sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)

	query = `
		SELECT
			id,
			bloods_type,
			created_at,
			updated_at
		FROM blood_types
		WHERE ` + whereField + ` = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&bloods_type,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.BloodType{
		Id:         id.String,
		BloodsType: bloods_type.String,
		CreatedAt:  createdAt.String,
		UpdatedAt:  updatedAt.String,
	}, nil
}

func (r *blood_typeRepo) GetList(ctx context.Context, req *models.BloodTypeGetListRequest) (*models.BloodTypeGetListResponse, error) {

	var (
		resp   = &models.BloodTypeGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			bloods_type,
			created_at,
			updated_at
		FROM blood_types
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND bloods_type ILIKE '%' || '` + req.Search + `' || '%'`
	}

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id          sql.NullString
			bloods_type sql.NullString
			createdAt   sql.NullString
			updatedAt   sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&bloods_type,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.BloodTypes = append(resp.BloodTypes, &models.BloodType{
			Id:         id.String,
			BloodsType: bloods_type.String,
			CreatedAt:  createdAt.String,
			UpdatedAt:  updatedAt.String,
		})
	}

	return resp, nil
}

func (r *blood_typeRepo) Update(ctx context.Context, req *models.UpdateBloodType) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			blood_types
		SET
		id = :id,
		bloods_type = :bloods_type,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":          req.Id,
		"bloods_type": req.BloodsType,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *blood_typeRepo) Delete(ctx context.Context, req *models.BloodTypePrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM blood_types WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}

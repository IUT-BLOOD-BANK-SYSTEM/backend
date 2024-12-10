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

type sis_adminRepo struct {
	db *pgxpool.Pool
}

func NewSisAdminRepo(db *pgxpool.Pool) *sis_adminRepo {
	return &sis_adminRepo{
		db: db,
	}
}

func (r *sis_adminRepo) Create(ctx context.Context, req *models.CreateSisAdmin) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO sis_admins(id, email,password, updated_at)
		VALUES ($1, $2, $3,NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Email,
		req.Password,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *sis_adminRepo) GetByID(ctx context.Context, req *models.SisAdminPrimaryKey) (*models.SisAdmin, error) {

	var whereField = "id"
	if len(req.Email) > 0 {
		whereField = "email"
		req.Id = req.Email
	}
	var (
		query string

		id        sql.NullString
		email     sql.NullString
		password  sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	query = `
		SELECT
			id,
			email,
			password,
			created_at,
			updated_at
		FROM sis_admins
		WHERE ` + whereField + ` = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&email,
		&password,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.SisAdmin{
		Id:        id.String,
		Email:     email.String,
		Password:  password.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}, nil
}

func (r *sis_adminRepo) GetList(ctx context.Context, req *models.SisAdminGetListRequest) (*models.SisAdminGetListResponse, error) {

	var (
		resp   = &models.SisAdminGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			email,
			password,
			created_at,
			updated_at
		FROM sis_admins
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND email ILIKE '%' || '` + req.Search + `' || '%'`
	}

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id        sql.NullString
			email     sql.NullString
			password  sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&email,
			&password,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.SisAdmins = append(resp.SisAdmins, &models.SisAdmin{
			Id:        id.String,
			Email:     email.String,
			Password:  password.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return resp, nil
}

func (r *sis_adminRepo) Update(ctx context.Context, req *models.UpdateSisAdmin) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			sis_admins
		SET
		id = :id,
		email = :email,
		password = :password,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":       req.Id,
		"email":    req.Email,
		"password": req.Password,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *sis_adminRepo) Delete(ctx context.Context, req *models.SisAdminPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM sis_admins WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}

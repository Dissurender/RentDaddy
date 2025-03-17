// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: leases.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createLease = `-- name: CreateLease :one
INSERT INTO leases (
    lease_version, lease_file_key, lease_template_id, tenant_id, landlord_id, apartment_id, 
    lease_start_date, lease_end_date, rent_amount, lease_status,
    created_by, updated_by
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING id
`

type CreateLeaseParams struct {
	LeaseVersion    int64          `json:"lease_version"`
	LeaseFileKey    pgtype.Text    `json:"lease_file_key"`
	LeaseTemplateID pgtype.Int4    `json:"lease_template_id"`
	TenantID        int64          `json:"tenant_id"`
	LandlordID      int64          `json:"landlord_id"`
	ApartmentID     pgtype.Int8    `json:"apartment_id"`
	LeaseStartDate  pgtype.Date    `json:"lease_start_date"`
	LeaseEndDate    pgtype.Date    `json:"lease_end_date"`
	RentAmount      pgtype.Numeric `json:"rent_amount"`
	LeaseStatus     string         `json:"lease_status"`
	CreatedBy       int64          `json:"created_by"`
	UpdatedBy       int64          `json:"updated_by"`
}

func (q *Queries) CreateLease(ctx context.Context, arg CreateLeaseParams) (int32, error) {
	row := q.db.QueryRow(ctx, createLease,
		arg.LeaseVersion,
		arg.LeaseFileKey,
		arg.LeaseTemplateID,
		arg.TenantID,
		arg.LandlordID,
		arg.ApartmentID,
		arg.LeaseStartDate,
		arg.LeaseEndDate,
		arg.RentAmount,
		arg.LeaseStatus,
		arg.CreatedBy,
		arg.UpdatedBy,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const createLeaseTemplate = `-- name: CreateLeaseTemplate :one
INSERT INTO lease_templates (template_name, s3_key, created_by)
VALUES ($1, $2, $3)
RETURNING id
`

type CreateLeaseTemplateParams struct {
	TemplateName string `json:"template_name"`
	S3Key        string `json:"s3_key"`
	CreatedBy    int32  `json:"created_by"`
}

func (q *Queries) CreateLeaseTemplate(ctx context.Context, arg CreateLeaseTemplateParams) (int32, error) {
	row := q.db.QueryRow(ctx, createLeaseTemplate, arg.TemplateName, arg.S3Key, arg.CreatedBy)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getLatestLeaseByApartment = `-- name: GetLatestLeaseByApartment :one
SELECT id, lease_version, lease_file_key, lease_template_id, tenant_id, landlord_id, apartment_id, lease_start_date, lease_end_date, rent_amount, lease_status, created_by, updated_by, created_at, updated_at FROM leases WHERE apartment_id = $1 ORDER BY lease_version DESC LIMIT 1
`

func (q *Queries) GetLatestLeaseByApartment(ctx context.Context, apartmentID pgtype.Int8) (Lease, error) {
	row := q.db.QueryRow(ctx, getLatestLeaseByApartment, apartmentID)
	var i Lease
	err := row.Scan(
		&i.ID,
		&i.LeaseVersion,
		&i.LeaseFileKey,
		&i.LeaseTemplateID,
		&i.TenantID,
		&i.LandlordID,
		&i.ApartmentID,
		&i.LeaseStartDate,
		&i.LeaseEndDate,
		&i.RentAmount,
		&i.LeaseStatus,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getLatestLeaseByTenant = `-- name: GetLatestLeaseByTenant :one
SELECT id, lease_version, lease_file_key, lease_template_id, tenant_id, landlord_id, apartment_id, lease_start_date, lease_end_date, rent_amount, lease_status, created_by, updated_by, created_at, updated_at FROM leases WHERE tenant_id = $1 ORDER BY lease_version DESC LIMIT 1
`

func (q *Queries) GetLatestLeaseByTenant(ctx context.Context, tenantID int64) (Lease, error) {
	row := q.db.QueryRow(ctx, getLatestLeaseByTenant, tenantID)
	var i Lease
	err := row.Scan(
		&i.ID,
		&i.LeaseVersion,
		&i.LeaseFileKey,
		&i.LeaseTemplateID,
		&i.TenantID,
		&i.LandlordID,
		&i.ApartmentID,
		&i.LeaseStartDate,
		&i.LeaseEndDate,
		&i.RentAmount,
		&i.LeaseStatus,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getLeaseByID = `-- name: GetLeaseByID :one
SELECT id, lease_version, lease_file_key, lease_template_id, tenant_id, landlord_id, apartment_id, lease_start_date, lease_end_date, rent_amount, lease_status, created_by, updated_by, created_at, updated_at FROM leases WHERE id = $1 LIMIT 1
`

func (q *Queries) GetLeaseByID(ctx context.Context, id int32) (Lease, error) {
	row := q.db.QueryRow(ctx, getLeaseByID, id)
	var i Lease
	err := row.Scan(
		&i.ID,
		&i.LeaseVersion,
		&i.LeaseFileKey,
		&i.LeaseTemplateID,
		&i.TenantID,
		&i.LandlordID,
		&i.ApartmentID,
		&i.LeaseStartDate,
		&i.LeaseEndDate,
		&i.RentAmount,
		&i.LeaseStatus,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getLeaseWithTemplate = `-- name: GetLeaseWithTemplate :one
SELECT leases.id, leases.lease_version, leases.lease_file_key, leases.lease_template_id, leases.tenant_id, leases.landlord_id, leases.apartment_id, leases.lease_start_date, leases.lease_end_date, leases.rent_amount, leases.lease_status, leases.created_by, leases.updated_by, leases.created_at, leases.updated_at, lease_templates.s3_key AS template_s3_key
FROM leases
JOIN lease_templates ON leases.lease_template_id = lease_templates.id
WHERE leases.id = $1
`

type GetLeaseWithTemplateRow struct {
	ID              int32            `json:"id"`
	LeaseVersion    int64            `json:"lease_version"`
	LeaseFileKey    pgtype.Text      `json:"lease_file_key"`
	LeaseTemplateID pgtype.Int4      `json:"lease_template_id"`
	TenantID        int64            `json:"tenant_id"`
	LandlordID      int64            `json:"landlord_id"`
	ApartmentID     pgtype.Int8      `json:"apartment_id"`
	LeaseStartDate  pgtype.Date      `json:"lease_start_date"`
	LeaseEndDate    pgtype.Date      `json:"lease_end_date"`
	RentAmount      pgtype.Numeric   `json:"rent_amount"`
	LeaseStatus     string           `json:"lease_status"`
	CreatedBy       int64            `json:"created_by"`
	UpdatedBy       int64            `json:"updated_by"`
	CreatedAt       pgtype.Timestamp `json:"created_at"`
	UpdatedAt       pgtype.Timestamp `json:"updated_at"`
	TemplateS3Key   string           `json:"template_s3_key"`
}

func (q *Queries) GetLeaseWithTemplate(ctx context.Context, id int32) (GetLeaseWithTemplateRow, error) {
	row := q.db.QueryRow(ctx, getLeaseWithTemplate, id)
	var i GetLeaseWithTemplateRow
	err := row.Scan(
		&i.ID,
		&i.LeaseVersion,
		&i.LeaseFileKey,
		&i.LeaseTemplateID,
		&i.TenantID,
		&i.LandlordID,
		&i.ApartmentID,
		&i.LeaseStartDate,
		&i.LeaseEndDate,
		&i.RentAmount,
		&i.LeaseStatus,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.TemplateS3Key,
	)
	return i, err
}

const listLeases = `-- name: ListLeases :many
SELECT id, lease_version, lease_file_key, lease_template_id, tenant_id, landlord_id, apartment_id, lease_start_date, lease_end_date, rent_amount, lease_status, created_by, updated_by, created_at, updated_at FROM leases ORDER BY created_at DESC
`

func (q *Queries) ListLeases(ctx context.Context) ([]Lease, error) {
	rows, err := q.db.Query(ctx, listLeases)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Lease
	for rows.Next() {
		var i Lease
		if err := rows.Scan(
			&i.ID,
			&i.LeaseVersion,
			&i.LeaseFileKey,
			&i.LeaseTemplateID,
			&i.TenantID,
			&i.LandlordID,
			&i.ApartmentID,
			&i.LeaseStartDate,
			&i.LeaseEndDate,
			&i.RentAmount,
			&i.LeaseStatus,
			&i.CreatedBy,
			&i.UpdatedBy,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const renewLease = `-- name: RenewLease :exec
UPDATE leases
SET 
    lease_version = lease_version + 1,  
    lease_end_date = $1, 
    rent_amount = $2, 
    lease_status = 'renewed', 
    updated_by = $3, 
    updated_at = now()
WHERE id = $4
RETURNING id
`

type RenewLeaseParams struct {
	LeaseEndDate pgtype.Date    `json:"lease_end_date"`
	RentAmount   pgtype.Numeric `json:"rent_amount"`
	UpdatedBy    int64          `json:"updated_by"`
	ID           int32          `json:"id"`
}

func (q *Queries) RenewLease(ctx context.Context, arg RenewLeaseParams) error {
	_, err := q.db.Exec(ctx, renewLease,
		arg.LeaseEndDate,
		arg.RentAmount,
		arg.UpdatedBy,
		arg.ID,
	)
	return err
}

const terminateLease = `-- name: TerminateLease :exec
UPDATE leases
SET 
    lease_status = 'terminated', 
    updated_by = $1, 
    updated_at = now()
WHERE id = $2
RETURNING id
`

type TerminateLeaseParams struct {
	UpdatedBy int64 `json:"updated_by"`
	ID        int32 `json:"id"`
}

func (q *Queries) TerminateLease(ctx context.Context, arg TerminateLeaseParams) error {
	_, err := q.db.Exec(ctx, terminateLease, arg.UpdatedBy, arg.ID)
	return err
}

const updateLease = `-- name: UpdateLease :exec
UPDATE leases
SET 
    lease_start_date = $1,
    lease_end_date = $2,
    rent_amount = $3,
    lease_status = $4,
    updated_by = $5,
    updated_at = now()
WHERE id = $6
RETURNING id
`

type UpdateLeaseParams struct {
	LeaseStartDate pgtype.Date    `json:"lease_start_date"`
	LeaseEndDate   pgtype.Date    `json:"lease_end_date"`
	RentAmount     pgtype.Numeric `json:"rent_amount"`
	LeaseStatus    string         `json:"lease_status"`
	UpdatedBy      int64          `json:"updated_by"`
	ID             int32          `json:"id"`
}

func (q *Queries) UpdateLease(ctx context.Context, arg UpdateLeaseParams) error {
	_, err := q.db.Exec(ctx, updateLease,
		arg.LeaseStartDate,
		arg.LeaseEndDate,
		arg.RentAmount,
		arg.LeaseStatus,
		arg.UpdatedBy,
		arg.ID,
	)
	return err
}

const updateLeaseFileKey = `-- name: UpdateLeaseFileKey :exec
UPDATE leases
SET lease_file_key = $1, updated_by = $2, updated_at = now()
WHERE id = $3
`

type UpdateLeaseFileKeyParams struct {
	LeaseFileKey pgtype.Text `json:"lease_file_key"`
	UpdatedBy    int64       `json:"updated_by"`
	ID           int32       `json:"id"`
}

func (q *Queries) UpdateLeaseFileKey(ctx context.Context, arg UpdateLeaseFileKeyParams) error {
	_, err := q.db.Exec(ctx, updateLeaseFileKey, arg.LeaseFileKey, arg.UpdatedBy, arg.ID)
	return err
}

package postgres

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"github.com/odahu/odahu-flow/packages/operator/pkg/apis/training"
	odahuErrors "github.com/odahu/odahu-flow/packages/operator/pkg/errors"
	"github.com/odahu/odahu-flow/packages/operator/pkg/utils/filter"
)

const (
	toolchainIntegrationTable   = "odahu_operator_toolchain_integration"
	uniqueViolationPostgresCode = pq.ErrorCode("23505") // unique_violation
)

var (
	MaxSize   = 500
	FirstPage = 0
)

type ToolchainRepo struct {
	DB *sql.DB
}

func (tr ToolchainRepo) GetToolchainIntegration(name string) (*training.ToolchainIntegration, error) {

	ti := new(training.ToolchainIntegration)

	err := tr.DB.QueryRow(
		fmt.Sprintf("SELECT id, spec, status, created, updated FROM %s WHERE id = $1", toolchainIntegrationTable),
		name,
	).Scan(&ti.ID, &ti.Spec, &ti.Status, &ti.CreatedAt, &ti.UpdatedAt)

	switch {
	case err == sql.ErrNoRows:
		return nil, odahuErrors.NotFoundError{Entity: name}
	case err != nil:
		return nil, err
	default:
		return ti, nil
	}

}

func (tr ToolchainRepo) GetToolchainIntegrationList(options ...filter.ListOption) (
	[]training.ToolchainIntegration, error,
) {

	listOptions := &filter.ListOptions{
		Filter: nil,
		Page:   &FirstPage,
		Size:   &MaxSize,
	}
	for _, option := range options {
		option(listOptions)
	}

	offset := *listOptions.Size * (*listOptions.Page)

	stmt := "SELECT id, spec, status, created, updated " +
		"FROM odahu_operator_toolchain_integration ORDER BY id LIMIT $1 OFFSET $2"

	rows, err := tr.DB.Query(stmt, *listOptions.Size, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tis []training.ToolchainIntegration

	for rows.Next() {
		ti := new(training.ToolchainIntegration)
		err := rows.Scan(&ti.ID, &ti.Spec, &ti.Status, &ti.CreatedAt, &ti.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tis = append(tis, *ti)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return tis, nil

}

func (tr ToolchainRepo) DeleteToolchainIntegration(name string) error {

	// First try to check that row exists otherwise raise exception to fit interface
	_, err := tr.GetToolchainIntegration(name)
	if err != nil {
		return err
	}

	// If exists, delete it

	sqlStatement := fmt.Sprintf("DELETE FROM %s WHERE id = $1", toolchainIntegrationTable)
	_, err = tr.DB.Exec(sqlStatement, name)
	if err != nil {
		return err
	}
	return nil
}

func (tr ToolchainRepo) UpdateToolchainIntegration(md *training.ToolchainIntegration) error {

	// First try to check that row exists otherwise raise exception to fit interface
	oldTi, err := tr.GetToolchainIntegration(md.ID)
	if err != nil {
		return err
	}

	md.Status = oldTi.Status

	sqlStatement := fmt.Sprintf("UPDATE %s SET spec = $1, status = $2, updated = $3 WHERE id = $4",
		toolchainIntegrationTable)
	_, err = tr.DB.Exec(sqlStatement, md.Spec, md.Status, md.UpdatedAt, md.ID)
	if err != nil {
		return err
	}
	return nil
}

func (tr ToolchainRepo) SaveToolchainIntegration(md *training.ToolchainIntegration) error {

	_, err := tr.DB.Exec(
		fmt.Sprintf("INSERT INTO %s (id, spec, status, created, updated) VALUES($1, $2, $3, $4, $5)",
			toolchainIntegrationTable),
		md.ID, md.Spec, md.Status, md.CreatedAt, md.UpdatedAt,
	)
	if err != nil {
		pqError, ok := err.(*pq.Error)
		if ok && pqError.Code == uniqueViolationPostgresCode {
			return odahuErrors.AlreadyExistError{Entity: md.ID}
		}
		return err
	}
	return nil

}

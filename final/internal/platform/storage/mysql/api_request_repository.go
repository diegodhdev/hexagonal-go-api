package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	mooc "github.com/diegodhdev/hexagonal-go-api/final/internal"
	"github.com/huandu/go-sqlbuilder"
)

// ApiRequestRepository is a MySQL mooc.ApiRequestRepository implementation.
type ApiRequestRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

// NewApiRequestRepository initializes a MySQL-based implementation of mooc.NewApiRequestRepository.
func NewApiRequestRepository(db *sql.DB, dbTimeout time.Duration) *ApiRequestRepository {
	return &ApiRequestRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

// Save implements the mooc.CourseRepository interface.
func (r *ApiRequestRepository) Save(ctx context.Context, course mooc.ApiRequest) error {
	apiRequestSQLStruct := sqlbuilder.NewStruct(new(sqlApiRequest))
	query, args := apiRequestSQLStruct.InsertInto(sqlApiRequestTable, sqlApiRequest{
		ID: course.ID().String(),
	}).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist api_request on database: %v", err)
	}

	return nil
}

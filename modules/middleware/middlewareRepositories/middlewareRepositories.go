package middlewareRepositories

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/nattrio/go-ecommerce/modules/middleware"
)

type IMiddlewareRepository interface {
	FindAccessToken(userId, AccessToken string) bool
	FindRole() ([]*middleware.Role, error)
}

type middlewareRepository struct {
	db *sqlx.DB
}

func MiddlewareRepository(db *sqlx.DB) IMiddlewareRepository {
	return &middlewareRepository{
		db: db,
	}
}

func (r *middlewareRepository) FindAccessToken(userId, AccessToken string) bool {
	query := `
	SELECT
		(CASE WHEN COUNT(*) = 1 THEN true ELSE false END)
	FROM oauth
	WHERE user_id = $1
	AND access_token = $2;
	`

	var check bool
	if err := r.db.Get(&check, query, userId, AccessToken); err != nil {
		return false
	}
	return true
}

func (r *middlewareRepository) FindRole() ([]*middleware.Role, error) {
	query := `
	SELECT
		"id",
		"title"
	FROM "roles"
	ORDER BY "id" DESC;`

	roles := make([]*middleware.Role, 0)
	if err := r.db.Select(&roles, query); err != nil {
		return nil, fmt.Errorf("roles are empty")
	}
	return roles, nil
}

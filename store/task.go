package store

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/keisas/golang.git/entity"
)

func (r *Repository) ListTasks(
	ctx context.Context, db *sqlx.DB,
) (entity.Tasks, error) {
	sql := `SELECT 
				id, 
				title, 
				status, 
				created, 
				modified
			FROM task;`
	rows, err := db.QueryContext(ctx, sql)
	if err := db.SelectContext(ctx, &tasks, sql); err != nil {
		return nil, err
	}
	return tasks, nil
}

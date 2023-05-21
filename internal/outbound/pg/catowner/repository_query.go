package catowner

import (
	"context"
	"fmt"

	"catownerclub/internal/domain/catowner"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repositoryQuery struct {
	db *pgxpool.Pool

	catOwnerID int

	getCatsFunc func(context.Context) (map[int]*catowner.Cat, error)
}

func newRepositoryQuery(db *pgxpool.Pool, catOwnerID int) *repositoryQuery {
	return &repositoryQuery{
		db:         db,
		catOwnerID: catOwnerID,
		getCatsFunc: func(context.Context) (map[int]*catowner.Cat, error) {
			return make(map[int]*catowner.Cat), nil
		},
	}
}

func (q *repositoryQuery) WithCats() catowner.RepositoryQuery {
	sql := `
		select
			id,
			name
		from cats
		where owner_id = $1
	`

	cats := make(map[int]*catowner.Cat)

	q.getCatsFunc = func(ctx context.Context) (map[int]*catowner.Cat, error) {
		rows, err := q.db.Query(ctx, sql, q.catOwnerID)
		if err != nil {
			return nil, fmt.Errorf("failed to get all cats: %w", err)
		}

		defer rows.Close()

		for rows.Next() {
			var cat catowner.Cat

			err := rows.Scan(
				&cat.ID,
				&cat.Name,
			)

			if err != nil {
				return nil, fmt.Errorf("failed to scan a cat: %w", err)
			}

			cats[cat.ID] = &cat
		}

		return cats, nil
	}

	return q
}

func (q *repositoryQuery) WithCat(catID int) catowner.RepositoryQuery {
	sql := `
		select
			id,
			name
		from cats
		where id = $1 and owner_id = $2
	`

	q.getCatsFunc = func(ctx context.Context) (map[int]*catowner.Cat, error) {
		row := q.db.QueryRow(ctx, sql, catID, q.catOwnerID)

		var cat catowner.Cat

		err := row.Scan(
			&cat.ID,
			&cat.Name,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan a cat: %w", err)
		}

		return map[int]*catowner.Cat{cat.ID: &cat}, nil
	}

	return q
}

func (q *repositoryQuery) Load(ctx context.Context) (*catowner.CatOwner, error) {
	sql := `
		select
			id,
			name
		from cat_owners
		where id = $1
	`

	row := q.db.QueryRow(ctx, sql, q.catOwnerID)

	var (
		catOwnerID   int
		catOwnerName string
	)

	err := row.Scan(
		&catOwnerID,
		&catOwnerName,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to scan a cat owner: %w", err)
	}

	cats, err := q.getCatsFunc(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get cats of the cat owner: %w", err)
	}

	catOwner := catowner.New(catOwnerID, catOwnerName)

	catOwner.Cats = cats

	return catOwner, nil
}

package catowner

import (
	"context"
	"fmt"

	"catownerclub/internal/domain/catowner"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Get(id int) catowner.RepositoryQuery {
	return newRepositoryQuery(r.db, id)
}

func (r *Repository) Add(ctx context.Context, catOwner *catowner.CatOwner) error {
	batch := pgx.Batch{}

	batch.Queue(`
		insert into cat_owners (id, name)
		values ($1, $2)
	`,
		catOwner.ID,
		catOwner.Name,
	)

	newCatSynchronizer(&batch).sync(catOwner)

	if err := r.db.SendBatch(ctx, &batch).Close(); err != nil {
		return fmt.Errorf("failed to add a cat owner in the repository: %w", err)
	}

	return nil
}

func (r *Repository) Update(ctx context.Context, catOwner *catowner.CatOwner) error {
	batch := pgx.Batch{}

	if catOwner.Updated {
		batch.Queue(`
			update cat_owners set
				name = $1
			where id = $2
		`,
			catOwner.Name,
			catOwner.ID)
	}

	newCatSynchronizer(&batch).sync(catOwner)

	if err := r.db.SendBatch(ctx, &batch).Close(); err != nil {
		return fmt.Errorf("failed to update a cat owner in the repository: %w", err)
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, catOwner *catowner.CatOwner) error {
	batch := pgx.Batch{}

	newCatSynchronizer(&batch).sync(catOwner)

	batch.Queue(`
		delete from cat_owners
		where id = $1 
	`, catOwner.ID)

	if err := r.db.SendBatch(ctx, &batch).Close(); err != nil {
		return fmt.Errorf("failed to delete a cat owner from the repository: %w", err)
	}

	return nil
}

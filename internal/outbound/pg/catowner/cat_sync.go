package catowner

import (
	"catownerclub/internal/domain/catowner"

	"github.com/jackc/pgx/v5"
)

type catSynchronizer struct {
	batch *pgx.Batch
}

func newCatSynchronizer(batch *pgx.Batch) *catSynchronizer {
	return &catSynchronizer{
		batch: batch,
	}
}

func (s *catSynchronizer) sync(catOwner *catowner.CatOwner) {
	s.insertCats(catOwner.ID, catOwner.InsertedCats)
	s.updateCats(catOwner.UpdatedCats)
	s.deleteCats(catOwner.DeletedCats)
}

func (s *catSynchronizer) insertCats(catOwnerID int, cats map[int]*catowner.Cat) {
	for _, cat := range cats {
		s.batch.Queue(`
			insert into cats (id, owner_id, name)
			values ($1, $2, $3)
		`, cat.ID, catOwnerID, cat.Name)
	}
}

func (s *catSynchronizer) updateCats(cats map[int]*catowner.Cat) {
	for _, cat := range cats {
		s.batch.Queue(`
			update cats set
				name = $1
			where id = $2
		`, cat.Name, cat.ID)
	}
}

func (s *catSynchronizer) deleteCats(cats map[int]*catowner.Cat) {
	for _, cat := range cats {
		s.batch.Queue(`
			delete from cats
			where id = $1
		`, cat.ID)
	}
}

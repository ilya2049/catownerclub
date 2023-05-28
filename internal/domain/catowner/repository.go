package catowner

import "context"

type Repository interface {
	Get(id int) RepositoryQuery

	Add(ctx context.Context, catOwner *CatOwner) error
	Update(ctx context.Context, catOwner *CatOwner) error
	Delete(ctx context.Context, catOwner *CatOwner) error
}

type RepositoryQuery interface {
	ForUpdate() RepositoryQuery

	WithCats() RepositoryQuery
	WithCat(catID int) RepositoryQuery
	Load(context.Context) (*CatOwner, error)
}

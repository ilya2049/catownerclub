package catowner

type CatOwner struct {
	InsertedCats map[int]*Cat
	UpdatedCats  map[int]*Cat
	DeletedCats  map[int]*Cat
	Updated      bool

	ID   int
	Name string
	Cats map[int]*Cat
}

func New(id int, name string) *CatOwner {
	return &CatOwner{
		InsertedCats: map[int]*Cat{},
		UpdatedCats:  map[int]*Cat{},
		DeletedCats:  map[int]*Cat{},
		Updated:      false,
		ID:           id,
		Name:         name,
		Cats:         map[int]*Cat{},
	}
}

func (owner *CatOwner) TakePossessionOfCat(cat *Cat) {
	owner.InsertedCats[cat.ID] = cat

	owner.Cats[cat.ID] = cat
}

func (owner *CatOwner) GiveUpCat(catID int) {
	if cat, ok := owner.Cats[catID]; ok {
		owner.DeletedCats[cat.ID] = cat

		delete(owner.Cats, cat.ID)
	}
}

func (owner *CatOwner) GiveUpAllCats() {
	for _, cat := range owner.Cats {
		owner.DeletedCats[cat.ID] = cat

		delete(owner.Cats, cat.ID)
	}
}

func (owner *CatOwner) RenameCat(catID int, newCatName string) {
	if cat, ok := owner.Cats[catID]; ok {
		owner.UpdatedCats[catID] = cat

		cat.Rename(newCatName)
	}
}

func (owner *CatOwner) Rename(newName string) {
	owner.Name = newName

	owner.Updated = true
}

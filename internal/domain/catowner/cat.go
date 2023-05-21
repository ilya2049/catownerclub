package catowner

type Cat struct {
	ID   int
	Name string
}

func NewCat(
	id int,
	name string,
) *Cat {
	return &Cat{
		ID:   id,
		Name: name,
	}
}

func (c *Cat) Rename(newName string) {
	c.Name = newName
}

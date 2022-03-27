package domain

type Pet struct {
	Id   string
	Name string
	Tag  string
}

func NewPet(id string, name string, tag string) (*Pet, error) {
	return &Pet{
		Id:   id,
		Name: name,
		Tag:  tag,
	}, nil
}

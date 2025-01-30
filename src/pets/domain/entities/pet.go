package entities

type Pet struct {
	ID   int
	Name string
	Raza string
}

var increment = 0

func NewPet(name, raza string) *Pet {
	increment++
	pet := Pet{ID: increment, Name: name, Raza: raza}
	return &pet
}

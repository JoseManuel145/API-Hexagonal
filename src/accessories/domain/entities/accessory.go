package entities

type Accessory struct {
	Id          int
	Name        string
	Description string
}

var increment = 0

func NewAccessory(name, description string) *Accessory {
	increment++
	accessory := Accessory{Id: increment, Name: name, Description: description}
	return &accessory
}

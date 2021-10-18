package autotransport

import "fmt"

var AllEntities = []Water{
	{Id: 1, Name: "One", Model: "X1", Manufacturer: "BMW", Material: "Iron", Speed: 10},
	{Id: 2, Name: "Two", Model: "X2", Manufacturer: "Toyota", Material: "Iron", Speed: 20},
	{Id: 3, Name: "Three", Model: "X3", Manufacturer: "Tesla", Material: "Wood", Speed: 30},
	{Id: 4, Name: "Four", Model: "X4", Manufacturer: "Ford", Material: "Wood", Speed: 40},
	{Id: 5, Name: "Five", Model: "X5", Manufacturer: "Kia", Material: "Plastic", Speed: 50},
}

type Water struct {
	Id uint64
	Name string
	Model string
	Manufacturer string
	Material string
	Speed uint
}

func NewWater(id uint64, name string, model string, manufacturer string, material string, speed uint) *Water {
	return &Water{
		id,
		name,
		model,
		manufacturer,
		material,
		speed,
	}
}

func (a Water) String() string {
	return fmt.Sprintf("id=%d, name=%s, model=%s, manufacturer=%s, material=%s, speed=%d", a.Id, a.Name, a.Model, a.Manufacturer, a.Material, a.Speed)
}
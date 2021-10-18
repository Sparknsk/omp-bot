package autotransport

import "fmt"

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
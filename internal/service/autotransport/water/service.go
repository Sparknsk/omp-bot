package water

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/autotransport"
)

//TODO: возможно не лучшее место для хранения
var allEntities = []autotransport.Water{
	{Id: 1, Name: "One", Model: "X1", Manufacturer: "BMW", Material: "Iron", Speed: 10},
	{Id: 2, Name: "Two", Model: "X2", Manufacturer: "Toyota", Material: "Iron", Speed: 20},
	{Id: 3, Name: "Three", Model: "X3", Manufacturer: "Tesla", Material: "Wood", Speed: 30},
	{Id: 4, Name: "Four", Model: "X4", Manufacturer: "Ford", Material: "Wood", Speed: 40},
	{Id: 5, Name: "Five", Model: "X5", Manufacturer: "Kia", Material: "Plastic", Speed: 50},
}
var waterIdMax uint64

type DummyWaterService struct {}

func NewDummyWaterService() *DummyWaterService {
	return &DummyWaterService{}
}

func (s *DummyWaterService) Describe(waterId uint64) (*autotransport.Water, error) {
	for _, item := range allEntities {
		if item.Id == waterId {
			return &item, nil
		}
	}

	return nil, fmt.Errorf("сущность с ID=%d не найдена", waterId)
}

func (s *DummyWaterService) List(cursor uint64, limit uint64) []autotransport.Water {
	var entities []autotransport.Water

	offset := (cursor-1)*limit
	last := cursor*limit

	if last > uint64(s.Count()) {
		entities = allEntities[offset:]
	} else {
		entities = allEntities[offset:last]
	}

	return entities
}

func (s *DummyWaterService) Count() int {
	return len(allEntities)
}

func (s *DummyWaterService) Create(water autotransport.Water) (uint64, error) {
	for _, item := range allEntities {
		if item.Id > waterIdMax {
			waterIdMax = item.Id
		}
	}
	waterIdMax = waterIdMax+1

	water.Id = waterIdMax

	allEntities = append(allEntities, water)

	return water.Id, nil
}

func (s *DummyWaterService) Update(waterId uint64, water autotransport.Water) error {
	for key, item := range allEntities {
		if item.Id == waterId {
			allEntities[key] = water
			return nil
		}
	}

	return fmt.Errorf("сущность с ID=%d не найдена", waterId)
}

func (s *DummyWaterService) Remove(waterId uint64) (bool, error) {
	for key, item := range allEntities {
		if item.Id == waterId {
			allEntities = append(allEntities[:key], allEntities[key+1:]...)
			return true, nil
		}
	}

	return false, fmt.Errorf("сущность с ID=%d не найдена", waterId)
}

package water

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/autotransport"
	"sync/atomic"
)

type DummyWaterService struct {
	allEntities []autotransport.Water
	entityIdMax uint64
}

func NewDummyWaterService() *DummyWaterService {
	entities := []autotransport.Water{
		{Id: 1, Name: "One", Model: "X1", Manufacturer: "BMW", Material: "Iron", Speed: 10},
		{Id: 2, Name: "Two", Model: "X2", Manufacturer: "Toyota", Material: "Iron", Speed: 20},
		{Id: 3, Name: "Three", Model: "X3", Manufacturer: "Tesla", Material: "Wood", Speed: 30},
		{Id: 4, Name: "Four", Model: "X4", Manufacturer: "Ford", Material: "Wood", Speed: 40},
		{Id: 5, Name: "Five", Model: "X5", Manufacturer: "Kia", Material: "Plastic", Speed: 50},
	}

	s := &DummyWaterService{
		allEntities: entities,
		entityIdMax: uint64(len(entities)),
	}

	return s
}

func (s *DummyWaterService) Describe(waterId uint64) (*autotransport.Water, error) {
	for _, item := range s.allEntities {
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
		entities = s.allEntities[offset:]
	} else {
		entities = s.allEntities[offset:last]
	}

	return entities
}

func (s *DummyWaterService) Count() int {
	return len(s.allEntities)
}

func (s *DummyWaterService) Create(water autotransport.Water) (uint64, error) {
	atomic.AddUint64(&s.entityIdMax, 1)

	water.Id = atomic.LoadUint64(&s.entityIdMax)

	s.allEntities = append(s.allEntities, water)

	return water.Id, nil
}

func (s *DummyWaterService) Update(waterId uint64, water autotransport.Water) error {
	for key, item := range s.allEntities {
		if item.Id == waterId {
			s.allEntities[key] = water
			return nil
		}
	}

	return fmt.Errorf("сущность с ID=%d не найдена", waterId)
}

func (s *DummyWaterService) Remove(waterId uint64) (bool, error) {
	for key, item := range s.allEntities {
		if item.Id == waterId {
			s.allEntities = append(s.allEntities[:key], s.allEntities[key+1:]...)
			return true, nil
		}
	}

	return false, fmt.Errorf("сущность с ID=%d не найдена", waterId)
}

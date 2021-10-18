package water

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/autotransport"
)

type DummyWaterService struct {}

func NewDummyWaterService() *DummyWaterService {
	return &DummyWaterService{}
}

func (s *DummyWaterService) Describe(waterId uint64) (*autotransport.Water, error) {
	for _, item := range autotransport.AllEntities {
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
	count := s.Count()

	var entitiesSlice []autotransport.Water
	if last > uint64(count) {
		entitiesSlice = autotransport.AllEntities[offset:]
	} else {
		entitiesSlice = autotransport.AllEntities[offset:last]
	}

	for _, item := range entitiesSlice {
		entities = append(entities, item)
	}

	return entities
}

func (s *DummyWaterService) Count() int {
	return len(autotransport.AllEntities)
}

func (s *DummyWaterService) Create(water autotransport.Water) (uint64, error) {
	maxId := uint64(0)
	for _, item := range autotransport.AllEntities {
		if item.Id > maxId {
			maxId = item.Id
		}
	}

	water.Id = maxId+1

	autotransport.AllEntities = append(autotransport.AllEntities, water)

	return water.Id, nil
}

func (s *DummyWaterService) Update(waterId uint64, water autotransport.Water) error {
	for key, item := range autotransport.AllEntities {
		if item.Id == waterId {
			autotransport.AllEntities[key] = water
			return nil
		}
	}

	return fmt.Errorf("сущность с ID=%d не найдена", waterId)
}

func (s *DummyWaterService) Remove(waterId uint64) (bool, error) {
	for key, item := range autotransport.AllEntities {
		if item.Id == waterId {
			autotransport.AllEntities = append(autotransport.AllEntities[:key], autotransport.AllEntities[key+1:]...)
			return true, nil
		}
	}

	return false, fmt.Errorf("сущность с ID=%d не найдена", waterId)
}

package service

import (
	"github.com/JackDaniells/pack-service/domain/contracts"
	"github.com/JackDaniells/pack-service/domain/entity"
)

type packService struct {
	repository contracts.PackRepository
}

func NewPackService(repository contracts.PackRepository) *packService {
	return &packService{
		repository: repository,
	}
}

func (p *packService) Calculate(items int) []entity.Pack {
	var packs = p.repository.GetAll()

	packsNeeded := calculate(packs, items)

	var response []entity.Pack
	for key, value := range packsNeeded {
		response = append(response, entity.Pack{
			Size:     key,
			Quantity: value,
		})
	}

	return response

}

func calculate(packs []int, items int) map[int]int {
	packsNeeded := make(map[int]int)

	// fully divides items into packs
	for i := len(packs) - 1; i >= 0; i-- {
		pack := packs[i]
		packsUsed := items / pack
		if packsUsed > 0 {
			packsNeeded[pack] = packsUsed
			items %= pack
		}
	}

	// Distribute remaining items across available pack sizes
	for i := 0; i < len(packs) && items > 0; i++ {
		pack := packs[i]
		// If the pack can accommodate the remaining items, use it
		if items <= pack {
			packsNeeded[pack]++
			items = 0
		}
	}

	return packsNeeded
}

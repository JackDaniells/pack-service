package service

import "github.com/JackDaniells/pack-service/domain/entity"

type packService struct {
}

func NewPackService() *packService {
	return &packService{}
}

func (p *packService) Calculate(orderItems int) []entity.Pack {
	var packSizes = []int{250, 500, 1000, 2000, 5000}

	packsNeeded := make(map[int]int)

	// fully divides items into packs
	for i := len(packSizes) - 1; i >= 0; i-- {
		pack := packSizes[i]
		packsUsed := orderItems / pack
		if packsUsed > 0 {
			packsNeeded[pack] = packsUsed
			orderItems %= pack
		}
	}

	// Distribute remaining items across available pack sizes
	for i := 0; i < len(packSizes) && orderItems > 0; i++ {
		pack := packSizes[i]
		// If the pack can accommodate the remaining items, use it
		if orderItems <= pack {
			packsNeeded[pack]++
			orderItems = 0
		}
	}

	var response []entity.Pack
	for key, value := range packsNeeded {
		response = append(response, entity.Pack{
			Size:     key,
			Quantity: value,
		})
	}

	return response

}

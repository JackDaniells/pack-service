package service

import (
	"errors"
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

func (p *packService) Create(pack int) error {
	if pack <= 0 {
		return errors.New("invalid pack size")
	}

	p.repository.Create(pack)
	return nil
}

func (p *packService) GetAll() ([]entity.Pack, error) {
	var response []entity.Pack
	for _, pack := range p.repository.GetAll() {
		response = append(response, entity.Pack{Size: pack})
	}
	return response, nil
}

func (p *packService) Remove(pack int) error {
	p.repository.Remove(pack)
	return nil
}

func (p *packService) Addlist(packs []int) error {
	p.repository.AddList(packs)
	return nil
}

func (p *packService) RemoveList(packs []int) error {
	p.repository.RemoveList(packs)
	return nil
}

func (p *packService) Calculate(items int) (response []entity.Pack, err error) {
	if items <= 0 {
		return nil, errors.New("invalid items size")
	}

	var packs = p.repository.GetAll()
	if len(packs) == 0 {
		return
	}

	packsQty := calculatePacksQuantities(packs, items)

	for i, packQty := range packsQty {
		if packQty > 0 {
			response = append(response, entity.Pack{
				Size:     packs[i],
				Quantity: packQty,
			})
		}
	}

	return response, nil

}

func calculatePacksQuantities(packs []int, items int) []int {
	// try to calculate by finding the perfect distribution of items in packs
	t1, i := calcPrioritizingItemsDistribution(packs, items)
	if i == 0 {
		return t1
	}

	// if the perfect distribution doesn't exist, try to calculate it by prioritizing the use of the smallest number of packs
	packsQty := calcPrioritizingPacksDistribution(packs, items)
	return packsQty
}

func calcPrioritizingPacksDistribution(packs []int, items int) []int {
	packsQty := make([]int, len(packs))
	// calc the number of packages that can be used to store items
	for i, pack := range packs {
		if pack <= items {
			packsQty[i] = items / pack
			items %= pack
		}

		// compare pack difference with the smallest pack
		thisPackDiff := pack - items
		// if the difference is smaller, put it in the current package
		if thisPackDiff < packs[len(packs)-1] {
			packsQty[i]++
			items = 0
			break
		}
	}
	return packsQty
}

func calcPrioritizingItemsDistribution(packs []int, items int) ([]int, int) {
	packsQty := make([]int, len(packs))
	// calc the number of packages that can be used to store items
	for i, pack := range packs {
		packsQty[i] = items / pack
		items %= pack
	}
	return packsQty, items
}

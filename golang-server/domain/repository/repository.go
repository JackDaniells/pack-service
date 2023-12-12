package repository

import (
	"slices"
	"sort"
)

type packRepository struct {
	packs []int
}

func NewPackRepository() *packRepository {
	return &packRepository{
		packs: []int{250, 500, 1000, 2000, 5000},
	}
}

func (p *packRepository) GetAll() []int {
	packs := p.packs
	sort.Slice(packs, func(i, j int) bool {
		return packs[i] > packs[j]
	})
	return packs
}

func (p *packRepository) Create(pack int) {
	if !slices.Contains(p.packs, pack) {
		p.packs = append(p.packs, pack)
	}
}

func (p *packRepository) Remove(packToRemove int) {
	var packs []int
	for _, pack := range p.packs {
		if pack != packToRemove {
			packs = append(packs, pack)
		}
	}
	p.packs = packs
}

func (p *packRepository) AddList(packsToAdd []int) {
	for _, pack := range packsToAdd {
		if !slices.Contains(p.packs, pack) {
			p.packs = append(p.packs, pack)
		}
	}
}

func (p *packRepository) UpdateList(packs []int) {
	p.packs = packs
}

func (p *packRepository) RemoveList(packsToRemove []int) {

	for _, packToRemove := range packsToRemove {
		for i, pack := range p.packs {
			if pack == packToRemove {
				p.packs[i] = p.packs[len(p.packs)-1]
				p.packs = p.packs[:len(p.packs)-1]
				break
			}
		}
	}
}

package repository

type packRepository struct {
	packs []int
}

func NewPackRepository() *packRepository {
	return &packRepository{
		packs: []int{250, 500, 1000, 2000, 5000},
	}
}

func (p *packRepository) GetAll() []int {
	return p.packs
}

package contracts

import (
	"github.com/JackDaniells/pack-service/domain/entity"
)

type PackService interface {
	Calculate(orderItems int) ([]entity.Pack, error)
	GetAll() ([]entity.Pack, error)
	Create(pack int) error
	Remove(pack int) error
	Addlist(packs []int) error
	RemoveList(packs []int) error
	UpdateList(packs []int) error
}

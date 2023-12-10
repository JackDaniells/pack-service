package contracts

import "github.com/JackDaniells/pack-service/domain/entity"

type PackService interface {
	Calculate(orderItems int) ([]entity.Pack, error)
	Create(pack int) error
	Remove(pack int) error
}

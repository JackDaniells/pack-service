package contracts

import "github.com/JackDaniells/pack-service/domain/entity"

type PackService interface {
	Calculate(orderItems int) []entity.Pack
	Create(pack int)
	Remove(pack int)
}

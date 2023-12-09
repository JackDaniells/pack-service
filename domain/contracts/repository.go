package contracts

type PackRepository interface {
	GetAll() []int
	Create(pack int)
	Remove(packToRemove int)
}

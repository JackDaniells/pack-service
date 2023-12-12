package contracts

type PackRepository interface {
	GetAll() []int
	Create(pack int)
	Remove(packToRemove int)
	AddList(packsToAdd []int)
	RemoveList(packsToRemove []int)
	UpdateList(packs []int)
}

package contracts

type PackService interface {
	Calculate(orderItems int) map[int]int
}

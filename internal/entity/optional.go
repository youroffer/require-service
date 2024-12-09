package entity

type Optional[T any] struct {
	Value T
	Set   bool
}

func NewOptional[T any](value T) Optional[T] {
	return Optional[T]{
		Value: value,
		Set:   true,
	}
}

package lazy

type Producer[T any] func() T

type EProducer[T any] func() (T, error)

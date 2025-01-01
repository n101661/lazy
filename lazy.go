package lazy

type Loader[T any] struct {
	cache   T
	produce Producer[T]
}

func NewLoader[T any](f Producer[T]) *Loader[T] {
	return &Loader[T]{
		produce: f,
	}
}

func (loader *Loader[T]) Load() T {
	if loader.produce != nil {
		loader.cache = loader.produce()
		loader.produce = nil
	}
	return loader.cache
}

func (loader *Loader[T]) Loaded() bool {
	return loader.produce == nil
}

type ELoader[T any] struct {
	cache   T
	err     error
	produce EProducer[T]
}

func NewELoader[T any](f EProducer[T]) *ELoader[T] {
	return &ELoader[T]{
		produce: f,
	}
}

func (loader *ELoader[T]) Load() (T, error) {
	if loader.produce != nil {
		loader.cache, loader.err = loader.produce()
		loader.produce = nil
	}
	return loader.cache, loader.err
}

func (loader *ELoader[T]) Loaded() bool {
	return loader.produce == nil
}

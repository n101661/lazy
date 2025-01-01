package lazy

import (
	"sync"
)

type AtomicLoader[T any] struct {
	mu      sync.Mutex
	cache   T
	produce Producer[T]
}

func NewAtomicLoader[T any](f Producer[T]) *AtomicLoader[T] {
	return &AtomicLoader[T]{
		produce: f,
	}
}

func (loader *AtomicLoader[T]) Load() T {
	if loader.produce != nil {
		loader.mu.Lock()

		if loader.produce != nil {
			loader.cache = loader.produce()
			loader.produce = nil
		}

		loader.mu.Unlock()
	}
	return loader.cache
}

func (loader *AtomicLoader[T]) Loaded() bool {
	if loader.produce == nil {
		return true
	}

	loader.mu.Lock()
	defer loader.mu.Unlock()

	return loader.produce == nil
}

type AtomicELoader[T any] struct {
	mu      sync.Mutex
	cache   T
	err     error
	produce EProducer[T]
}

func NewAtomicELoader[T any](f EProducer[T]) *AtomicELoader[T] {
	return &AtomicELoader[T]{
		produce: f,
	}
}

func (loader *AtomicELoader[T]) Load() (T, error) {
	if loader.produce != nil {
		loader.mu.Lock()

		if loader.produce != nil {
			loader.cache, loader.err = loader.produce()
			loader.produce = nil
		}

		loader.mu.Unlock()
	}
	return loader.cache, loader.err
}

func (loader *AtomicELoader[T]) Loaded() bool {
	if loader.produce == nil {
		return true
	}

	loader.mu.Lock()
	defer loader.mu.Unlock()

	return loader.produce == nil
}

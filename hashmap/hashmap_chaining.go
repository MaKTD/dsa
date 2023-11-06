package hashmap

import (
	"golang.org/x/exp/slices"
	"hash/maphash"
)

const InitialCapacity = 32
const GrowThreshold = 3
const GrowFactor = 4
const ShrinkThreshold = 6
const ShrinkFactor = 4

type Entry[T any] struct {
	key   string
	value T
}

type StrWithChaining[T any] struct {
	size int
	seed maphash.Seed
	d    [][]Entry[T]
}

func (r *StrWithChaining[T]) Size() int {
	return r.size
}

func (r *StrWithChaining[T]) IsEmpty() bool {
	return r.size == 0
}

func (r *StrWithChaining[T]) Insert(key string, value T) {
	if ((r.size / len(r.d)) >= GrowThreshold) && len(r.d) != InitialCapacity {
		r.resize(len(r.d) * GrowFactor)
	} else if r.size > InitialCapacity && ((len(r.d) / r.size) >= ShrinkThreshold) {
		r.resize(len(r.d) / ShrinkFactor)
	}

	hash := r.hash(key)
	r.size += r.insert(r.d, hash, key, value)
}

func (r *StrWithChaining[T]) Delete(str string) bool {
	if ((r.size / len(r.d)) >= GrowThreshold) && len(r.d) != InitialCapacity {
		r.resize(len(r.d) * GrowFactor)
	} else if r.size > InitialCapacity && ((len(r.d) / r.size) >= ShrinkThreshold) {
		r.resize(len(r.d) / ShrinkFactor)
	}

	hash := r.hash(str)
	bucket := r.d[hash]
	if bucket == nil {
		return false
	}

	for i, entry := range bucket {
		if entry.key == str {
			r.d[hash] = slices.Delete(bucket, i, i+1)
			r.size -= 1
			return true
		}
	}

	return false
}

func (r *StrWithChaining[T]) Find(str string) (T, bool) {
	hash := r.hash(str)
	bucket := r.d[hash]
	if bucket == nil {
		var empty T
		return empty, false
	}

	for _, entry := range bucket {
		if entry.key == str {
			return entry.value, true
		}
	}

	var empty T
	return empty, false
}

func (r *StrWithChaining[T]) hash(str string) uint64 {
	return maphash.String(r.seed, str) % uint64(len(r.d))
}

func (r *StrWithChaining[T]) resize(newCap int) {
	newD := make([][]Entry[T], newCap)
	newSeed := maphash.MakeSeed()
	newSize := 0

	for _, bucket := range r.d {
		for _, entry := range bucket {
			newHash := maphash.String(newSeed, entry.key) % uint64(len(newD))
			newSize += r.insert(newD, newHash, entry.key, entry.value)
		}
	}

	r.seed = newSeed
	r.d = newD
	r.size = newSize
}

func (r *StrWithChaining[T]) insert(d [][]Entry[T], hash uint64, key string, value T) int {
	bucket := d[hash]
	if bucket == nil {
		d[hash] = []Entry[T]{Entry[T]{key: key, value: value}}
		return 1
	}

	for i, entry := range bucket {
		if entry.key == key {
			bucket[i] = Entry[T]{key: key, value: value}
			return 0
		}
	}

	d[hash] = append(d[hash], Entry[T]{key: key, value: value})
	return 1
}

func NewStrWithChaining[T any]() StrWithChaining[T] {
	return StrWithChaining[T]{
		seed: maphash.MakeSeed(),
		d:    make([][]Entry[T], InitialCapacity),
	}
}

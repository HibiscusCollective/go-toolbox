package fxmap

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"hash"
	"hash/fnv"
	"reflect"
)

type Hashable interface {
	Hash(fn hash.Hash) []byte
}

type Byteable interface {
	Bytes() []byte
}

type HashError interface {
	error
	HashFn() HashFunctionDescription
}

type HashMap[K, V any, SZ hashSize] struct {
	hash hash.Hash
}

type HashFunctionDescription struct {
	Algorithm string
	Size      int
	BlockSize int
}

type hashSize interface {
	uint32
}

type hashComparable interface {
	comparable
}

type pair[K, V any] struct {
	key K
	val V
}

type hashError struct {
	error
	hashFn HashFunctionDescription
}

func New[K hashComparable, V any](m map[K]V) (HashMap[K, V, uint32], error) {
	return NewWith32bitHashFunction(m, fnv.New32a())
}

func NewWith32bitHashFunction[K hashComparable, V any](m map[K]V, fn hash.Hash32) (HashMap[K, V, uint32], error) {
	hm := HashMap[K, V, uint32]{
		hash: fn,
	}

	for k := range m {
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		if err := enc.Encode(k); err != nil {
			panic(err)
		}

		_, err := fn.Write(buf.Bytes())
		if err != nil {
			return HashMap[K, V, uint32]{}, fmtHashError(fn, "failed to compute hash: %w", err)
		}
	}

	return hm, nil
}

func Invert[K, V comparable](m map[K]V) map[V]K {
	if m == nil {
		return nil
	}

	r := make(map[V]K, len(m))
	for k, v := range m {
		r[v] = k
	}

	return r
}

func fmtHashError(fn hash.Hash, str string, args ...any) hashError {
	return hashError{
		error: fmt.Errorf(str, args...),
		hashFn: HashFunctionDescription{
			Algorithm: reflect.TypeOf(fn).String(),
			Size:      fn.Size() * 8,
			BlockSize: fn.BlockSize() * 8,
		},
	}
}

func (e hashError) HashFn() HashFunctionDescription { return e.hashFn }

// Package fxslice provides functions to extend the slices package in the standard library
package fxslice

import (
	"errors"
	"fmt"
	"reflect"
)

// Transformer is a function that transforms a value of type T to a value of type R
type Transformer[T, R any] func(T) R

// TryTransformer is a function that transforms a value of type T to a value of type R but may return an error if it fails
type TryTransformer[T, R any] func(T) (R, error)

// Cast casts a slice of type T to a slice of type R
func Cast[T, R any](srcs []T) ([]R, error) {
	return TryTransform(srcs, func(src T) (R, error) {
		dstType := reflect.TypeOf(new(R)).Elem()
		v := reflect.ValueOf(src)
		if !v.IsValid() || !v.CanConvert(dstType) {
			return *new(R), NewCastError(src, dstType)
		}

		return v.Convert(dstType).Interface().(R), nil
	})
}

// Transform transforms a slice of type T to a slice of type R
func Transform[T, R any](src []T, fn Transformer[T, R]) []R {
	rs, _ := TryTransform(src, succeed(fn))

	return rs
}

// TryTransform transforms a slice of type T to a slice of type R
func TryTransform[T, R any](src []T, fn TryTransformer[T, R]) ([]R, error) {
	if src == nil || len(src) == 0 {
		return nil, nil
	}

	var out []R
	var errs []error

	for i, elem := range src {
		r, err := fn(elem)
		if err != nil {
			errs = append(errs, fmt.Errorf("index %d: %w", i, err))
		} else {
			out = append(out, r)
		}
	}

	if len(out) == 0 {
		return nil, errors.Join(errs...)
	}

	return out, errors.Join(errs...)
}

func succeed[T, R any](fn Transformer[T, R]) func(T) (R, error) {
	return func(t T) (R, error) {
		return fn(t), nil
	}
}

// CastError is an error that indicates that type casting failed
type CastError[T any] interface {
	error
	Src() T
	Dst() reflect.Type
}

type castError[T any] struct {
	src T
	dst reflect.Type
}

func NewCastError[T any](src T, dst reflect.Type) castError[T] {
	return castError[T]{src: src, dst: dst}
}

// Src returns the source value that caused the error.
func (c castError[T]) Src() T { return c.src }

// Dst returns the destination type that the source value could not be cast to.
func (c castError[T]) Dst() reflect.Type { return c.dst }

// Error implements error.
func (c castError[T]) Error() string {
	return fmt.Sprintf("cannot cast %[1]v of type %[1]T to %[2]s", c.src, c.dst)
}

package null

import "fmt"

type Null[T any] struct {
	value T
	valid bool
}

// New creates a new invalid Null[T].
func New[T any]() Null[T] {
	return Null[T]{
		valid: false,
	}
}

// From creates a new valid Null[T] from the given value.
func From[T any](from T) Null[T] {
	return Null[T]{
		value: from,
		valid: true,
	}
}

// FromPtr creates a new Null[T] which is invalid if the given pointer is not nil. Otherwise, the
// Null[T] value will be invalid.
func FromPtr[T any](from *T) Null[T] {
	if from != nil {
		return From(*from)
	}
	return New[T]()
}

// Map calls the transform function on the value inside src if it's valid. The result of the transform
// function is returned inside of a new Null[O]. If src is not valid, a new invalid Null[O] is returned.
//
// Example:
// 		null.Map(null.From(3), func (x int) string { return strconv.Itoa(x) })
func Map[T, O any](src Null[T], transform func(T) O) Null[O] {
	if src.valid {
		return From(transform(src.value))
	}
	return New[O]()
}

// IsValid returns whether this Null[T] is valid.
func (n Null[T]) IsValid() bool {
	return n.valid
}

// Get returns the value inside this Null[T]. If it's invalid, the function will panic.
func (n Null[T]) Get() T {
	if !n.valid {
		panic("Get called on invalid Null")
	}
	return n.value
}

// GetOr returns the value inside of this Null[T]. If it's invalid, the given value is returned.
func (n Null[T]) GetOr(value T) T {
	if !n.valid {
		return value
	}
	return n.value
}

// GetOrZero returns the value inside this Null[T]. If it's invalid, the zero value of type T is returned.
func (n Null[T]) GetOrZero() T {
	var zero T
	return n.GetOr(zero)
}

// GetPtr returns the value inside of this Null[T]. If it's invalid, nil is returned. Note that it's
// possible to modify the value contained in the Null[T] using the returned pointer.
func (n *Null[T]) GetPtr() *T {
	if !n.valid {
		return nil
	}
	return &n.value
}

// Set sets the value inside of the Null[T] and makes it valid.
func (n *Null[T]) Set(value T) {
	n.valid = true
	n.value = value
}

// SetPtr sets the value inside of the Null[T] and makes it valid if the given pointer is not nil.
func (n *Null[T]) SetPtr(value *T) {
	if value != nil {
		n.valid = true
		n.value = *value
	} else {
		var zero T
		n.valid = false
		n.value = zero
	}
}

func (n Null[T]) String() string {
	if !n.valid {
		return "<nil>"
	}
	return fmt.Sprintf("%v", n.value)
}

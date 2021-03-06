package null

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func addrof[T any](val T) *T {
	return &val
}

func TestNew(t *testing.T) {
	assert.Equal(t, Null[int]{valid: false, initialized: true, value: 0}, New[int]())
}

func TestFrom(t *testing.T) {
	assert.Equal(t, Null[int]{valid: true, initialized: true, value: 3}, From(3))
}

func TestFromPtr(t *testing.T) {
	assert.Equal(t, Null[int]{valid: true, initialized: true, value: 3}, FromPtr(addrof(3)))
	assert.Equal(t, Null[int]{valid: false, initialized: true, value: 0}, FromPtr[int](nil))
}

func TestMap(t *testing.T) {
	assert.Equal(t, Null[string]{valid: true, initialized: true, value: "3"}, Map(From(3), strconv.Itoa))
	assert.Equal(t, Null[string]{valid: false, initialized: true, value: ""}, Map(New[int](), func(x int) string { return strconv.Itoa(x) }))
}

func TestIsValid(t *testing.T) {
	assert.True(t, Null[int]{valid: true}.IsValid())
	assert.False(t, Null[int]{valid: false}.IsValid())
}

func TestIsInitialized(t *testing.T) {
	assert.True(t, Null[int]{initialized: true}.IsInitialized())
	assert.False(t, Null[int]{initialized: false}.IsInitialized())
}

func TestGet(t *testing.T) {
	assert.Equal(t, 3, From(3).Get())
	assert.Panics(t, func() {
		New[int]().Get()
	})
}

func TestGetOr(t *testing.T) {
	assert.Equal(t, 3, From(3).GetOr(33))
	assert.Equal(t, 33, New[int]().GetOr(33))
}

func TestGetOrZero(t *testing.T) {
	assert.Equal(t, "3", From("3").GetOrZero())
	assert.Equal(t, "", New[string]().GetOrZero())
}

func TestGetPtr(t *testing.T) {
	a := From(3)
	b := New[int]()

	assert.Equal(t, 3, *a.GetPtr())
	assert.Equal(t, (*int)(nil), b.GetPtr())

	*a.GetPtr() = 33
	assert.Equal(t, 33, a.Get())
}

func TestSet(t *testing.T) {
	a := New[int]()
	a.Set(3)

	assert.Equal(t, 3, a.Get())
}

func TestSetPtr(t *testing.T) {
	a := New[int]()
	b := New[int]()
	a.SetPtr(addrof(3))
	b.SetPtr(nil)

	assert.Equal(t, Null[int]{valid: true, initialized: true, value: 3}, a)
	assert.Equal(t, Null[int]{valid: false, initialized: true, value: 0}, b)
}

func TestUnmarshalJSON(t *testing.T) {
	var x Null[int]
	assert.NoError(t, x.UnmarshalJSON([]byte("1337")))
	assert.Equal(t, x.Get(), 1337)
}

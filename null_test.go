package null

import (
	"testing"
	"strconv"

	"github.com/stretchr/testify/assert"
)

func addrof[T any](val T) *T {
	return &val
}

func TestNew(t *testing.T) {
	assert.Equal(t, Null[int]{valid: false, value: 0}, New[int]())
}

func TestFrom(t *testing.T) {
	assert.Equal(t, Null[int]{valid: true, value: 3}, From(3))
}

func TestFromPtr(t *testing.T) {
	assert.Equal(t, Null[int]{valid: true, value: 3}, FromPtr(addrof(3)))
	assert.Equal(t, Null[int]{valid: false, value: 0}, FromPtr[int](nil))
}

func TestMap(t *testing.T) {
	assert.Equal(t, Null[string]{valid: true, value: "3"}, Map(From(3), func (x int) string { return strconv.Itoa(x)}))
	assert.Equal(t, Null[string]{valid: false, value: ""}, Map(New[int](), func (x int) string { return strconv.Itoa(x)}))
}

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMappingLine_GetDestination(t *testing.T) {
	t.Run("input is within mapping", func(t *testing.T) {
		mappingLine := NewMappingLine(0, 11, 42)
		in := NewRange(11, 3)
		out := mappingLine.GetDestination(in)
		assert.Equal(t, []Range{{Start: 0, End: 3}}, out)
	})

	t.Run("input is larger than mapping", func(t *testing.T) {
		mappingLine := NewMappingLine(0, 11, 3)
		in := NewRange(10, 5)
		out := mappingLine.GetDestination(in)
		assert.Equal(t, []Range{{Start: 0, End: 3}, {Start: 10, End: 11}, {Start: 14, End: 15}}, out)
	})
}

func TestRange_Split(t *testing.T) {
	t.Run("input is within mapping", func(t *testing.T) {
		mapping := Range{Start: 45, End: 60}
		in := Range{Start: 50, End: 55}
		out := mapping.Split(in)
		assert.Equal(t, []Range{{Start: 50, End: 55}}, out)
	})

	t.Run("input is larger than mapping", func(t *testing.T) {
		mapping := Range{Start: 50, End: 55}
		in := Range{Start: 45, End: 60}
		out := mapping.Split(in)
		assert.Equal(t, []Range{{Start: 50, End: 55}, {Start: 45, End: 50}, {Start: 55, End: 60}}, out)
	})

	t.Run("input does not match mapping", func(t *testing.T) {
		mapping := Range{Start: 50, End: 55}
		in := Range{Start: 10, End: 15}
		out := mapping.Split(in)
		assert.Equal(t, []Range{{Start: 10, End: 15}}, out)
	})
}

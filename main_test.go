package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMapping(t *testing.T) {
	tcs := []struct {
		Lines    []MappingLine
		Expected Mapping
	}{
		{
			Lines: []MappingLine{
				{Destination: 50, Source: 98, Length: 2},
			},
			Expected: Mapping{98: 50, 99: 51},
		},
		{
			Lines: []MappingLine{
				{Destination: 50, Source: 98, Length: 2},
				{Destination: 52, Source: 50, Length: 3},
			},
			Expected: Mapping{50: 52, 51: 53, 52: 54, 98: 50, 99: 51},
		},
	}

	for _, tc := range tcs {
		actual := NewMapping(tc.Lines)
		assert.Equal(t, tc.Expected, actual)
	}
}

func TestMapping_GetDst(t *testing.T) {
	mapping := Mapping{50: 52, 51: 53}
	t.Run("when mapping found, it is returned", func(t *testing.T) {
		dst := mapping.GetDst(51)
		assert.Equal(t, 53, dst)
	})
	t.Run("when mapping not found, input is returned", func(t *testing.T) {
		dst := mapping.GetDst(75)
		assert.Equal(t, 75, dst)
	})
}

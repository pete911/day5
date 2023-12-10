package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	// MappingNames length has to be exactly the same as Mappings, name is retrieved by index
	MappingNames = []string{
		"soil",
		"fertilizer",
		"water",
		"light",
		"temperature",
		"humidity",
		"location",
	}

	// Mappings length has to be exactly the same as MappingNames, name is retrieved by index
	Mappings = []Mapping{
		NewMapping([]MappingLine{
			{Destination: 50, Source: 98, Length: 2},
			{Destination: 52, Source: 50, Length: 48},
		}),
		NewMapping([]MappingLine{
			{Destination: 0, Source: 15, Length: 37},
			{Destination: 37, Source: 52, Length: 2},
			{Destination: 39, Source: 0, Length: 15},
		}),
		NewMapping([]MappingLine{
			{Destination: 49, Source: 53, Length: 8},
			{Destination: 0, Source: 11, Length: 42},
			{Destination: 42, Source: 0, Length: 7},
			{Destination: 57, Source: 7, Length: 4},
		}),
		NewMapping([]MappingLine{
			{Destination: 88, Source: 18, Length: 7},
			{Destination: 18, Source: 25, Length: 70},
		}),
		NewMapping([]MappingLine{
			{Destination: 44, Source: 77, Length: 23},
			{Destination: 81, Source: 45, Length: 19},
			{Destination: 68, Source: 64, Length: 13},
		}),
		NewMapping([]MappingLine{
			{Destination: 0, Source: 69, Length: 1},
			{Destination: 1, Source: 0, Length: 69},
		}),
		NewMapping([]MappingLine{
			{Destination: 60, Source: 56, Length: 37},
			{Destination: 56, Source: 93, Length: 4},
		}),
	}
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		dst, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("cannot convert input %s to integer", arg)
			os.Exit(1)
		}

		fmt.Printf("seed %d   ", dst)
		for i, name := range MappingNames {
			dst = Mappings[i].GetDst(dst)
			fmt.Printf("%s: %d   ", name, dst)
		}
		fmt.Println()
	}
}

type MappingLine struct {
	Destination int
	Source      int
	Length      int
}

// Mapping contains source and destination mappings
type Mapping map[int]int

// NewMapping return new mapping from supplied lines
func NewMapping(lines []MappingLine) Mapping {
	out := make(map[int]int)
	for _, line := range lines {
		for i := 0; i < line.Length; i++ {
			out[line.Source+i] = line.Destination + i
		}
	}
	return out
}

// GetDst returns destination if it is in mapping or input (src) if there's no mapping
func (m Mapping) GetDst(src int) int {
	if out, ok := m[src]; ok {
		return out
	}
	return src
}

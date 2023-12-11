package main

import "fmt"

// Output where key is name (e.g. location) and values are ranges
type Output map[string][]Range

func main() {
	out := Output{"seeds": Seeds}
	for i, name := range MappingNames {
		dst := Seeds
		dst = Mappings[i].GetDestination(dst)
		out[name] = dst
	}

	fmt.Println("seeds")
	fmt.Printf("%+v\n", out["seeds"])
	fmt.Println()
	for _, name := range MappingNames {
		fmt.Println(name)
		fmt.Printf("%+v\n", out[name])
		fmt.Println()
	}

	var lowestLocation int
	for i, v := range out["location"] {
		if i == 0 {
			lowestLocation = v.Start
			continue
		}
		if lowestLocation > v.Start {
			lowestLocation = v.Start
		}
	}
	fmt.Printf("lowest location: %d\n", lowestLocation)
}

// Mapping contains source and destination mappings
type Mapping []MappingLine

func (m Mapping) GetDestination(source []Range) []Range {
	var out []Range
	for _, mappingLine := range m {
		for _, sourceLine := range source {
			out = append(out, mappingLine.GetDestination(sourceLine)...)
		}
	}
	return out
}

type MappingLine struct {
	Destination Range
	Source      Range
}

func NewMappingLine(destination, source, length int) MappingLine {
	return MappingLine{
		Destination: NewRange(destination, length),
		Source:      NewRange(source, length),
	}
}

func (m MappingLine) GetDestination(source Range) []Range {
	if m.Source.Contains(source) {
		increment := m.Destination.Start - m.Source.Start
		out := m.Source.Split(source)
		out[0].Start = out[0].Start + increment
		out[0].End = out[0].End + increment
		return out
	}
	return []Range{source}
}

type Range struct {
	Start int
	End   int
}

func NewRange(start, length int) Range {
	return Range{Start: start, End: start + length}
}

func (r Range) Contains(v Range) bool {
	return r.Start <= v.End && v.Start <= r.End
}

// Split returns intersecting ranges, if there is no overlapping, input is returned,
// otherwise first range is the one that overlaps
func (r Range) Split(v Range) []Range {
	// no matching intersection
	if !r.Contains(v) {
		return []Range{v}
	}

	var unmapped []Range
	var start, end int
	if v.Start < r.Start {
		start = r.Start
		unmapped = append(unmapped, Range{Start: v.Start, End: r.Start})
	}
	if v.Start >= r.Start {
		start = v.Start
	}
	if v.End > r.End {
		end = r.End
		unmapped = append(unmapped, Range{Start: r.End, End: v.End})
	}
	if v.End <= r.End {
		end = v.End
	}
	mapped := Range{Start: start, End: end}
	if len(unmapped) == 0 {
		return []Range{mapped}
	}
	return append([]Range{mapped}, unmapped...)
}

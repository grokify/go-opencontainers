package manifest

import (
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type Descriptors []v1.Descriptor

func (ds Descriptors) DigestsMap() map[string]int {
	out := map[string]int{}
	for _, d := range ds {
		de := DescriptorEdit{Descriptor: d}
		out[string(de.Digest())] += 1
	}
	return out
}

func (ds Descriptors) DigestsUnique() bool {
	return histogramUnique(ds.DigestsMap())
}

func (ds Descriptors) TitlesMap() map[string]int {
	out := map[string]int{}
	for _, d := range ds {
		de := DescriptorEdit{Descriptor: d}
		out[de.Title()] += 1
	}
	return out
}

func (ds Descriptors) TitlesUnique() bool {
	return histogramUnique(ds.TitlesMap())
}

func histogramUnique(m map[string]int) bool {
	for _, v := range m {
		if v != 1 && v != 0 {
			return false
		}
	}
	return true
}

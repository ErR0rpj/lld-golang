package filters

import "main/internal/core"

type AndFilter struct {
	Filters []core.FileFilter
}

func (f *AndFilter) Matches(file core.File) bool {
	for _, filter := range f.Filters {
		if !filter.Matches(file) {
			return false
		}
	}
	return true
}

type OrFilter struct {
	Filters []core.FileFilter
}

func (f *OrFilter) Matches(file core.File) bool {
	for _, filter := range f.Filters {
		if filter.Matches(file) {
			return true
		}
	}
	return false
}

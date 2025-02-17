package filters

import (
	"main/internal/core"
	"strings"
)

type NameFilter struct {
	Pattern string
}

func (f *NameFilter) Matches(file core.File) bool {
	return strings.Contains(file.Name, f.Pattern)
}

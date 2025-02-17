package filters

import "main/internal/core"

type SizeGreaterThan struct {
	Size int64
}

func (f *SizeGreaterThan) Matches(file core.File) bool {
	return file.Size > f.Size
}

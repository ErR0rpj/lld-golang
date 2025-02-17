package main

import (
	"fmt"
	"main/internal/core"
	"main/internal/filters"
)

func main() {
	dir := &core.Directory{
		Files: []core.File{
			{Name: "resume.pdf", Size: 2 * 1024 * 1024},
			{Name: "importantdoc.pdf", Size: 10 * 1024 * 1024},
			{Name: "mostimportantinfo.pdf", Size: 100 * 1024 * 1024},
		},
	}

	filter := &filters.AndFilter{
		Filters: []core.FileFilter{
			&filters.NameFilter{Pattern: "important"},
			&filters.SizeGreaterThan{Size: 10},
		},
	}

	results := dir.Search(filter)
	for _, result := range results {
		fmt.Println(result.Name)
	}

}

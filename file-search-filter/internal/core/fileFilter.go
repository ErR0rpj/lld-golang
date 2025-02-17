package core

type FileFilter interface {
	Matches(file File) bool
}

func (d *Directory) Search(filter FileFilter) []File {
	var results []File

	for _, file := range d.Files {
		if filter.Matches(file) {
			results = append(results, file)
		}
	}

	return results
}

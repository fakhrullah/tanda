package collection

import "io/fs"

func Includes(collections []string, searchElement string) bool {
	for _, element := range collections {
		if element == searchElement {
			return true
		}
	}

	return false
}

func Map(files []fs.FileInfo) []string {
	var allFilenames []string
	for _, file := range files {
		allFilenames = append(allFilenames, file.Name())
	}
	return allFilenames
}

package day07

import (
	"path"
)

func FindDirsUpTo100K(input string) (int, error) {
	fs, err := ParseFileSystem(input)
	if err != nil {
		return 0, err
	}

	dirSizes := make(map[string]int)

	populateDirSizeMap(dirSizes, fs.root, "/")

	result := 0

	for _, size := range dirSizes {
		if size <= 100000 {
			result += size
		}
	}

	return result, nil
}

func populateDirSizeMap(dsm map[string]int, dir *directory, dirName string) {
	for name, de := range dir.entries {
		if dir, ok := de.(*directory); ok {
			thisDirPath := path.Join(dirName, name)
			dsm[thisDirPath] = dir.totalSize()
			populateDirSizeMap(dsm, dir, thisDirPath)
		}
	}
}

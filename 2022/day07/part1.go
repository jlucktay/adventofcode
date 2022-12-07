package day07

func FindDirsUpTo100K(input string) (int, error) {
	fs, err := ParseFileSystem(input)
	if err != nil {
		return 0, err
	}

	dirSizes := make(map[string]int)

	populateDirSizeMap(dirSizes, fs.root)

	result := 0

	for _, size := range dirSizes {
		if size <= 100000 {
			result += size
		}
	}

	return result, nil
}

func populateDirSizeMap(dsm map[string]int, dir *directory) {
	for name, de := range dir.entries {
		if de.entryType() == det_dir {
			dsm[name] = de.totalSize()

			populateDirSizeMap(dsm, de.(*directory))
		}
	}
}

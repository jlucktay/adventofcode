package day07

import (
	"errors"
	"sort"
)

const (
	TOTAL_DISK_SPACE  = 70_000_000
	NEED_UNUSED_SPACE = 30_000_000
)

func FindDirToDelete(input string) (int, error) {
	fs, err := ParseFileSystem(input)
	if err != nil {
		return 0, err
	}

	dirSizes := make(map[string]int)
	populateDirSizeMap(dirSizes, fs.root, "/")

	currentUnusedSpace := TOTAL_DISK_SPACE - fs.root.totalSize()
	needToDeleteSize := NEED_UNUSED_SPACE - currentUnusedSpace
	deletionCandidates := make([]int, 0)

	for _, dirSize := range dirSizes {
		if dirSize > needToDeleteSize {
			deletionCandidates = append(deletionCandidates, dirSize)
		}
	}

	if len(deletionCandidates) == 0 {
		return 0, errors.New("finding candidates for deletion")
	}

	sort.Ints(deletionCandidates)

	return deletionCandidates[0], nil
}

// Package day09 for Advent of Code 2024, day 9, part 2.
// https://adventofcode.com/2024/day/9
package day09

func Part2(input string) (int64, error) {
	d, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	d.compactFiles()

	return d.checksum(true), nil
}

func (d Disk) fileStart(workBackFrom int) int {
	fileID := d[workBackFrom]

	for i := workBackFrom; i > 1; i-- {
		if d[i-1] != fileID {
			return i
		}
	}

	return 0
}

func (d Disk) firstFreeBlock(requiredSize int) int {
	currentFreeBlockSize := 0

	for i := range d {
		switch d[i] {
		case -1:
			currentFreeBlockSize++

			if currentFreeBlockSize >= requiredSize {
				return i - currentFreeBlockSize + 1
			}

		default:
			currentFreeBlockSize = 0
		}
	}

	return -1
}

func (d Disk) moveFile(fileStart, blockStart, size int) {
	for i := range size {
		d[blockStart+i] = d[fileStart+i]
		d[fileStart+i] = -1
	}
}

func (d Disk) compactFiles() {
	for i := len(d) - 1; i >= 0; i-- {
		if d[i] == -1 {
			continue
		}

		fileStart := d.fileStart(i)

		size := i - fileStart + 1

		freeBlockStart := d.firstFreeBlock(size)

		if freeBlockStart != -1 && freeBlockStart < fileStart {
			d.moveFile(fileStart, freeBlockStart, size)
		}

		i = fileStart
	}
}

package day07

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type FileSystem struct {
	root, cwd *directory
}

type dirEntryType string

const (
	det_dir  dirEntryType = "directory"
	det_file dirEntryType = "file"
)

type dirEntry interface {
	prettyPrint(name string, indent int) string
	entryType() dirEntryType
	totalSize() int
}

type directory struct {
	parent  *directory
	entries map[string]dirEntry
}

func (d *directory) prettyPrint(name string, indent int) string {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("%*s- %s (dir)\n", indent, "", name))

	// Make sorted list of directories/files and print those with an indent.
	entryNames := make([]string, 0)

	for name := range d.entries {
		entryNames = append(entryNames, name)
	}

	sort.Strings(entryNames)

	for _, en := range entryNames {
		sb.WriteString(d.entries[en].prettyPrint(en, indent+2))
	}

	return sb.String()
}

func (d *directory) entryType() dirEntryType { return det_dir }

func (d *directory) totalSize() int {
	size := 0

	for _, de := range d.entries {
		size += de.totalSize()
	}

	return size
}

type file struct {
	size int
}

func (f *file) prettyPrint(name string, indent int) string {
	return fmt.Sprintf("%*s- %s (file, size=%d)\n", indent, "", name, f.size)
}

func (f *file) entryType() dirEntryType { return det_file }

func (f *file) totalSize() int {
	return f.size
}

func NewFileSystem() *FileSystem {
	newFs := &FileSystem{}
	newFs.root = &directory{entries: map[string]dirEntry{}}
	newFs.cwd = newFs.root

	return newFs
}

func ParseFileSystem(input string) (*FileSystem, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	resultFS := NewFileSystem()

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "$ ") {
			if err := resultFS.parseCommand(line[2:]); err != nil {
				return nil, err
			}
		} else if strings.HasPrefix(line, "dir ") {
			if err := resultFS.parseDirectory(line[4:]); err != nil {
				return nil, err
			}
		} else {
			if err := resultFS.parseFile(line); err != nil {
				return nil, err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanning input: %v", err)
	}

	return resultFS, nil
}

func (fs *FileSystem) parseCommand(input string) error {
	xInput := strings.Split(input, " ")
	if len(xInput) == 0 || len(xInput) > 2 {
		return fmt.Errorf("parsing command '%s'", input)
	}

	switch xInput[0] {
	case "cd":
		targetDir := xInput[1]
		if err := fs.changeDirectory(targetDir); err != nil {
			return err
		}
	case "ls":
		// This is a no-op; the scanner will parse in subsequent directories and files until the next command.
	default:
		return fmt.Errorf("unknown command '%s'", xInput[0])
	}

	return nil
}

func (fs *FileSystem) parseDirectory(input string) error {
	if fs.cwd.entries == nil {
		fs.cwd.entries = make(map[string]dirEntry)
	}

	fs.cwd.entries[input] = &directory{
		parent: fs.cwd,
	}

	return nil
}

func (fs *FileSystem) changeDirectory(targetDir string) error {
	if targetDir == "/" {
		fs.cwd = fs.root

		return nil
	}

	if targetDir == ".." {
		fs.cwd = fs.cwd.parent

		return nil
	}

	for name, entry := range fs.cwd.entries {
		if entry.entryType() != det_dir {
			continue
		}

		if name == targetDir {
			cd, ok := fs.cwd.entries[targetDir].(*directory)
			if !ok {
				return fmt.Errorf("could not assert key '%s' for target dir from entries '%+v'",
					targetDir, fs.cwd.entries)
			}

			fs.cwd = cd

			return nil
		}
	}

	return fmt.Errorf("change to non-existant '%s' directory", targetDir)
}

func (fs *FileSystem) parseFile(input string) error {
	xInput := strings.Split(input, " ")
	if len(xInput) == 0 || len(xInput) > 2 {
		return fmt.Errorf("parsing file '%s'", input)
	}

	fileSize, err := strconv.ParseInt(xInput[0], 10, 32)
	if err != nil {
		return fmt.Errorf("parsing file size '%s': %w", xInput[0], err)
	}

	newFile := &file{size: int(fileSize)}

	if fs.cwd.entries == nil {
		fs.cwd.entries = make(map[string]dirEntry)
	}

	fs.cwd.entries[xInput[1]] = newFile

	return nil
}

func (fs *FileSystem) String() string {
	return fs.root.prettyPrint("/", 0)
}

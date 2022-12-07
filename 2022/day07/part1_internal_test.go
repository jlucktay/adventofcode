package day07

import (
	"testing"

	"github.com/matryer/is"
)

const INPUT = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`

func TestTotalSize(t *testing.T) {
	is := is.New(t)

	fs, err := ParseFileSystem(INPUT)
	is.NoErr(err)

	is.Equal(584, fs.root.entries["a"].(*directory).entries["e"].(*directory).totalSize())
	is.Equal(94853, fs.root.entries["a"].(*directory).totalSize())
	is.Equal(24933642, fs.root.entries["d"].(*directory).totalSize())
	is.Equal(48381165, fs.root.totalSize())
}

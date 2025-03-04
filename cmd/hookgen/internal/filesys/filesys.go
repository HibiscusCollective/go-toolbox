// Package filesys provides an adapter layer for interacting with the filesystem
package filesys

import (
	"io"

	"github.com/spf13/afero"
)

// FS is an interface that defines the logic for interacting with the filesystem
type FS interface {
	FileCreator
	DirMaker
}

// File is an interface that defines the logic for reading and writing a file
type File interface {
	io.ReadSeekCloser
	io.WriteCloser
}

// FileSystemReader is an interface that defines the logic for listing files and reading a file
type FileSystemReader interface {
	ListFiles(path string) ([]string, error)
	ReadFile(path string) (File, error)
}

// FileCreator is an interface that defines the logic for creating a writeable file
type FileCreator interface {
	Create(name string) (File, error)
}

// DirMaker is an interface that defines the logic for creating a directory
type DirMaker interface {
	MkdirAll(path string) error
}

type aferoFS struct {
	fs afero.Fs
}

// New creates a new instance of a FS adapter
func New(afs afero.Fs) FS {
	return aferoFS{fs: afs}
}

func (a aferoFS) Create(name string) (File, error) {
	return a.fs.Create(name)
}

func (a aferoFS) MkdirAll(path string) error {
	return a.fs.MkdirAll(path, 0o644)
}

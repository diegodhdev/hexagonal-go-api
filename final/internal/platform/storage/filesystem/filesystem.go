package filesystem

import (
	"os"
)

type Filesystem struct {
	storageFolder string
}

func NewFilesystem(storagePath string) *Filesystem {
	return &Filesystem{
		storageFolder: storagePath,
	}
}

func (f *Filesystem) Save(content []byte, filename string) error {
	err := os.WriteFile(f.storageFolder+filename, content, 0777)
	if err != nil {

		return err
	}

	return nil
}

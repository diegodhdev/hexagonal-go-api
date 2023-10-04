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

func (f *Filesystem) Save(content []byte, folder string, filename string) error {

	if _, err := os.Stat(f.storageFolder + folder + filename); os.IsNotExist(err) {
		err := os.MkdirAll(f.storageFolder+folder, 0700)
		if err != nil {
			return err
		}
	}

	err := os.WriteFile(f.storageFolder+folder+filename, content, 0777)
	if err != nil {
		return err
	}

	return nil
}

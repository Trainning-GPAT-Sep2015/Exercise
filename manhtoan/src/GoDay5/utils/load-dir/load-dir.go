package loadDir

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func ShowFileList(dir string) ([]os.FileInfo, error) {
	absPath, err := filepath.Abs(dir)
	if err != nil {
		return []os.FileInfo{}, err
	}

	files, err := ioutil.ReadDir(absPath)
	if err != nil {
		return []os.FileInfo{}, err
	}

	return files, nil
}

package utils

import (
	"fmt"
	"gopkg.in/src-d/go-billy.v4"
	"path/filepath"
)

func printDir(fs billy.Filesystem, dir string) error {
	fis, err := fs.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, fi := range fis {
		fpath := filepath.Join(dir, fi.Name())
		fmt.Println(fpath)
		if fi.IsDir() {
			if err := printDir(fs, fpath); err != nil {
				return err
			}
		}
	}
	return nil
}

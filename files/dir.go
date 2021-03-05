package files

import (
	"os"
	"path/filepath"
)

// 创建文件夹
func CreateDir(path string) error {
	// 路径转换，window下是"\",转为"/"
	path = filepath.ToSlash(path)

	// 获取到dir
	dir := filepath.Dir(path)

	if !IsExist(dir) {
		err := os.MkdirAll(dir, 0744)
		if err != nil {
			return err
		}
	}
	return nil
}

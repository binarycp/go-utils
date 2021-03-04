package files

import "os"

// 创建文件夹
func CreateDir(dir string) error {
	if !IsExist(dir) {
		err := os.MkdirAll(dir, 0744)
		if err != nil {
			return err
		}
	}
	return nil
}

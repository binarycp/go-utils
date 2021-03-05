//Package 文件工具包
package files

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 读取文件内容
func Read(path string) (string, error) {
	read, err := ReadAsBytes(path)
	return string(read), err
}

// 读取文件内容
func ReadAsBytes(path string) ([]byte, error) {
	open, err := os.Open(path)
	if open != nil {
		defer open.Close()
	}
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(open)
}

// 读取文件的每一行
func ReadLine(path string) ([][]byte, error) {
	open, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer open.Close()

	reader := bufio.NewReader(open)
	var result [][]byte

	for {
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			break
		}

		for isPrefix && err == nil {
			var bs []byte
			bs, isPrefix, err = reader.ReadLine()
			line = append(line, bs...)
		}

		// 不能直接使用line，否则最后一行会替换第一行代码
		lineStr := string(line)

		result = append(result, []byte(lineStr))
	}

	return result, nil
}

// 创建文件（包括文件夹）
func Create(path string) error {
	if !IsExist(path) {
		// 路径转换，window下是"\",转为"/"
		path = filepath.ToSlash(path)
		err := CreateDir(path)
		if err != nil {
			return err
		}

		_, err = os.Create(path)
		if err != nil {
			return err
		}
	}

	return nil
}

// 打开某个文件，文件不存在会创建,注意这里没有关闭文件句柄的
func Open(path string, flag int, perm os.FileMode) (*os.File, error) {
	err := Create(path)
	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(path, flag, perm)
	if err != nil {
		return nil, err
	}

	return file, nil
}

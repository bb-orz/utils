package xos

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 获取文件的绝对路径
func GetAbsPath() (string,error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}


// 获取文件所属目录
func GetFileDir(filePath string) string {
	return filepath.Dir(filePath)
}


// 创建文件所属目录
func CreateFileDir(filePath string) error {
	// 创建本地目录
	err := os.MkdirAll(GetFileDir(filePath), 0777)
	if err != nil {
		return err
	}
	return nil
}


// 简单复制文件，用于小文件复制
func SimpleCopyFile(src string, dst string) error {

	bytes, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dst, bytes, 0666)
	if err != nil {
		return err
	}

	return nil
}

// 带缓冲的文件复制，用于大文件
func CopyFile(src string, dst string) error {
	// 只读模式打开文件
	srcFile, err := os.OpenFile(src, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// 只写模式打开创建文件
	dstFile, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// 创建缓冲区[]byte
	var n int = 0
	bufbytes := make([]byte, 1024)
	reader := bufio.NewReader(srcFile)
	writer := bufio.NewWriter(dstFile)

	for {
		n++
		_, err = reader.Read(bufbytes)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		_, err := writer.Write(bufbytes)
		if err != nil {
			return err
		}
	}

	return nil
}
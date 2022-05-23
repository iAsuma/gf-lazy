package qutil

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// PathExists 判断路径（文件/文件夹）是否存在
func PathExists(path string) (bool, os.FileInfo) {
	f, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true, f
		}
		return false, nil
	}

	return true, f
}

// IsExistPath 判断路径（文件/文件夹）是否存在
func IsExistPath(path string) bool {
	is, _ := PathExists(path)
	return is
}

// IsDir 判断是否文件夹
func IsDir(dir string) bool {
	is, f := PathExists(dir)
	return is && f.IsDir()
}

// IsFile 判断是否文件
func IsFile(file string) bool {
	is, f := PathExists(file)
	return is && !f.IsDir()
}

// WriteFile 写文件
func WriteFile(filename string, content string) {
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = f.Write([]byte(content))
	}
}

// GetCurrentPath 获取当前路径
func GetCurrentPath() string {
	dir, err := os.Getwd() //当前的目录
	if err != nil {
		dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Println("can not get current path")
		}
	}
	return dir
}

package main

import (
	"fmt"
	"os"
	//"errors"
	"io/ioutil"
	"path/filepath"
)

func main() {
	dirs, err := ioutil.ReadDir(".")
	
	if err != nil {
		fmt.Println("无法找到目录")
	}
	
	for _, fi := range dirs {
		//fmt.Println(fi.Name())
		if !fi.IsDir() {
			continue
		}
	}
	
	directorys := make([]string, 0, 50)
	
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			directorys = append(directorys, info.Name())
		}
		return nil
	})
	
	for _, dir := range directorys {
		fmt.Println(dir)
	}
}
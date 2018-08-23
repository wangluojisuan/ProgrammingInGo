﻿package main

import (
	"os"
	"io"
	"fmt"
	"strings"
	"strconv"
	"errors"
	"path/filepath"
	"github.com/xuri/excelize"
	"github.com/gonutz/w32"
)

const (
	RESLUT_FILENAME = "Result.txt"
	XLSX_FILENAME = "1.xlsx"
)

type FileInfo struct {
	FileName string
	FileVersion string
}

func GetAllSameLevelDirectory() ([]string, error) {
	return nil, nil
}

func GetFileList() []string {
	fileList := make([]string, 0, 30)
	
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fileList = append(fileList, info.Name())
		}
		return nil
	})
	
	if err != nil {
		return []string{}
	}
	
	return fileList
}

// 获取文件版本
func GetFileVersion(path string) (version string, err error) {
	size := w32.GetFileVersionInfoSize(path)
    if size <= 0 {
        err = errors.New("GetFileVersionInfoSize failed")
		return
    }
	
	info := make([]byte, size)
    ok := w32.GetFileVersionInfo(path, info)
    if !ok {
        err = errors.New("GetFileVersionInfo failed")
		return
    }
	
	fixed, ok := w32.VerQueryValueRoot(info)
    if !ok {
        panic("VerQueryValueRoot failed")
    }
    fileVersion := fixed.FileVersion()
	strFileVersions := [4]string{}
	strFileVersions[0] = strconv.FormatUint((fileVersion&0xFFFF000000000000)>>48, 10)
	strFileVersions[1] = strconv.FormatUint((fileVersion&0x0000FFFF00000000)>>32, 10)
	strFileVersions[2] = strconv.FormatUint((fileVersion&0x00000000FFFF0000)>>16, 10)
	strFileVersions[3] = strconv.FormatUint((fileVersion&0x000000000000FFFF)>>0, 10)

	//fmt.Println(size)
	version = strings.Join(strFileVersions[:], ".")
	return
}

// 文件是否存在
func isFileExist(filename string) bool {
	var isExist = true
	
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		isExist = false
	}
	
	return isExist
}

// 写入日志
func WriteResult(log string) {
	var f *os.File
	var err error
	
	if isFileExist(RESLUT_FILENAME) {
		f, err = os.OpenFile(RESLUT_FILENAME, os.O_APPEND, 0666)		
	} else {
		f, err = os.OpenFile(RESLUT_FILENAME, os.O_RDWR|os.O_CREATE, 0666)
	}
	if err != nil {
		return
	}
	defer f.Close()
	
	_, err = io.WriteString(f, log+"\r\n")
}

func main() {
	ignore := os.Remove(RESLUT_FILENAME)
	if ignore!=nil {
		fmt.Println(ignore)
	} 
	
	isXlsxFileExist := isFileExist(XLSX_FILENAME)
	if !isXlsxFileExist {
		WriteResult("没有找到输入的"+XLSX_FILENAME)
	}
	
	xlsx, err := excelize.OpenFile(XLSX_FILENAME)
	if err!=nil {
		fmt.Println(err)
		return
	}
	//cell := xlsx.GetCellValue("ATS系统发布项信息表", "A2")
	//fmt.Println(cell)
	
	rows := xlsx.GetRows("ATS系统发布项信息表")
	//capacity := len(rows)
	//fileSlice := make([]FileInfo, capacity)
	fileSlice := []FileInfo{}
	for index, row := range rows {
		if index<=2 {
			continue
		}
		strVersion := strings.Trim(row[2], " ")
		strVersion = strings.Trim(strVersion, "\n")
		strNames := strings.Trim(row[5], " ")
		if strVersion=="" || strNames=="" {
			continue
		}
		names := strings.Split(strNames, "\n")
		if len(names)>1 {
			for _, name := range names {
				if len(name)!=len([]rune(name)) {
					//fmt.Println("检测到中文字符串")
					continue
				}
				fileSlice = append(fileSlice, FileInfo{FileName:name, FileVersion:strVersion})
				//fmt.Println(index, "\t", strVersion, "\t", name)
			}
			continue
		}		
		fileSlice = append(fileSlice, FileInfo{FileName:names[0], FileVersion:strVersion})
		//fmt.Println(index, "\t", strVersion, "\t", names[0])
	}
	
	fileList := GetFileList()
	notSupport, notMatch, match := 0, 0, 0
	for _, file := range fileList {
		if strings.EqualFold(file, XLSX_FILENAME) || strings.EqualFold(file, RESLUT_FILENAME) {
			continue
		}
		
		version, err := GetFileVersion(file);
		if err!=nil {
			info := "Oppos\t["+file+"] file not support."
			fmt.Println(info)
			WriteResult(info)
			notSupport++
			continue
		}
		for _, fileInfo := range fileSlice {
			if !strings.EqualFold(fileInfo.FileName, file) {
				continue
			}
			if strings.EqualFold(fileInfo.FileVersion, version) {
				info := "Check\t["+file+"] version match."
				fmt.Println(info)
				WriteResult(info)
				match++
			} else {
				info := "Error\t["+file+"] version not match."
				fmt.Println(info)
				WriteResult(info)
				notMatch++
			}
		}		
	}	
	
	var strResult string = "\r\n结果：版本匹配【"+strconv.Itoa(match)+"】，版本差异【"+strconv.Itoa(notMatch)+"】，类型不支持【"+strconv.Itoa(notSupport)+"】"
	WriteResult(strResult)
}
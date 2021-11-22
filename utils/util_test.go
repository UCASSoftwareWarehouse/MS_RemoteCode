package utils

import (
	"fmt"
	"os"
	"testing"
)

func TestGetUUID(t *testing.T) {
	fmt.Println(GetUUID())
}

func TestCompress(t *testing.T) {
	pwd, _ := os.Getwd()
	pwd = GetParentDirectory(pwd)
	dirName := pwd + "/data"
	filePath := fmt.Sprintf("%s/%s.zip", dirName, "numpy")
	println(dirName)
	println(filePath)
	file, _ := os.Open(dirName + "/numpy")
	defer file.Close()
	var files = []*os.File{file}
	Compress(files, filePath)
}

func TestZip(t *testing.T) {
	pwd, _ := os.Getwd()
	pwd = GetParentDirectory(pwd)
	dirName := pwd + "/internal/data"
	filePath := fmt.Sprintf("%s/%s.zip", dirName, "numpy2")
	println(dirName)
	println(filePath)
	//Zip(dirName+"/numpy/test.txt",filePath)
	Zip(dirName+"/numpy", filePath)
}

func TestJudgeFileType(t *testing.T) {
	pwd, _ := os.Getwd()
	pwd = GetParentDirectory(pwd)
	//dirName:=pwd+"/internal/data/numpy"
	dirName := pwd + "/internal/data/b16db197c7e14bb38f8c0b1c66e2523a"

	println(dirName)
	fileName := JudgeSingleFileType(dirName)
	println(fileName)
}

package utils

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestGetUUID(t *testing.T) {
	fmt.Println(GetUUID())
}

//func TestCompress(t *testing.T) {
//	pwd, _ := os.Getwd()
//	pwd = GetParentDirectory(pwd)
//	dirName := pwd + "/data"
//	filePath := fmt.Sprintf("%s/%s.zip", dirName, "numpy")
//	println(dirName)
//	println(filePath)
//	file, _ := os.Open(dirName + "/numpy")
//	defer file.Close()
//	var files = []*os.File{file}
//	//Compress(files, filePath)
//}

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

func TestUntargz(t *testing.T) {
	srcPath := "/Users/zhujianxing/GoLandProjects/code/MS_RemoteCode/internal/data/panda-0.3.1.tar.gz"
	destPath := "/Users/zhujianxing/GoLandProjects/code/MS_RemoteCode/internal/data/"
	err := Untargz(srcPath, destPath)
	log.Printf("err=%+v", err)
}

func TestUnzip(t *testing.T) {
	srcPath := "/Users/zhujianxing/GoLandProjects/code/MS_RemoteCode/internal/data/numpy-1.19.3.zip"
	destPath := "/Users/zhujianxing/GoLandProjects/code/MS_RemoteCode/internal/data/"
	err := Unzip(srcPath, destPath)
	log.Printf("err=%+v", err)
}

func TestUnzip2(t *testing.T) {
	//.whl to .zip
	srcPath := "/Users/zhujianxing/GoLandProjects/code/MS_RemoteCode/internal/data/numpy-1.21.4-cp39-cp39-macosx_11_0_arm64.zip"
	destPath := "/Users/zhujianxing/GoLandProjects/code/MS_RemoteCode/internal/data/numpywheel/"
	err := Unzip(srcPath, destPath)
	log.Printf("err=%+v", err)
}

func TestUnWhell(t *testing.T) {
	srcPath := "/Users/zhujianxing/GoLandProjects/code/MS_RemoteCode/internal/data/numpy-1.21.4-cp39-cp39-macosx_11_0_arm64.whl"
	destPath := "/Users/zhujianxing/GoLandProjects/code/MS_RemoteCode/internal/data/numpywheel/"
	err := UnWheel(srcPath, destPath)
	log.Printf("err=%+v", err)
}

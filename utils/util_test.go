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

func TestZip2(t *testing.T) {
	srcPath := "/Users/zhujianxing/GoLandProjects/code/MS_RemoteCode/internal/data/ade98188483742d7896aeae430d72341"
	destPath := "/Users/zhujianxing/GoLandProjects/code/MS_RemoteCode/internal/data/ade98188483742d7896aeae430d72341/test.zip"
	Zip(srcPath, destPath)
}

func TestJudgeFileType(t *testing.T) {
	pwd, _ := os.Getwd()
	pwd = GetParentDirectory(pwd)
	//dirName:=pwd+"/internal/data/numpy"
	dirName := pwd + "/internal/data/pandas"

	println(dirName)
	fileName := JudgeSingleFileType(dirName)
	println(fileName)
}

func TestUntargz(t *testing.T) {
	srcPath := "/Users/zhujianxing/GoLandProjects/code/MS_RemoteCode/internal/data/numpy/panda-0.3.1.tar.gz"
	destPath := "/Users/zhujianxing/GoLandProjects/code/MS_RemoteCode/internal/data/numpy"
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

func TestGetSingleFileName(t *testing.T) {
	srcPath := "/Users/zhujianxing/GoLandProjects/code/MS_RemoteCode/internal/data/b16db197c7e14bb38f8c0b1c66e2523a"
	name := GetSingleFileName(srcPath)
	log.Println(name)
}

func TestWhl2Zip(t *testing.T) {
	srcPath := "/Users/zhujianxing/GoLandProjects/code/MS_RemoteCode/internal/data/numpy/numpy-1.21.4-cp39-cp39-macosx_11_0_arm64.whl"
	filePath, err := Whl2Zip(srcPath)
	log.Println(filePath)
	log.Println(err)
}

func TestTarGz2Zip(t *testing.T) {
	srcPath := "/Users/zhujianxing/GoLandProjects/code/MS_RemoteCode/internal/data/pandas/panda-0.3.1.tar.gz"
	dstPath := "/Users/zhujianxing/GoLandProjects/code/MS_RemoteCode/internal/data/pandas/panda-0.3.1.zip"
	TarGz2Zip(srcPath, dstPath)
}

func TestGetAllFile(t *testing.T) {
	srcPath := "/Users/zhujianxing/GoLandProjects/code/MS_RemoteCode/internal/data/ba9d95105e734442a0195162dadfe211"
	files, err := GetAllFile(srcPath)
	log.Println(err)
	for _, file := range files {
		log.Println(file)
	}
}

func TestGetFileName(t *testing.T) {
	srcPath := "/Users/zhujianxing/GoLandProjects/code/MS_RemoteCode/internal/data/pandas/panda-0.3.1.tar.gz"
	name := GetFileName(srcPath)
	log.Println(name)
}

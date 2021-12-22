package remote

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"remote_code/config"
	"remote_code/model"
	"remote_code/pb_gen"
	"testing"
)

func TestDownloadRemoteCode(t *testing.T) {
	//cmd:=exec.Command("/bin/bash/","pip3 download numpy -d ./numpy")
	//cmd:=exec.Command("python3","-V")
	cmd := exec.Command("bash", "-c", "pip3 download numpy -d ../data/numpy -i https://pypi.tuna.tsinghua.edu.cn/simple")
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	// 保证关闭输出流
	defer stdout.Close()
	// 运行命令
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	// 读取输出结果
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("stdout=", string(opBytes))
	opBytes, err = ioutil.ReadAll(stderr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("stderr=", string(opBytes))
}

//test zip
func TestDownloadRemoteCode2(t *testing.T) {
	config.InitConfigDefault()
	model.InitGorm()
	request := &pb_gen.DownloadRemoteCodeRequest{
		Metadata: &pb_gen.UploadMetadata2{
			ProjectId: 1,
			UserId:    1,
			FileInfo: &pb_gen.FileInfo2{
				FileName: "",
				FileType: 0,
			},
		},
		Platform: "win_amd64",
		NoDeps:   true,
		//OnlyBinary:    ":all:",
		PythonVersion: "",
		Package:       "numpy",
		Version:       "1.19.1",
	}
	code, err := DownloadRemoteCode(context.Background(), request)
	//os.RemoveAll("./e867c42b-6b78-4090-a656-72dc0cfd88f4")
	fmt.Printf("%+v", code)
	fmt.Printf("%+v", err)
}

//whl test
func TestDownloadRemoteCode3(t *testing.T) {
	os.Setenv("ENV", "prd")
	os.Setenv("CONFIG_PATH", "/home/zhujianxing/saas/MS_RemoteCode/config.yml")
	os.Setenv("NETWORK_INTERFACE", "docker_gwbridge")
	config.InitConfig()
	model.InitGorm()
	request := &pb_gen.DownloadRemoteCodeRequest{
		Metadata: &pb_gen.UploadMetadata2{
			ProjectId: 6,
			UserId:    1,
			FileInfo: &pb_gen.FileInfo2{
				FileName: "",
				FileType: 1,
			},
		},
		Platform: "",
		NoDeps:   false,
		//OnlyBinary:    ":all:",
		PythonVersion: "",
		Package:       "numpy",
		Version:       "",
	}
	code, err := DownloadRemoteCode(context.Background(), request)
	//os.RemoveAll("./e867c42b-6b78-4090-a656-72dc0cfd88f4")
	fmt.Printf("%+v", code)
	fmt.Printf("%+v", err)
}

func TestDownloadRemoteCode5(t *testing.T) {
	os.Setenv("ENV", "prd")
	os.Setenv("CONFIG_PATH", "/home/zhujianxing/saas/MS_RemoteCode/config.yml")
	os.Setenv("NETWORK_INTERFACE", "docker_gwbridge")
	config.InitConfig()
	model.InitGorm()
	request := &pb_gen.DownloadRemoteCodeRequest{
		Metadata: &pb_gen.UploadMetadata2{
			ProjectId: 6,
			UserId:    1,
			FileInfo: &pb_gen.FileInfo2{
				FileName: "",
				FileType: 1,
			},
		},
		Platform: "",
		NoDeps:   true,
		//OnlyBinary:    ":all:",
		PythonVersion: "",
		Package:       "request2",
		Version:       "",
	}
	code, err := DownloadRemoteCode(context.Background(), request)
	//os.RemoveAll("./e867c42b-6b78-4090-a656-72dc0cfd88f4")
	fmt.Printf("%+v", code)
	fmt.Printf("%+v", err)
}

//tag.gz test
func TestDownloadRemoteCode4(t *testing.T) {
	config.InitConfigDefault()
	model.InitGorm()
	request := &pb_gen.DownloadRemoteCodeRequest{
		Metadata: &pb_gen.UploadMetadata2{
			ProjectId: 1,
			UserId:    1,
			FileInfo: &pb_gen.FileInfo2{
				FileName: "",
				FileType: 0,
			},
		},
		Platform: "",
		NoDeps:   false,
		//OnlyBinary:    ":all:",
		PythonVersion: "",
		Package:       "panda",
		Version:       "",
	}
	code, err := DownloadRemoteCode(context.Background(), request)
	fmt.Printf("%+v", code)
	fmt.Printf("%+v", err)
}

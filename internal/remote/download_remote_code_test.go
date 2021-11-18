package remote

import (
	"io/ioutil"
	"log"
	"os/exec"
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

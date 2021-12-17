package utils

import (
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func CommandBash(command string) (string, string, error) {
	cmd := exec.Command("bash", "-c", command)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
		return "", "", err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
		return "", "", err
	}
	// 保证关闭输出流
	defer stdout.Close()
	defer stderr.Close()
	// 运行命令
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
		return "", "", err
	}
	// 读取输出结果
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
		return "", "", err
	}
	opBytes2, err := ioutil.ReadAll(stderr)
	if err != nil {
		log.Fatal(err)
		return "", "", err
	}
	errStr := handleStderr(string(opBytes2))
	return string(opBytes), errStr, nil
}

func handleStderr(stderr string) string {
	if strings.Contains(stderr, "ERROR:") || strings.Contains(stderr, "E:") {
		return stderr
	}
	return ""
}

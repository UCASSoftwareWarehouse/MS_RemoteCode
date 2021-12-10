package remote

import (
	"context"
	"fmt"
	"remote_code/config"
	"remote_code/model"
	"remote_code/pb_gen"
	"testing"
)

func TestDownloadAptDeb(t *testing.T) {
	config.InitConfigDefault()
	model.InitGorm()
	request := &pb_gen.DownloadAptDebRequest{
		UserId:  "1",
		Package: "jq",
		Version: "",
		Type:    "deb",
	}
	resp, err := DownloadAptDeb(context.Background(), request)
	fmt.Printf("%+v", resp)
	fmt.Printf("%+v", err)
}

func TestDownloadAptDeb2(t *testing.T) {
	config.InitConfigDefault()
	model.InitGorm()
	request := &pb_gen.DownloadAptDebRequest{
		UserId:  "1",
		Package: "jq",
		Version: "1.5",
		Type:    "deb",
	}
	resp, err := DownloadAptDeb(context.Background(), request)
	fmt.Printf("%+v", resp)
	fmt.Printf("%+v", err)
}

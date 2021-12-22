package remote

import (
	"context"
	"fmt"
	"os"
	"remote_code/config"
	"remote_code/model"
	"remote_code/pb_gen"
	"testing"
)

func TestDownloadAptDeb(t *testing.T) {
	config.InitConfigDefault()
	model.InitGorm()
	request := &pb_gen.DownloadAptDebRequest{
		Metadata: &pb_gen.UploadMetadata2{
			ProjectId: 1,
			UserId:    1,
			FileInfo: &pb_gen.FileInfo2{
				FileName: "",
				FileType: 2,
			},
		},
		Package: "jq",
		Version: "",
		Type:    "deb",
	}
	resp, err := DownloadAptDeb(context.Background(), request)
	fmt.Printf("%+v", resp)
	fmt.Printf("%+v", err)
}

func TestDownloadAptDeb2(t *testing.T) {
	// config.InitConfigDefault()
	os.Setenv("ENV", "prd")
	os.Setenv("CONFIG_PATH", "/home/zhujianxing/saas/MS_RemoteCode/config.yml")
	os.Setenv("NETWORK_INTERFACE", "docker_gwbridge")
	config.InitConfig()
	model.InitGorm()
	request := &pb_gen.DownloadAptDebRequest{
		Metadata: &pb_gen.UploadMetadata2{
			ProjectId: 5,
			UserId:    1,
			FileInfo: &pb_gen.FileInfo2{
				FileName: "",
				FileType: 2,
			},
		},
		Package: "jq",
		Version: "1.5",
		Type:    "deb",
	}
	resp, err := DownloadAptDeb(context.Background(), request)
	fmt.Printf("%+v", resp)
	fmt.Printf("%+v", err)
}

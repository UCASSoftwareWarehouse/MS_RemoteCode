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
	}
	resp, err := DownloadAptDeb(context.Background(), request)
	fmt.Printf("%+v", resp)
	fmt.Printf("%+v", err)
}

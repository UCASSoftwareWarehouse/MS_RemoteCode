package upload

import (
	pb_gen2 "MS_Local/pb_gen"
	"context"
	"log"
	"remote_code/config"
	"testing"
)

func TestUpload(t *testing.T) {
	config.InitConfigDefault()
	fpath := "../data/apt/jq_1.5+dfsg-2_amd64.deb"
	resp, err := Upload(context.Background(), 1, 1, fpath, pb_gen2.FileType_binary)
	log.Println(err)
	log.Println(resp)
}

package upload

import (
	pb_gen2 "MS_Local/pb_gen"
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"remote_code/consul"
	"time"
)

func Upload(ctx context.Context, uid uint64, pid uint64, fpath string, fileType pb_gen2.FileType) (*pb_gen2.UploadResponse, error) {

	file, err := os.Open(fpath)
	if err != nil {
		log.Printf("cannot open file: %+v", err)
		return nil, err
	}
	defer file.Close()

	client := &consul.GrpcClient{Name: "sw_ms_local"}
	client.RunConsulClient()
	local_code := pb_gen2.NewMSLocalClient(client.Conn)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	uploadClient, err := local_code.Upload(ctx)
	if err != nil {
		log.Printf("cannot upload file: %+v", err)
		return nil, err
	}

	finfo, err := os.Stat(fpath)
	if err != nil {
		log.Printf("cannot get file info %+v", err)
		return nil, err
	}

	request := &pb_gen2.UploadRequest{
		Data: &pb_gen2.UploadRequest_Metadata{
			Metadata: &pb_gen2.UploadMetadata{
				ProjectId: pid,
				UserId:    uid,
				FileInfo: &pb_gen2.FileInfo{
					FileName: finfo.Name(),
					FileType: fileType,
				},
			},
		},
	}
	log.Printf("request is:\n%v", request)

	err = uploadClient.Send(request)
	if err != nil {
		log.Printf("cannot send image info to server: %+v,%+v", err, uploadClient.RecvMsg(nil))
		return nil, err
	}
	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("cannot read chunk to buffer: %+v", err)
			return nil, err
		}
		req := &pb_gen2.UploadRequest{
			Data: &pb_gen2.UploadRequest_Content{
				Content: buffer[:n],
			},
		}

		err = uploadClient.Send(req)
		if err != nil {
			log.Printf("cannot send chunk to server: %+v,%+v", err, uploadClient.RecvMsg(nil))
			return nil, err
		}
	}
	resp, err := uploadClient.CloseAndRecv()
	if err != nil {
		log.Printf("cannot receive response: %+v", err)
		return nil, err
	}
	return resp, nil
}

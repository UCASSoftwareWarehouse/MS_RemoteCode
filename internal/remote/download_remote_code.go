package remote

import (
	"context"
	"log"
	"remote_code/constant"
	"remote_code/pb_gen"
	"remote_code/utils"
	"strings"
)

func DownloadRemoteCode(ctx context.Context, req *pb_gen.DownloadRemoteCodeRequest) (*pb_gen.DownloadRemoteCodeResponse, error) {

	var args []string
	resp := &pb_gen.DownloadRemoteCodeResponse{}

	if len(req.Platform) != 0 {
		args = append(args, "--platform "+req.Platform)
	}
	if req.NoDeps {
		args = append(args, "--no-deps ")
	}
	if len(req.OnlyBinary) != 0 {
		args = append(args, "--only-binary "+req.OnlyBinary)
	}
	if len(req.PythonVersion) != 0 {
		args = append(args, "--python-version "+req.PythonVersion)
	}
	if len(req.Package) == 0 {
		log.Fatal("package can not be empty")
		resp.Message = "package can not be empty"
		resp.Code = constant.STATUS_OK
		return resp, nil
	}
	command := "pip3 download " + req.Package + strings.Join(args, " ")
	log.Printf("command from %+v:%+v", req.UserId, command)
	stdout, stderr, err := utils.CommandBash(command)
	if err != nil {
		log.Fatal("DownloadRemoteCode CommandBash err:%+v", err)
		resp.Message = err.Error()
		resp.Code = constant.STATUS_BADREQUEST
		return resp, err
	}
	if len(stderr) != 0 {
		log.Fatal("DownloadRemoteCode CommandBash err:%+v", err)
		resp.Message = stderr
		resp.Code = constant.STATUS_BADREQUEST
		return resp, nil
	}
	log.Printf("stdout:%+v", stdout)
	resp.Message = constant.MESSAGE_OK
	resp.Code = constant.STATUS_OK
	return resp, nil
}

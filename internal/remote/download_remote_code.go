package remote

import (
	"context"
	"fmt"
	"log"
	"os"
	"remote_code/config"
	"remote_code/constant"
	"remote_code/model"
	"remote_code/pb_gen"
	"remote_code/utils"
	"strings"
)

func DownloadRemoteCode(ctx context.Context, req *pb_gen.DownloadRemoteCodeRequest) (*pb_gen.DownloadRemoteCodeResponse, error) {
	var args []string
	resp := &pb_gen.DownloadRemoteCodeResponse{}

	//校验userid todo 鉴权？
	user := &model.User{}
	userId := req.Metadata.UserId
	_, err := user.FindUserById(ctx, userId)
	if err != nil {
		resp.Message = "user_id doesn't exit"
		resp.Code = constant.STATUS_BADREQUEST
		log.Printf("user_id doesn't exit")
		return resp, err
	}

	//参数判断
	if len(req.Package) == 0 {
		log.Printf("package can not be empty")
		resp.Message = "package can not be empty"
		resp.Code = constant.STATUS_OK
		return resp, nil
	}
	if len(req.Platform) != 0 {
		//args = append(args, "--platform "+req.Platform+" --no-deps")
		args = append(args, "--platform "+req.Platform)
	}
	if req.NoDeps {
		args = append(args, "--no-deps")
	}
	if len(req.OnlyBinary) != 0 {
		args = append(args, "--only-binary="+req.OnlyBinary)
	}
	if len(req.PythonVersion) != 0 {
		args = append(args, "--python-version "+req.PythonVersion)
	}
	fileName := req.Package
	if len(req.Version) != 0 {
		fileName = fmt.Sprintf("%s==%s ", req.Package, req.Version)
	}

	//执行pip download命令
	pwd, _ := os.Getwd()
	for !strings.HasSuffix(pwd, "/MS_RemoteCode") {
		pwd = utils.GetParentDirectory(pwd)
	}
	log.Println(pwd)
	dirName := pwd + "/internal/data/" + utils.GetUUID()
	args = append(args, " -d "+dirName)
	var command string
	if config.IsProd() {
		command = "python3 -m pip download " + fileName + strings.Join(args, " ")
	} else {
		command = "pip3 download " + fileName + strings.Join(args, " ")
	}
	log.Printf("command from user %+v:%+v", req.Metadata.UserId, command)
	stdout, stderr, err := utils.CommandBash(command)
	if err != nil {
		log.Printf("DownloadRemoteCode CommandBash err:%+v", err)
		resp.Message = err.Error()
		resp.Code = constant.STATUS_BADREQUEST
		return resp, err
	}
	if len(stderr) != 0 {
		log.Printf("DownloadRemoteCode CommandBash err:%+v", err)
		resp.Message = stderr
		resp.Code = constant.STATUS_BADREQUEST
		return resp, nil
	}
	log.Printf("stdout:%+v", stdout)

	/*
		压缩
		pypi包下载位置 "../MS_RemoteCode/internal/data/bb5e44e4febb4fcc88ea6824db4f9689"
		打zip包位置 "../MS_RemoteCode/internal/data/numpy.zip"
	*/
	filePath := utils.GetParentDirectory(dirName)
	filePath = fmt.Sprintf("%s/%s.zip", filePath, req.Package)
	log.Println(dirName)
	log.Println(filePath)
	// todo judge fileType 统一为zip
	utils.Zip(dirName, filePath)
	// todo 上传接口 接入guohao mongodb_code
	//response, err := upload.Upload(ctx, req.Metadata.UserId, req.Metadata.ProjectId, filePath, pb_gen2.FileType(req.Metadata.FileInfo.FileType))
	//project:=response.ProjectInfo
	//resp.ProjectInfo=&pb_gen.Project{
	//	Id:                 project.Id,
	//	ProjectName:        project.ProjectName,
	//	UserId:             project.UserId,
	//	Tags:               project.Tags,
	//	License:            project.License,
	//	Updatetime:         project.Updatetime,
	//	ProjectDescription: project.ProjectDescription,
	//	CodeAddr:           project.CodeAddr,
	//	BinaryAddr:         project.BinaryAddr,
	//	Classifiers:        project.Classifiers,
	//}
	//删除文件
	//defer os.RemoveAll(dirName)

	resp.Message = constant.MESSAGE_OK
	resp.Code = constant.STATUS_OK
	return resp, nil
}

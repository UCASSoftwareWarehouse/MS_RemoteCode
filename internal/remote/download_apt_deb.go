package remote

import (
	pb_gen2 "MS_Local/pb_gen"
	"context"
	"fmt"
	"log"
	"os"
	"remote_code/constant"
	upload "remote_code/internal/upload_service"
	"remote_code/model"
	"remote_code/pb_gen"
	"remote_code/utils"
	"strings"
)

func DownloadAptDeb(ctx context.Context, req *pb_gen.DownloadAptDebRequest) (*pb_gen.DownloadAptDebResponse, error) {

	if req.Type == "deb" {
		return DownloadAptDebPkg(ctx, req)
	} else if req.Type == "code" {
		//return DownloadAptDebCode(ctx,req)
	}

	resp := &pb_gen.DownloadAptDebResponse{}
	resp.Code = constant.STATUS_BADREQUEST
	resp.Message = "invalid type"
	return resp, nil
}

func DownloadAptDebPkg(ctx context.Context, req *pb_gen.DownloadAptDebRequest) (*pb_gen.DownloadAptDebResponse, error) {
	resp := &pb_gen.DownloadAptDebResponse{}

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

	fileName := req.Package
	if len(req.Version) != 0 {
		fileName = fmt.Sprintf("%s=%s\\*", req.Package, req.Version)
	}

	//执行pip download命令
	pwd, _ := os.Getwd()
	for !strings.HasSuffix(pwd, "/MS_RemoteCode") {
		pwd = utils.GetParentDirectory(pwd)
	}
	log.Println(pwd)
	uuid := utils.GetUUID()
	//deb文件download地址： ../data/{uuid}/
	dirName := pwd + "/internal/data/" + uuid
	log.Printf("download path:%+v", dirName)
	utils.CommandBash("mkdir ../data/" + uuid)
	command := "apt-get download " + fileName
	log.Printf("command from %+v:%+v", req.Metadata.UserId, command)
	stdout, stderr, err := utils.CommandBash(command)
	if err != nil {
		log.Printf("DownloadAptDeb CommandBash err:%+v", err)
		resp.Message = err.Error()
		resp.Code = constant.STATUS_BADREQUEST
		return resp, err
	}
	if len(stderr) != 0 {
		log.Printf("DownloadAptDeb CommandBash err:%+v", stderr)
		resp.Message = stderr
		resp.Code = constant.STATUS_BADREQUEST
		return resp, nil
	}
	os.Mkdir(dirName, os.ModePerm)
	utils.CommandBash("mv ./*.deb " + dirName + "/")
	log.Printf("stdout:%+v", stdout)

	/*
		压缩
		deb包下载位置 "../MS_RemoteCode/internal/data/bb5e44e4febb4fcc88ea6824db4f9689"
		打zip包位置 "../MS_RemoteCode/internal/data/numpy.zip"
	*/
	debPath := fmt.Sprintf("%s/%s", dirName, utils.GetSingleFileName(dirName))
	log.Printf("debPath=%+v", debPath)
	response, err := upload.Upload(ctx, req.Metadata.UserId, req.Metadata.ProjectId, debPath, pb_gen2.FileType(req.Metadata.FileInfo.FileType))
	log.Println(err)
	log.Println(response)
	if err != nil {
		log.Printf("DownloadAptDeb err:%+v", err.Error())
		resp.Message = "上传失败"
		resp.Code = constant.STATUS_BADREQUEST
		return resp, nil
	}
	project := response.ProjectInfo
	resp.ProjectInfo = &pb_gen.Project2{
		Id:                 project.Id,
		ProjectName:        project.ProjectName,
		UserId:             project.UserId,
		Tags:               project.Tags,
		License:            project.License,
		Updatetime:         project.Updatetime,
		ProjectDescription: project.ProjectDescription,
		CodeAddr:           project.CodeAddr,
		BinaryAddr:         project.BinaryAddr,
		Classifiers:        project.Classifiers,
	}

	// todo 删除文件
	// os.RemoveAll(dirName)

	resp.Message = constant.MESSAGE_OK
	resp.Code = constant.STATUS_OK
	return resp, nil
}

//func DownloadAptDebCode(ctx context.Context, req *pb_gen.DownloadAptDebRequest) (*pb_gen.DownloadAptDebResponse, error) {
//
//}

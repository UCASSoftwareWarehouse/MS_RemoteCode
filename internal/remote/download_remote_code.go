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
	initPath := pwd + "/internal/data/" + utils.GetUUID()
	dirName := initPath + "/tmp"
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
		log.Printf("DownloadRemoteCode CommandBash err:%+v", stderr)
		resp.Message = stderr
		resp.Code = constant.STATUS_BADREQUEST
		return resp, nil
	}
	log.Printf("stdout:%+v", stdout)

	/*
		压缩
		pypi包下载位置 "../MS_RemoteCode/internal/data/bb5e44e4febb4fcc88ea6824db4f9689/tmp"
		打zip包位置 "../MS_RemoteCode/internal/data/bb5e44e4febb4fcc88ea6824db4f9689/"
	*/
	// filePath为最终文件路径
	filePath, err := handleFile(dirName, req.Package)
	if err != nil {
		resp.Message = err.Error()
		resp.Code = constant.STATUS_BADREQUEST
		return resp, nil
	}

	log.Printf("final file path:%+v", filePath)

	// todo 上传接口 接入guohao
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
	//defer os.RemoveAll(initPath)

	resp.Message = constant.MESSAGE_OK
	resp.Code = constant.STATUS_OK
	return resp, nil
}

func handleFile(dirName string, fileName string) (string, error) {
	fileList, err := utils.GetAllFile(dirName)
	var sourceName string
	flag := false
	if len(fileList) > 1 {
		flag = true
	}
	//对于可能存在依赖的情况，只解压package
	for _, file := range fileList {
		if strings.Contains(file, fileName) {
			end := strings.LastIndex(file, "/")
			sourceName = file[end+1:]
			break
		}
	}
	filePath := ""
	//sourceName:=utils.GetSingleFileName(dirName)
	srcPath := fmt.Sprintf("%s/%s", dirName, sourceName)
	log.Printf("sourcePath=%+v", srcPath)
	fileType := utils.JudgeFileTypeByPath(srcPath)
	// 有依赖的多个文件
	if flag {
		log.Println(dirName)
		name := utils.GetFileName(sourceName)
		newDirName := strings.ReplaceAll(dirName, "tmp", name)
		os.Rename(dirName, newDirName)
		dstPath := fmt.Sprintf("%s/%s.zip", utils.GetParentDirectory(newDirName), name)
		utils.Zip(newDirName, dstPath)
		filePath = dstPath
	} else {
		if fileType == "whl" {
			filePath, err = utils.Whl2Zip(srcPath)
			if err != nil {
				log.Printf("DownloadRemoteCode whl2Zip err:%+v", err)
				return "", err
			}
		} else if fileType == "gz" {
			dstZipName := strings.ReplaceAll(sourceName, "tar.gz", "zip")
			filePath = fmt.Sprintf("%s/%s", dirName, dstZipName)
			_, err = utils.TarGz2Zip(srcPath, filePath)
		} else {
			filePath = fmt.Sprintf("%s/%s", dirName, sourceName)
			if err != nil {
				log.Printf("DownloadRemoteCode whl2Zip err:%+v", err)
				return "", err
			}
		}
	}
	return filePath, nil
}

package models

import (
	"github.com/weilaihui/fdfs_client"
	"log"
)

//根据文件路径上传文件
func FDFSUploadByFilename(filename string) (fileId string, err error) {
	fdfsClient, err := fdfs_client.NewFdfsClient("./conf/client.conf")
	if err != nil {
		log.Fatal("create fdfs_client err!!loadByFilename")
		return
	}
	uploadRespons, err := fdfsClient.UploadByFilename(filename)
	if err != nil {
		log.Fatal("upload file err")
		return
	}
	fileId = uploadRespons.RemoteFileId
	return
}

//上传内存的文件
func FDFSUploadByBuffer(buffer []byte, postfix string) (fileId string, err error) {
	fdfsClient, err := fdfs_client.NewFdfsClient("./conf/client.conf")
	if err != nil {
		log.Fatal("create fdfs_client err!!loadBybuffer")
		return
	}
	uploadRespons, err := fdfsClient.UploadByBuffer(buffer, postfix)
	if err != nil {
		log.Fatal("upload file err")
		return
	}
	fileId = uploadRespons.RemoteFileId
	return
}

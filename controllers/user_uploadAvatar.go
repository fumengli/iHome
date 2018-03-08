package controllers

//api/v1.0/userss
//LoginNewUser

import (
	_ "encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"iHome/models"
	"path"
)

type UploadAvatarController struct {
	beego.Controller
}

func (this *UploadAvatarController) RetrunData(retmsg interface{}) {
	this.Data["json"] = retmsg
	this.ServeJSON()
}
func (this *UploadAvatarController) UploadAvatar() {
	beego.Info("=========upload avatar succ!=============")
	var retmsg = make(map[string]interface{})
	retmsg["errno"] = models.RECODE_OK
	retmsg["errmsg"] = models.RecodeText(models.RECODE_OK)
	defer this.RetrunData(retmsg)
	//1 得到头像图片的二进制文件,上传到fdfs
	file, header, err := this.GetFile("avatar")
	if err != nil {
		fmt.Println("xxxxx", err)
		retmsg["errno"] = models.RECODE_IOERR
		retmsg["errmsg"] = models.RecodeText(models.RECODE_IOERR)
		return
	}
	file_postfix := path.Ext(header.Filename)
	filebuffer := make([]byte, header.Size)
	file.Read(filebuffer)
	fileID, err := models.FDFSUploadByBuffer(filebuffer, file_postfix[1:])
	if err != nil {
		fmt.Println("sssssssss", err)
		retmsg["errno"] = models.RECODE_IOERR
		retmsg["errmsg"] = models.RecodeText(models.RECODE_IOERR)
		return
	}
	//2 从session中得到当前用户
	user_id := this.GetSession("user_id")
	curUser := models.User{Id: user_id.(int), Avatar_url: fileID}
	//3 图片fileid 更新到数据库
	o := orm.NewOrm()
	o.Update(&curUser, "Avatar_url")
	//3.1头像更新到session
	this.SetSession("avatar_url", curUser.Avatar_url)
	//4 生成返回数据
	url, _ := models.GetUrl()
	retUrl := make(map[string]interface{})
	retUrl["avatar_url"] = "http://" + url + fileID
	retmsg["data"] = retUrl

	return
}

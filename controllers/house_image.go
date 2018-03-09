package controllers

import (
	//"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"iHome/models"
	"path"
	"strconv"
)

type HouseImageController struct {
	beego.Controller
}

//返回结构变成json返回前端
func (this *HouseImageController) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

//上传房屋图片 [post]
func (this *HouseImageController) UploadImage() {
	beego.Info("========upload house image succ....")

	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	defer this.RetData(resp)

	//得到文件二进制数据
	file, header, err := this.GetFile("house_image")
	if err != nil {
		resp["errno"] = models.RECODE_SERVERERR
		resp["errmsg"] = models.RecodeText(models.RECODE_SERVERERR)
		return
	}

	fileBuffer := make([]byte, header.Size)
	if _, err := file.Read(fileBuffer); err != nil {
		resp["errno"] = models.RECODE_IOERR
		resp["errmsg"] = models.RecodeText(models.RECODE_IOERR)
		return
	}

	suffix := path.Ext(header.Filename)

	//将文件的二进制数据上传到fastdfs中

	fileId, err := models.FDFSUploadByBuffer(fileBuffer, suffix[1:])
	if err != nil {
		resp["errno"] = models.RECODE_IOERR
		resp["errmsg"] = models.RecodeText(models.RECODE_IOERR)
		beego.Info("upload house image to fastdfs error err = ", err)
		return
	}

	beego.Info("fdfs upload suc  fileid = ", fileId)

	//获取房屋id
	str_hid := this.Ctx.Input.Param(":id")
	beego.Info("str_hid  =", str_hid)
	hid, _ := strconv.Atoi(str_hid)
	beego.Info("===id = ", hid)

	house := models.House{Id: hid}

	//数据库的操作，
	o := orm.NewOrm()
	//images := make([]*models.HouseImage, 0)

	err = o.Read(&house)
	if err != nil {
		fmt.Println("house", house)
		beego.Info("read house err...")
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	//图片信息存入结构体
	houseImage := models.HouseImage{Url: fileId, House: &house}
	//房屋图片插入数据库
	image_id, err := o.Insert(&houseImage)
	if err != nil {
		beego.Info("isnsert houseimage err...")
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	beego.Info("insert houseimage succ image id = ", image_id)

	//房屋图片数组
	house.Images = append(house.Images, &houseImage)

	if house.Index_image_url == "" {
		house.Index_image_url = fileId
	}

	if _, err := o.Update(&house, "Images", "Index_image_url"); err != nil {
		beego.Info("update house image   Index_image_url  err", err)
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	//将fileid拼接成一个完整的url路径
	m_url, _ := models.GetUrl()

	image_url := m_url + "/" + fileId

	beego.Info("===============image_url=", image_url)
	//安装协议做出json返回给前端

	url_map := make(map[string]interface{})
	url_map["url"] = "http://" + image_url
	resp["data"] = url_map
	return
}

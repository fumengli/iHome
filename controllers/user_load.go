package controllers

//api/v1.0/userss
//LoginNewUser

import (
	//	"encoding/json"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/orm"
	"iHome/models"
)

type LoadUserController struct {
	beego.Controller
}

func (this *LoadUserController) RetrunData(retmsg interface{}) {
	this.Data["json"] = retmsg
	this.ServeJSON()
}
func (this *LoadUserController) LoadUser() {
	beego.Info("=========Load user succ!=============")
	var retmsg = make(map[string]interface{})
	retmsg["errno"] = models.RECODE_OK
	retmsg["errmsg"] = models.RecodeText(models.RECODE_OK)
	defer this.RetrunData(retmsg)

	//2 直接从当前的session中拿数据
	curuser := models.User{}
	if tmp := this.GetSession("name"); tmp != nil {
		curuser.Name = this.GetSession("name").(string)
	} else {
		beego.Info("name is empty")
	}
	url, _ := models.GetUrl()
	curuser.Id = this.GetSession("user_id").(int)
	curuser.Mobile = this.GetSession("mobile").(string)
	if tmp := this.GetSession("avatar_url"); tmp == nil {
		curuser.Avatar_url = "http://" + url + "static/images/defaultAvatar.jpg"
		this.SetSession("avatar_url", curuser.Avatar_url)
	}
	curuser.Avatar_url = this.GetSession("avatar_url").(string)

	retmsg["data"] = curuser
	return
}

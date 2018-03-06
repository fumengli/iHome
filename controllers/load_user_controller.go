package controllers

//api/v1.0/userss
//LoginNewUser

import (
	"encoding/json"
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
	//1 得到客户端请求的json数据
	var regRequestMap = make(map[string]interface{})
	json.Unmarshal(this.Ctx.Input.RequestBody, &regRequestMap)

	//2 直接从当前的session中拿数据
	curuser := models.User{}

	curuser.Name = this.GetSession("name").(string)
	curuser.Id = this.GetSession("user_id").(int)
	curuser.Mobile = this.GetSession("mobile").(string)
	retmsg["data"] = curuser
	return
}

package controllers

//api/v1.0/userss
//LoginNewUser

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"iHome/models"
)

type ModifyUserNameController struct {
	beego.Controller
}

func (this *ModifyUserNameController) RetrunData(retmsg interface{}) {
	this.Data["json"] = retmsg
	this.ServeJSON()
}
func (this *ModifyUserNameController) ModifyUserName() {
	beego.Info("=========Modify user name succ!=============")
	var retmsg = make(map[string]interface{})
	retmsg["errno"] = models.RECODE_OK
	retmsg["errmsg"] = models.RecodeText(models.RECODE_OK)
	defer this.RetrunData(retmsg)
	//1 得到客户端请求的json数据
	var regRequestMap = make(map[string]interface{})
	json.Unmarshal(this.Ctx.Input.RequestBody, &regRequestMap)
	beego.Info("get Name:", regRequestMap["name"])
	this.SetSession("name", regRequestMap["name"])
	//2 直接从当前的session中拿数据
	curuser := models.User{}
	curuser.Id = this.GetSession("user_id").(int)
	curuser.Name = this.GetSession("name").(string)
	//修改到数据库
	o := orm.NewOrm()
	if _, err := o.Update(&curuser, "Name"); err != nil {
		retmsg["errno"] = models.RECODE_DBERR
		retmsg["errmsg"] = models.RecodeText(models.RECODE_DBERR)

	}
	return
}

package controllers

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/orm"
	"iHome/models"
)

type SessionController struct {
	beego.Controller
}

func (this *SessionController) RetrunData(retmsg interface{}) {
	this.Data["json"] = retmsg
	this.ServeJSON()
}

func (this *SessionController) GetSessionName() {
	beego.Info("=========GetSession succ!=============")
	var retmsg = make(map[string]interface{})
	retmsg["errno"] = models.RECODE_SESSIONERR
	retmsg["errmsg"] = models.RecodeText(models.RECODE_SESSIONERR)
	defer this.RetrunData(retmsg)
	nameMap := make(map[string]interface{})
	name := this.GetSession("name")
	if name != nil {
		retmsg["errno"] = models.RECODE_OK
		retmsg["errmsg"] = models.RecodeText(models.RECODE_OK)
		nameMap["name"] = name
		retmsg["data"] = nameMap
	}
	return
}

func (this *SessionController) DelSessionName() {
	beego.Info("=========DelSession succ!=============")
	var retmsg = make(map[string]interface{})
	retmsg["errno"] = models.RECODE_OK
	retmsg["errmsg"] = models.RecodeText(models.RECODE_OK)
	defer this.RetrunData(retmsg)
	this.DelSession("name")
	this.DelSession("user_id")
	this.DelSession("mobile")
	this.DelSession("Password_hash")
	this.DelSession("Real_name")
	this.DelSession("Id_card")
	this.DelSession("avatar_url")

	return

}

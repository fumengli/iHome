package controllers

//api/v1.0/userss
//RegNewUser

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"iHome/models"
)

type RegNewUserController struct {
	beego.Controller
}

func (this *RegNewUserController) RetrunData(retmsg interface{}) {
	this.Data["json"] = retmsg
	this.ServeJSON()
}
func (this *RegNewUserController) RegNewUser() {
	beego.Info("=========GetRegInfo succ!=============")
	var retmsg = make(map[string]interface{})
	retmsg["errno"] = models.RECODE_SESSIONERR
	retmsg["errmsg"] = models.RecodeText(models.RECODE_SESSIONERR)
	defer this.RetrunData(retmsg)
	//1 得到客户端请求的json数据
	var regRequestMap = make(map[string]interface{})
	json.Unmarshal(this.Ctx.Input.RequestBody, &regRequestMap)

	//beego.Info("mobile= ", regRequestMap["mobile"])
	//beego.Info("password=", regRequestMap["password"])
	//beego.Info("sms_code=", regRequestMap["sms_code"])
	//2 判断用户数据合法性
	if regRequestMap["mobile"] == "" || regRequestMap["password"] == "" || regRequestMap["sms_code"] == nil {
		retmsg["errno"] = models.RECODE_PARAMERR
		retmsg["errmsg"] = models.RecodeText(models.RECODE_PARAMERR)
		return
	}
	//3 将数据存入mysql
	md5Ctx := md5.New()
	user := models.User{}
	user.Mobile = regRequestMap["mobile"].(string)
	md5Ctx.Write([]byte(regRequestMap["password"].(string)))
	user.Password_hash = hex.EncodeToString(md5Ctx.Sum(nil))
	user.Name = regRequestMap["mobile"].(string)
	o := orm.NewOrm()
	id, err := o.Insert(&user)
	if err != nil {
		beego.Info("insert error = ", err)
		retmsg["errno"] = models.RECODE_DBERR
		retmsg["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	//4 将当前的用户的信息存储到session中
	this.SetSession("name", user.Mobile)
	this.SetSession("user_id", id)
	this.SetSession("mobile", user.Mobile)
	return
}

package controllers

//api/v1.0/userss
//LoginNewUser

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"iHome/models"
)

type LoginUserController struct {
	beego.Controller
}

func (this *LoginUserController) RetrunData(retmsg interface{}) {
	this.Data["json"] = retmsg
	this.ServeJSON()
}
func (this *LoginUserController) LoginUser() {
	beego.Info("=========GetLoginInfo succ!=============")
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
	if regRequestMap["mobile"] == "" || regRequestMap["password"] == "" {
		retmsg["errno"] = models.RECODE_PARAMERR
		retmsg["errmsg"] = models.RecodeText(models.RECODE_PARAMERR)
		return
	}
	//3 从mysql中查询数据
	user := models.User{Mobile: regRequestMap["mobile"].(string)}

	o := orm.NewOrm()
	err := o.Read(&user, "Mobile")
	if err != nil {
		beego.Info("Read error = ", err)
		retmsg["errno"] = models.RECODE_NODATA
		retmsg["errmsg"] = models.RecodeText(models.RECODE_NODATA)
		return
	} //有该用户
	beego.Info("==========read succ===========", user.Name)
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(regRequestMap["password"].(string)))
	input_hash := hex.EncodeToString(md5Ctx.Sum(nil))
	if user.Password_hash == input_hash { // password ok
		retmsg["errno"] = models.RECODE_OK
		retmsg["errmsg"] = models.RecodeText(models.RECODE_PWDERR)
	} else {
		beego.Info("===========password err=============")
		beego.Info("db_hash", user.Password_hash)
		beego.Info("input_hash", input_hash)
		retmsg["errno"] = models.RECODE_PWDERR
		retmsg["errmsg"] = models.RecodeText(models.RECODE_PWDERR)
		return
	}
	//4 将当前的用户的信息存储到session中
	this.SetSession("name", user.Name)
	this.SetSession("user_id", user.Id)
	this.SetSession("mobile", user.Mobile)
	this.SetSession("avatar_url", user.Avatar_url)
	this.SetSession("real_name", user.Real_name)
	this.SetSession("id_card", user.Id_card)
	this.SetSession("password", user.Password_hash)
	return
}

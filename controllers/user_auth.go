package controllers

//api/v1.0/userss
//LoginNewUser

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"iHome/models"
)

type AuthUserController struct {
	beego.Controller
}

func (this *AuthUserController) RetrunData(retmsg interface{}) {
	this.Data["json"] = retmsg
	this.ServeJSON()
}

func (this *AuthUserController) AuthUser() {
	beego.Info("=========AuthUser GET succ!=============")
	var retmsg = make(map[string]interface{})
	retmsg["errno"] = models.RECODE_OK
	retmsg["errmsg"] = models.RecodeText(models.RECODE_OK)
	defer this.RetrunData(retmsg)

	//从当前session中得到数据
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
		beego.Info("USE default!url=", curuser.Avatar_url)
		this.SetSession("avatar_url", curuser.Avatar_url)
	}
	curuser.Avatar_url = this.GetSession("avatar_url").(string)
	curuser.Id_card = this.GetSession("id_card").(string)
	curuser.Real_name = this.GetSession("real_name").(string)
	curuser.Password_hash = this.GetSession("password").(string)
	retmsg["data"] = curuser

	return
}

//post
func (this *AuthUserController) AuthUserPOST() {
	beego.Info("=========AuthUserPOST succ!=============")
	var retmsg = make(map[string]interface{})
	retmsg["errno"] = models.RECODE_SESSIONERR
	retmsg["errmsg"] = models.RecodeText(models.RECODE_SESSIONERR)
	defer this.RetrunData(retmsg)
	//1 得到客户端请求的json数据
	var regRequestMap = make(map[string]interface{})
	json.Unmarshal(this.Ctx.Input.RequestBody, &regRequestMap)

	beego.Info("id_card= ", regRequestMap["id_card"])
	beego.Info("real_name=", regRequestMap["real_name"])
	//2 判断用户数据合法性
	if regRequestMap["id_card"] == "" || regRequestMap["real_name"] == "" {
		retmsg["errno"] = models.RECODE_PARAMERR
		retmsg["errmsg"] = models.RecodeText(models.RECODE_PARAMERR)
		return
	}
	//3 写入数据库
	mobile := this.GetSession("mobile").(string)
	user := models.User{Mobile: mobile}
	o := orm.NewOrm()
	if o.Read(&user, "mobile") == nil {

		user.Id_card = regRequestMap["id_card"].(string)
		user.Real_name = regRequestMap["real_name"].(string)
		_, err := o.Update(&user)

		if err != nil {
			beego.Info("Write error = ", err)
			retmsg["errno"] = models.RECODE_NODATA
			retmsg["errmsg"] = models.RecodeText(models.RECODE_NODATA)
			return
		}
	}
	beego.Info("update succ!===")
	//4 将新增的信息存储到session中
	this.SetSession("id_card", user.Id_card)
	this.SetSession("real_name", user.Real_name)
	//5 返回信息
	retmsg["errno"] = models.RECODE_OK
	retmsg["errmsg"] = models.RecodeText(models.RECODE_OK)
	retmsg["data"] = user
	return
}

package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"iHome/models"
)

type HouseLoadController struct {
	beego.Controller
}

func (this *HouseLoadController) RetrunData(retmsg interface{}) {
	this.Data["json"] = retmsg
	this.ServeJSON()
}

func (this *HouseLoadController) GetHouseLoad() {
	beego.Info("=========GetHouseLoad succ!=============")
	var retmsg = make(map[string]interface{})
	retmsg["errno"] = models.RECODE_SESSIONERR
	retmsg["errmsg"] = models.RecodeText(models.RECODE_SESSIONERR)
	defer this.RetrunData(retmsg)
	//1 通过session 拿到用户的user_id
	userId := this.GetSession("user_id")
	if userId == nil {
		beego.Info("userid session is empty")
		return
	}
	//根据id 查询该用户信息
	o := orm.NewOrm()
	var houses []models.House = []models.House{}
	num, err := o.QueryTable("house").Filter("user", userId.(int)).All(&houses)
	if err == nil {
		fmt.Printf("%d houses read\n", num)
	}
	//组织返回数据
	retmsg["errno"] = models.RECODE_OK
	retmsg["errmsg"] = models.RecodeText(models.RECODE_OK)
	retmsg["data"] = houses
	return
}

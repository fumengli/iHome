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

type houseType []models.House

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
	retHouse := make(map[string]houseType)
	var houses houseType
	num, err := o.QueryTable("house").Filter("user", userId.(int)).All(&houses)
	if err == nil {
		fmt.Printf(" houses read\n")
	}
	//组织返回数据

	var i int64
	for i = 0; i < num; i++ {
		houses[i].User.Id = userId.(int)
		houses[i].User.Name = this.GetSession("name").(string)
		houses[i].User.Password_hash = this.GetSession("password").(string)
		houses[i].User.Mobile = this.GetSession("mobile").(string)
		houses[i].User.Real_name = this.GetSession("real_name").(string)
		houses[i].User.Id_card = this.GetSession("id_card").(string)
		houses[i].User.Avatar_url = this.GetSession("avatar_url").(string)
		//houses[i].Area.Name = this.GetSession("aname").(string)
		houses = append(houses, houses[i])
	}
	retHouse["houses"] = houses
	retmsg["errno"] = models.RECODE_OK
	retmsg["errmsg"] = models.RecodeText(models.RECODE_OK)
	retmsg["data"] = retHouse
	return

}

package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"iHome/models"
)

type AreasController struct {
	beego.Controller
}

func (this *AreasController) RetrunData(retmsg interface{}) {
	this.Data["json"] = retmsg
	this.ServeJSON()
}

func (this *AreasController) GetAreas() {
	beego.Info("=========GetAreas succ!=============")
	o := orm.NewOrm()
	var areas []models.Area
	var retmsg = make(map[string]interface{})
	retmsg["errno"] = models.RECODE_OK
	retmsg["errmsg"] = models.RecodeText(models.RECODE_OK)
	defer this.RetrunData(retmsg)
	if ret, err := o.QueryTable("area").All(&areas); err != nil {
		retmsg["errno"] = models.RECODE_DBERR
		retmsg["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	} else if ret == 0 {
		retmsg["errno"] = models.RECODE_NODATA
		retmsg["errmsg"] = models.RecodeText(models.RECODE_NODATA)
		return
	}
	//succ
	retmsg["data"] = areas
	return
}

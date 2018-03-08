package controllers

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/orm"
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
	retmsg["errno"] = models.RECODE_OK
	retmsg["errmsg"] = models.RecodeText(models.RECODE_OK)
	defer this.RetrunData(retmsg)
	/*
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
	*/
}

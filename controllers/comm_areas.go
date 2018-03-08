package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	"iHome/models"
	"time"
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
	//1.尝试先从redis中读数据
	//url, _ := models.GetUrl()
	//{"key":"collectionName","conn":":6039","dbNum":"0","password":"thePassWord"}
	rd, err := cache.NewCache("redis", `{"key":"iHome","conn":"192.168.69.233:6379","dbNum":"0"}`)
	if err != nil {
		fmt.Println("err!!!!", err)
		retmsg["errno"] = models.RECODE_DBERR
		retmsg["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return

	}

	areas_struct := rd.Get("areasInfo")
	if areas_struct != nil {
		//beego.Info("read areas from redis")
		var areas_struct_data []models.Area
		json.Unmarshal(areas_struct.([]byte), &areas_struct_data)
		retmsg["data"] = areas_struct_data
		return
	}
	//2.从数据数据库中查询,并写入redis
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
	//save to redis
	if areas_info_jsonstr, err := json.Marshal(areas); err != nil {
		beego.Info("areas write to  redis fail!")
		return
	} else {
		beego.Info("areas write ti redis succ!")
		rd.Put("areasInfo", areas_info_jsonstr, time.Second*100)
	}
	return
}

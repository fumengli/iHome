package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"iHome/models"
	"strconv"
	_ "time"
)

type GetHouseDetailsController struct {
	beego.Controller
}

func (this *GetHouseDetailsController) RetrunData(retmsg interface{}) {
	this.Data["json"] = retmsg
	this.ServeJSON()
}

func (this *GetHouseDetailsController) GetHouseDetails() {
	beego.Info("=========GetHouseDetails succ!=============")
	var retmsg = make(map[string]interface{})
	retmsg["errno"] = models.RECODE_NODATA
	retmsg["errmsg"] = models.RecodeText(models.RECODE_NODATA)
	defer this.RetrunData(retmsg)
	userId := this.GetSession("user_id").(int)
	//根据houseId 查询该用户信息
	//得到用户请求的房屋id
	houseIdstr := this.Ctx.Input.Param(":id")
	houseId, _ := strconv.Atoi(houseIdstr)

	retHouse := make(map[string]interface{})
	house_map := make(map[string]interface{})
	//first find from cache
	var houseTemp models.House = models.House{Id: houseId}
	//var facilityTemp []*models.Facility
	//从数据库中查找
	o := orm.NewOrm()
	//组织返回数据
	//o.QueryTable("house").RelatedSel().One(&houseTemp)
	o.Read(&houseTemp)

	o.LoadRelated(&houseTemp, "Facilities")
	o.LoadRelated(&houseTemp, "Area")
	o.LoadRelated(&houseTemp, "Images")
	o.LoadRelated(&houseTemp, "User")
	o.LoadRelated(&houseTemp, "Orders")
	tmp_facilities := make([]int, 0)
	tmp_img_urls := make([]string, 0)
	for _, val1 := range houseTemp.Facilities {
		tmp_facilities = append(tmp_facilities, val1.Id)
	}
	for _, val := range houseTemp.Images {
		tmp_img_urls = append(tmp_img_urls, val.Url)
	}
	house_map["hid"] = houseId
	house_map["acreage"] = houseTemp.Acreage
	house_map["address"] = houseTemp.Address
	house_map["beds"] = houseTemp.Beds
	house_map["capacity"] = houseTemp.Capacity
	house_map["comments"] = nil
	house_map["deposit"] = houseTemp.Deposit
	house_map["facilities"] = tmp_facilities
	house_map["img_urls"] = []string{"", ""}
	house_map["max_days"] = houseTemp.Max_days
	house_map["min_days"] = houseTemp.Max_days
	house_map["price"] = houseTemp.Price
	house_map["room_count"] = houseTemp.Room_count
	house_map["title"] = houseTemp.Title
	house_map["unit"] = houseTemp.Unit
	house_map["user_avatar"] = houseTemp.User.Avatar_url
	house_map["user_id"] = houseTemp.User.Id
	house_map["user_name"] = houseTemp.User.Name

	retHouse["house"] = house_map
	retHouse["user_id"] = userId
	//组织返回json数据
	retmsg["errno"] = models.RECODE_OK
	retmsg["errmsg"] = models.RecodeText(models.RECODE_OK)

	retmsg["data"] = retHouse

	return
}

package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"iHome/models"
	"strconv"
)

type UploadHouseInfoController struct {
	beego.Controller
}

func (this *UploadHouseInfoController) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *UploadHouseInfoController) UpHouseInfo() {
	beego.Info("=========post uploadhouseinfo succ========")

	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	defer this.RetData(resp)
	var uploadHouseData = make(map[string]interface{})
	//huoqu客户端输入的房屋信息
	json.Unmarshal(this.Ctx.Input.RequestBody, &uploadHouseData)

	if uploadHouseData["title"] == "" {
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}

	//定义房屋
	house := models.House{}
	o := orm.NewOrm()

	//根据客户端信息获取user结构体
	user_id := this.GetSession("user_id").(int)
	user := models.User{Id: user_id}
	if err := o.Read(&user); err != nil {
		beego.Info("read user err...")
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	//根据客户端信息获取area结构体
	area_id, _ := strconv.Atoi(uploadHouseData["area_id"].(string))
	area := models.Area{Id: area_id}
	if err := o.Read(&area); err != nil {
		beego.Info("read area err...")
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	//房屋信息
	house.User = &user
	house.Area = &area
	house.Title = uploadHouseData["title"].(string)
	house.Price = uploadHouseData["price"].(int)
	house.Address = uploadHouseData["address"].(string)
	house.Room_count = uploadHouseData["room_count"].(int)
	house.Acreage = uploadHouseData["acreage"].(int)
	house.Unit = uploadHouseData["unit"].(string)
	house.Capacity = uploadHouseData["capacity"].(int)
	house.Beds = uploadHouseData["beds"].(string)
	house.Deposit = uploadHouseData["deposit"].(int)
	house.Min_days = uploadHouseData["min_days"].(int)
	house.Max_days = uploadHouseData["max_days"].(int)

	//数据库中插入房屋信息
	id, err := o.Insert(&house)
	if err != nil {
		beego.Info("insert house err")
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	//房屋设施
	facilitys := make([]*models.Facility, 0)
	//房屋设施编号获取数据
	faciCode := uploadHouseData["facility"].([]string)
	for i := 0; i < len(faciCode); i++ {
		fid, _ := strconv.Atoi(faciCode[i])
		faci := models.Facility{Id: fid}
		if err := o.Read(&faci); err != nil {
			beego.Info("read facility err")
			resp["errno"] = models.RECODE_DBERR
			resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
			return
		}

		facilitys = append(facilitys, &faci)
	}
	//房屋信息中心插入房屋设施信息
	m_house := models.House{Id: int(id)}
	m2m := o.QueryM2M(&m_house, "Facilitys")
	num, err := m2m.Add(facilitys)
	if err == nil {
		beego.Info("Added nums: ", num)
	}

	//返回给前端的数据
	house_map := make(map[string]int64)
	house_map["house_id"] = id
	resp["data"] = house_map
}

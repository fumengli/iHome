package models

import (
	_ "fmt"
	_ "github.com/astaxie/beego"
)

/*
func SetUserToSession(user User, ctl beego.Controller) {
	user.Avatar_url = ctl.SetSession("avatar_url")
	user.Id = ctl.SetSession("user_id")
	user.Id_card = ctl.SetSession("id_card")
	user.Mobile = ctl.SetSession("mobile")
	user.Name = ctl.SetSession("name")
	user.Password_hash = ctl.SetSession("password_hash")
	user.Real_name = ctl.SetSession("real_name")
}
func GetUserFromSession(ctl beego.Controller) (user User) {
	user.Avatar_url = ctl.GetSession("avatar_url").(string)
	user.Id = ctl.GetSession("user_id").(int)

	user.Id_card = ctl.GetSession("id_card").(string)

	user.Mobile = ctl.GetSession("mobile").(string)

	user.Name = ctl.GetSession("name").(string)

	user.Password_hash = ctl.GetSession("password_hash").(string)

	user.Real_name = ctl.GetSession("real_name").(string)

	return
}
func DelUserFromSession(ctl beego.Controller) {
	ctl.DelSession("avatar_url")
	ctl.DelSession("user_id")
	ctl.DelSession("id_card")
	crl.DelSession("mobile")
	ctl.DelSession("name")
	ctl.DelSession("password_hash")
	ctl.DelSession("real_name")
}

/*
func GetUserFromSession(ctl beego.Controller) (user User) {
}
func DelUserFromSession(ctl beego.Controller) {

}
*/

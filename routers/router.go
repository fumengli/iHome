package routers

import (
	"github.com/astaxie/beego"
	"iHome/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//初始化地区选择 areas
	beego.Router("/api/v1.0/areas", &controllers.AreasController{}, "get:GetAreas")
	//初始化api/v1.0/houses/index
	beego.Router("/api/v1.0/houses/index", &controllers.HouseIndexController{}, "get:GetHouseIndex")
	//初始化api/v1.0/session reg
	beego.Router("/api/v1.0/session", &controllers.SessionController{}, "get:GetSessionName;delete:DelSessionName")
	//初始化api/v1.0/session login
	beego.Router("/api/v1.0/sessions", &controllers.LoginUserController{}, "post:LoginUser")

	//初始化 api/v1.0/users
	beego.Router("/api/v1.0/users", &controllers.RegNewUserController{}, "post:RegNewUser")
	//登录成功后加载当前用户信息
	beego.Router("/api/v1.0/user", &controllers.LoadUserController{}, "get:LoadUser")
	//加载用户头像
	beego.Router("/api/v1.0/user/avatar", &controllers.UploadAvatarController{}, "post:UploadAvatar")
	//修改用户名 /api/v1.0/user/name
	beego.Router("/api/v1.0/user/name", &controllers.ModifyUserNameController{}, "put:ModifyUserName")
}

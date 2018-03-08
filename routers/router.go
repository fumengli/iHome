package routers

import (
	"github.com/astaxie/beego"
	"iHome/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//初始化地区选择 areas
	beego.Router("/api/v1.0/areas", &controllers.AreasController{}, "get:GetAreas")
	//初始化api/v1.0/session reg
	beego.Router("/api/v1.0/session", &controllers.SessionController{}, "get:GetSessionName;delete:DelSessionName")
	//初始化api/v1.0/session login
	beego.Router("/api/v1.0/sessions", &controllers.LoginUserController{}, "post:LoginUser")
	//====================USER================================
	//初始化 api/v1.0/users
	beego.Router("/api/v1.0/users", &controllers.RegNewUserController{}, "post:RegNewUser")
	//登录成功后加载当前用户信息
	beego.Router("/api/v1.0/user", &controllers.LoadUserController{}, "get:LoadUser")
	//加载用户头像  [/api/v1.0/user/avatar]
	beego.Router("/api/v1.0/user/avatar", &controllers.UploadAvatarController{}, "post:UploadAvatar")
	//修改用户名 /api/v1.0/user/name
	beego.Router("/api/v1.0/user/name", &controllers.ModifyUserNameController{}, "put:ModifyUserName")
	//实名认证检测  api/v1.0/user--GET&POST //发布房源
	beego.Router("/api/v1.0/user/auth", &controllers.AuthUserController{}, "get:AuthUser;post:AuthUserPOST")

	//发布房源 GET [/api/v1.0/user/houses]
	beego.Router("/api/v1.0/user/houses", &controllers.HouseLoadController{}, "get:GetHouseLoad")

	//====================House========================
	//初始化   api/v1.0/houses/index
	beego.Router("/api/v1.0/houses/index", &controllers.HouseIndexController{}, "get:GetHouseIndex")

}

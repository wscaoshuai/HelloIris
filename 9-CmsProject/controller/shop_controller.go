package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"HelloIris/9-CmsProject/service"
)

//商店功能模块控制结构体
type ShopController struct {
	//上下文对象
	Ctx     iris.Context
	Service service.ShopService
	Session *sessions.Session
}

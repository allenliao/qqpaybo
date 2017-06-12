package routers

import (
	"qqpaybo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/Login", &controllers.LoginController{})
	beego.Router("/Index", &controllers.IndexController{})

	beego.Router("/Payment", &controllers.PaymentController{})
	beego.Router("/Payment/:QQ([0-9]+)", &controllers.PaymentController{})
	beego.Router("/Payment/AddQQPay", &controllers.PaymentAddQQPayController{})

	beego.Router("/PaymentEnquiry", &controllers.PaymentEnquiryController{})
	beego.Router("/GroupManagement", &controllers.GroupManagementController{})

	beego.Router("/GroupManagement/AddQQGroup", &controllers.GroupManagementAddQQGroupController{})
	beego.Router("/GroupManagement/DelQQGroup/:id([0-9]+)", &controllers.GroupManagementDelQQGroupController{})
	beego.Router("/GroupManagement/EditQQGroup", &controllers.GroupManagementEditQQGroupController{})

}

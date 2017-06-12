package controllers

import "qqpaybo/storage"

type PaymentController struct {
	BaseController
}

func (this *PaymentController) Get() {
	this.VerifyLogin()
	QQ := this.Input().Get("QQ")
	this.Data["qqExpireInfoMap"] = storage.DB_GetQQExpireData(QQ)

	this.Data["groupInfoMap"] = storage.DB_GetAllQQGroup()
	this.TplName = "payment.html"
}

func (this *PaymentController) Post() {
	//查詢
	this.VerifyLogin()
	this.TplName = "payment.html"
}

type PaymentAddQQPayController struct {
	BaseController
}

func (this *PaymentAddQQPayController) Post() {
	//新增
	this.VerifyLogin()
	qq := this.Input().Get("qq")
	idGroup := this.Input().Get("idGroup")
	startdate := this.Input().Get("startdate")
	extenddays := this.Input().Get("extenddays")
	amount := this.Input().Get("amount")
	storage.DB_AddQQPay(qq, idGroup, startdate, extenddays, amount)

	this.Redirect("/Payment?QQ="+qq, 302)
}

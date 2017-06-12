package controllers

type PaymentEnquiryController struct {
	BaseController
}

func (this *PaymentEnquiryController) Get() {
	this.VerifyLogin()
	this.TplName = "paymentenquiry.html"
}

func (this *PaymentEnquiryController) Post() {

	this.TplName = "paymentenquiry.html"
}

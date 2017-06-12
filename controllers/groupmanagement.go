package controllers

import "qqpaybo/storage"

type GroupManagementController struct {
	BaseController
}

func (this *GroupManagementController) Get() {
	this.VerifyLogin()

	this.Data["groupInfoMap"] = storage.DB_GetAllQQGroup()
	this.TplName = "groupmanagement.html"
}

type GroupManagementAddQQGroupController struct {
	BaseController
}

func (this *GroupManagementAddQQGroupController) Post() {
	this.VerifyLogin()
	groupQQ := this.Input().Get("groupQQ")
	groupname := this.Input().Get("groupname")
	storage.DB_AddQQGroup(groupQQ, groupname)

	this.Redirect("/GroupManagement", 302)

}

type GroupManagementDelQQGroupController struct {
	BaseController
}

func (this *GroupManagementDelQQGroupController) Get() {
	this.VerifyLogin()
	IdGroup := this.Ctx.Input.Param(":id")

	storage.DB_DelQQGroup(IdGroup)
	this.Redirect("/GroupManagement", 302)
}

type GroupManagementEditQQGroupController struct {
	BaseController
}

func (this *GroupManagementEditQQGroupController) Post() {
	this.VerifyLogin()
	IdGroup := this.Input().Get("idGroup")
	groupQQ := this.Input().Get("groupQQ")
	groupname := this.Input().Get("groupname")

	storage.DB_EditQQGroup(IdGroup, groupQQ, groupname)
	this.Redirect("/GroupManagement", 302)
}

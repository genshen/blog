package admin

type PanelController struct {
	BaseController
}

func (this *PanelController) Get() {
	this.Data["Email"] = "astaxie@gmail.com"
	this.EnableRender = true
	this.TplName = "admin/panel/index.html"
}
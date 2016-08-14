package admin

type PanelController struct {
	BaseController
}

func (this *PanelController) Get() {
	//this.Ctx.Input.
	this.Data["Email"] = "astaxie@gmail.com"
	this.EnableRender = true
	this.TplName = "home/index.html"
}
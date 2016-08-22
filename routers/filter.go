package routers

import (
	"strings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"gensh.me/blog/controllers/admin"
	"net/http"
	"os"
)

func intiFilter() {
	beego.InsertFilter(admin.AdminPrefix + "/*", beego.BeforeRouter, FilterAuth)
	if beego.BConfig.RunMode == beego.DEV{
		beego.InsertFilter(adminStaticPrefix + "/*", beego.BeforeRouter, ServeAdminStaticDev)
	}else{
		//beego.InsertFilter(adminStaticPrefix + "/*", beego.BeforeRouter, ServeAdminStatic)
	}
}

func FilterAuth(ctx *context.Context) {
	var baseUri string
	index := strings.IndexByte(ctx.Request.RequestURI, '?')
	if index == -1 {
		index = len(ctx.Request.RequestURI)
	}
	baseUri = ctx.Request.RequestURI[0:index]
	if baseUri != admin.AdminAuthUri && baseUri != admin.AdminSignOutUri {
		_, ok := ctx.Input.Session(admin.UserId).(string)
		if !ok {
			var urlTail = ctx.Request.RequestURI[len(admin.AdminPrefix):] //thr router filter means len(RequestURI)>= len(AdminPrefix)
			if strings.HasPrefix(urlTail, adminApi) {
				ctx.Output.Status = 401
				ctx.Output.Header("Content-type","application/json")
				ctx.Output.Body([]byte("{\"Ststau\":2,\"Error\":\"需要登录后才能操作\"}")) //todo
			} else {
				ctx.Redirect(302, admin.AdminAuthUri + "?next=" + ctx.Request.RequestURI)
			}
		}
	}
}

//for controller,if some output has sent,the code will not run
//so context.Abort is net necessary
func ServeAdminStaticDev(context *context.Context) {
	if _, ok := context.Input.Session(admin.UserId).(string); !ok {
		context.Output.Status = 401
		context.Output.Body([]byte("UnAuthenticated"))
	} else {
		//var filePath = context.Request.RequestURI[len(adminStaticPrefix):]
		//fmt.Println(filePath)
		fi, err := os.Stat("static"+ context.Request.RequestURI)
		if err != nil {
			context.Output.Status = 404
			context.Output.Body([]byte("not found"))
			return
		}
		file, _ := os.Open("static"+context.Request.RequestURI)
		defer file.Close()
		if strings.HasSuffix(context.Request.RequestURI,"css"){
			context.Output.Header("Content-Type", "text/css; charset=utf-8")
		}
		http.ServeContent(context.ResponseWriter, context.Request, "", fi.ModTime(), file)
	}
}
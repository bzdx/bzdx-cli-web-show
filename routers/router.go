package routers

import (
	"{{ appname }}/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// -------------------------------------- 前端路由 --------------------------------------
    beego.Router("/", &controllers.PageController{})

	// -------------------------------------- seo --------------------------------------
	beego.Router("/sitemap.xml", &controllers.SEOFile{}, "GET:SitemapXML")
	beego.Router("/sitemap.html", &controllers.SEOFile{}, "GET:SitemapHTML")
	beego.Router("/sitemap.txt", &controllers.SEOFile{}, "GET:SitemapTXT")
	beego.Router("/sitemap.xsl", &controllers.SEOFile{}, "GET:SitemapXSL")
	beego.Router("/MI4D3ErdTnJ0H1F8CxjPJugBTdwd313onvNxPS6V8j0.html", &controllers.SEOFile{}, "GET:GoogleSEO")

	// -------------------------------------- 404 --------------------------------------
	beego.ErrorController(&controllers.ShowErr{})
}

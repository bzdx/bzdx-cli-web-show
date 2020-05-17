package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
)

// 页面路由
type PageController struct {
	beego.Controller
}
// seo文件
type SEOFile struct {
	beego.Controller
}
// 错误处理
type ShowErr struct {
	beego.Controller
}

// -------------------------------------- 前端路由 --------------------------------------
func (c *PageController) Get() {
	c.TplName = "index.html"
}

// -------------------------------------- SEO文件 --------------------------------------
// sitemap
func (c *SEOFile) SitemapXML() {
	file, _ := ioutil.ReadFile("./sitemap.xml")
	c.Ctx.WriteString(string(file))
}
func (c *SEOFile) SitemapXSL() {
	file, _ := ioutil.ReadFile("./sitemap.xsl")
	c.Ctx.WriteString(string(file))
}
func (c *SEOFile) SitemapTXT() {
	file, _ := ioutil.ReadFile("./sitemap.txt")
	c.Ctx.WriteString(string(file))
}
func (c *SEOFile) SitemapHTML() {
	file, _ := ioutil.ReadFile("./sitemap.html")
	c.Ctx.WriteString(string(file))
}
func (this *SEOFile) GoogleSEO() {
	this.TplName = "MI4D3ErdTnJ0H1F8CxjPJugBTdwd313onvNxPS6V8j0.html"
}


// -------------------------------------- err --------------------------------------
func (this *ShowErr) Error404() {
	this.TplName = "404.html"
}
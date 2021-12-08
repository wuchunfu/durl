package controllers

import "github.com/beego/beego/v2/core/config"

// Index 首页入口
func (c *BackendController) Index() {
	// xsrf 值
	c.Data["xsrf_token"] = c.XSRFToken()

	// 百度统计key
	sConf, _ := config.String("Statistical_Baidu")
	if sConf != "" {
		c.Data["Statistical_Baidu_Key"] = sConf
	}
	c.TplName = "index.html"
}

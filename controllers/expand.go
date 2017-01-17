package controllers

import (
    "github.com/astaxie/beego"
)

type ExpandController struct {
    beego.Controller
}

func (c *ExpandController) Get() {
    var result ShortResult
    shorturl := c.Input().Get("shorturl")
    result.UrlShort = shorturl

    if urlcache.IsExist(shorturl) {
        result.UrlLong = urlcache.Get(shorturl).(string)
    } else {
        result.UrlLong = ""
    }

    c.Data["json"] = result
    c.ServeJSON()
}

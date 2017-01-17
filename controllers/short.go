package controllers

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/cache"
    "github.com/Sai628/beego-sample-shorturl/models"
)

var (
    urlcache cache.Cache
)

func init() {
    urlcache, _ = cache.NewCache("memory", `{"interval: 0}`)
}

type ShortResult struct {
    UrlShort string
    UrlLong string
}

type ShortController struct {
    beego.Controller
}

func (c *ShortController) Get() {
    var result ShortResult
    longurl := c.Input().Get("longurl")
    result.UrlLong = longurl
    
    urlmd5 := models.GetMD5(longurl)
    beego.Info(urlmd5)
    if urlcache.IsExist(urlmd5) {
        beego.Info("has exist: ", longurl)
        result.UrlShort = urlcache.Get(urlmd5).(string)
    } else {
        result.UrlShort = models.Generate()
        beego.Info("not exist, generate a new short url:", result.UrlShort)
        err := urlcache.Put(urlmd5, result.UrlShort, 9999999)
        if err != nil {
            beego.Info(err)
        }
        err = urlcache.Put(result.UrlShort, longurl, 9999999)
        if err != nil {
            beego.Info(err)
        }
    }

    c.Data["json"] = result
    c.ServeJSON()
}

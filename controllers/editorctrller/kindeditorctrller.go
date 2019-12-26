package editorctrller

import (
	"fmt"
	"github.com/astaxie/beego"
	"html_beego/conf"
	"html_beego/utils/fileutil"
	"io"
	"os"
)

type Kindeditorctrller struct {
	beego.Controller
}

func (c *Kindeditorctrller) Index() {

}

func (c *Kindeditorctrller) Upload() {

	if c.Ctx.Input.IsPost(){

		path := conf.Info["file_upload_collection"]

		fResult := func(err error, filename ...string){
			m := make(map[string]interface{})

			if err != nil{
				m["error"] = 1
				m["message"] = err.Error()
			}else{
				m["error"] = 0
				m["url"] = conf.Dmain + path + filename[0]
			}

			c.Data["json"] = m
			c.ServeJSON()
		}

		f, fHeader, err := c.GetFile("imgFile")
		if err != nil{
			fResult(err)
			return
		}
		defer f.Close()

		//logs.Debug("fileinfo:%#v", fHeader)

		if !fileutil.PathIsExist(path){
			err = os.Mkdir(path, os.ModePerm)
			if err != nil{
				fResult(err)
				return
			}
		}

		fw, err := os.Create(path+fHeader.Filename)
		if err != nil{
			fResult(err)
			return
		}
		defer fw.Close()

		bytes, err := io.Copy(fw, f)
		if err != nil{
			fResult(err)
			return
		}

		if bytes != fHeader.Size{
			fResult(fmt.Errorf("保存数据出错"))
			return
		}

		fResult(nil, fHeader.Filename)
	}
}
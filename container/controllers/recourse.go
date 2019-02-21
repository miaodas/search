package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"search/container/models"
	"search/hiendy"
)

type ResourceController struct {
	beego.Controller
}

type Test struct {
	Name string
	Sex  string
	id   int
}

type ReParams struct {
	Limit  int
	Offset int
}

func (this *ResourceController) Get() {
	var list []hiendy.Resource
	limit, err := this.GetInt("limit")
	offset, err := this.GetInt("offset", 0)
	if err != nil {
		this.Data["json"] = struct {
			Success  string
			Code     int
			ErrorDes string
		}{
			"failed",
			-1,
			"params error! unsupport params",
		}
		this.ServeJSON()
		return
	}
	err = models.GetResources(offset, limit, &list)
	if err != nil {
		this.Data["json"] = struct {
			Success string
			Code    int
		}{
			"failed",
			-1,
		}
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = struct {
			Success string
			Code    int
			Result  []hiendy.Resource
		}{
			"success",
			1,
			list,
		}
	}
	this.ServeJSON()
}

func (this *ResourceController) Post() {
	carNumber := this.GetString("CarNumber")
	id := this.GetString(":id")
	if carNumber == "" {
		this.Ctx.WriteString("CarNumber is empty")
		return
	}
	fmt.Printf(carNumber)
	this.Data["json"] = struct {
		Success   string
		Code      int
		CarNumber string
		Id        string
	}{
		"success",
		0,
		carNumber,
		id,
	}
	this.ServeJSON()
}

func (this *ResourceController) Put() {
	id, err := this.GetInt(":id", 2)
	if err != nil {
		this.Data["json"] = struct {
			Success string
			Code    int
		}{
			"put failed",
			-1,
		}
	}
	this.Data["json"] = struct {
		Success string
		Code    int
		Id      int
	}{
		"success",
		0,
		id,
	}
	this.ServeJSON()
}

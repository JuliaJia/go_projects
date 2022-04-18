package main

import (
	"fmt"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Ctx.WriteString("Get")
}

func (c *HomeController) Post() {
	c.Ctx.WriteString("Post")
}

func (c *HomeController) Delete() {
	c.Ctx.WriteString("Delete")
}

func (c *HomeController) Put() {
	c.Ctx.WriteString("Put")
}

type HostController struct {
	beego.Controller
}

func (c *HostController) Get() {
	id := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString(fmt.Sprintf("Host: %s", id))
}

type TaskController struct {
	beego.Controller
}

func (c *TaskController) Query() {
	id := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString(fmt.Sprintf("Task: %s", id))
}

func (c *TaskController) Modify() {
	id := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString(fmt.Sprintf("ModifyTask: %s", id))
}

func main() {
	beego.Router("/home/", &HomeController{})
	beego.Router("/host/?:id/", &HostController{})
	beego.Router("/task/?:id/", &TaskController{}, "get:Query;Put:Modify")
	beego.Run()
}

package web

import (
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris"
)

type UserController struct {
	mvc.C
}

func (c *UserController) Get() (string, string) {
	return "{\"name\":123}", "Application/Json"
}

func (c *UserController) GetBy(name string) string {
	return "Hello, " + name
}
func (c *UserController) GetWelcome() (string, int) {
	return "This is GetWelcome action func...", iris.StatusOK
}
func (c *UserController) GetWelcomeBy(name string, numTimes int) {
	c.Ctx.Writef("Hello %s, times %d", name, numTimes)
}

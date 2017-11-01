package web

import (
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris"
	"fmt"
)

type AddressController struct {
	mvc.C
}

func (c *AddressController) Get() (int, error) {
	return iris.StatusBadRequest, fmt.Errorf("{\"good\":11}")
}

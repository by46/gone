package web

import (
	"testing"
	"github.com/kataras/iris/httptest"
	"github.com/kataras/iris"
)

func TestAuthBasic(t *testing.T) {
	app := IrisServeWithBasic()
	e := httptest.New(t, app)

	e.GET("/").Expect().Status(iris.StatusUnauthorized)
	e.GET("/admin").Expect().Status(iris.StatusUnauthorized)
	e.GET("/admin").WithBasicAuth("benjamin", "12345").
		Expect().Status(iris.StatusOK).Body().Equal("/admin benjamin: 12345")
	e.GET("/admin/settings").WithBasicAuth("wendy1", "12345").
		Expect().Status(iris.StatusUnauthorized)
}

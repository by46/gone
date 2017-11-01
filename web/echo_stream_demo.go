package web

import (
	"github.com/labstack/echo"
	"net/http"
	"encoding/json"
	"time"
	"os"
)

type GeoLocation struct {
	Altitude  float64
	Latitude  float64
	Longitude float64
}

var locations = []GeoLocation{
	{-97, 37.819929, -122.478255},
	{1899, 39.096849, -120.032351},
	{2619, 37.865101, -119.538329},
	{42, 33.812092, -117.918974},
	{15, 37.77493, -122.419416},
}

func EchoStreamServe() *echo.Echo {
	app := echo.New()

	app.GET("/", func(ctx echo.Context) error {
		ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextPlain)
		ctx.Response().WriteHeader(http.StatusOK)
		for _, l := range locations {
			if err := json.NewEncoder(ctx.Response()).Encode(l); err != nil {
				return err
			}
			ctx.Response().Flush()
			time.Sleep(1 * time.Second)
		}
		return nil
	})

	app.GET("/profile", func(ctx echo.Context) error {
		ctx.Response().Header().Set(echo.HeaderContentType, "Image/jpeg")
		src, err := os.Open("land.jpg")
		if err != nil {
			return err
		}
		defer src.Close()
		ctx.Response().WriteHeader(http.StatusOK)
		buff := make([]byte, 1024*4)
		count := len(buff)
		for count == len(buff) {
			count, err = src.Read(buff)
			if err != nil {
				return err
			}
			if count <= 0 {
				break
			}
			ctx.Response().Write(buff[:count])
			ctx.Response().Flush()
		}
		return nil
	})
	return app
}

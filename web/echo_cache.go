package web

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"time"
	"strings"
	"fmt"
)

const (
	RFC822                = "Mon, 02 Jan 2006 15:04:05 GMT"
	HeaderLastModified    = "Last-Modified"
	HeaderCacheControl    = "Cache-Control"
	HeaderExpires         = "Expires"
	HeaderETag            = "ETag"
	HeaderIfModifiedSince = "If-Modified-Since"
	HeaderIfNoneMatch     = "If-None-Match"
)

func Time2RFC(time time.Time) string {
	buf := make([]byte, 0, len(RFC822))
	buf = time.AppendFormat(buf, RFC822)
	return string(buf)
}

func RFC2Time(line string) (time.Time, error) {
	return time.Parse(RFC822, line)
}
func ServeEchoCache() http.Handler {
	app := echo.New()

	app.Use(middleware.Logger())
	app.GET("/cache2", func(ctx echo.Context) error {
		now := time.Now().UTC().Add(5 * time.Second)
		ctx.Response().Header().Add("Expires", Time2RFC(now))
		ctx.Response().Header().Add("Cache-Control", "max-age=5")
		return ctx.HTML(http.StatusOK, "<html><head><link rel='stylesheet' href='4.css'></head><body>hello<img src='2.jpg'/></body></html>")

	})

	app.GET("/2.jpg", func(ctx echo.Context) error {
		return ctx.File("1.jpg")
	})
	app.GET("/4.css", func(ctx echo.Context) error {
		lastModified, _ := RFC2Time("Sat, 16 Dec 2017 15:03:28 GMT")
		now := time.Now().UTC().Add(5 * time.Second)
		etag := "v1.0"
		if checkFresh(ctx, lastModified, etag) {
			ctx.Response().Header().Add(HeaderLastModified, Time2RFC(lastModified))
			ctx.Response().Header().Add(HeaderETag, fmt.Sprintf("\"%s\"", etag))
			ctx.Response().Header().Add(HeaderExpires, Time2RFC(now))
			ctx.Response().Header().Add(HeaderCacheControl, "max-age=15")
			ctx.Response().WriteHeader(304)
			return nil
		}
		ctx.Response().Header().Add(HeaderETag, fmt.Sprintf("\"%s\"", etag))
		ctx.Response().Header().Add(HeaderExpires, Time2RFC(now))
		ctx.Response().Header().Add(HeaderCacheControl, "max-age=15")
		ctx.Response().Header().Add(HeaderLastModified, Time2RFC(lastModified))
		return ctx.String(200, "body {}")
	})

	return app
}

func checkFresh(ctx echo.Context, lastModified time.Time, etag string) bool {
	var fresh bool
	since := ctx.Request().Header.Get(HeaderIfModifiedSince)
	if since != "" {
		sinceTime, _ := RFC2Time(since)
		if lastModified.Before(sinceTime) || lastModified.Equal(sinceTime) {
			return true
		}
	}
	tmp := ctx.Request().Header.Get(HeaderIfNoneMatch)
	if tmp != "" && !fresh {
		for _, segment := range strings.Split(tmp, ",") {
			segment = strings.TrimSpace(segment)
			if segment == etag {
				return true
			}
		}
	}
	return fresh
}

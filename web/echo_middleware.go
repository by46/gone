package web

import (
	"github.com/labstack/echo"
	"net/http"
	"time"
	"sync"
	"strconv"
)

type Stats struct {
	Uptime       time.Time      `json:"uptime"`
	RequestCount uint64         `json:"requestCount"`
	Status       map[string]int `json:"status"`
	mutex        sync.RWMutex
}

func NewStats() *Stats {
	return &Stats{
		Uptime: time.Now(),
		Status: map[string]int{},
	}
}
func (s *Stats) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if err := next(ctx); err != nil {
			ctx.Error(err)
		}
		s.mutex.Lock()
		defer s.mutex.Unlock()
		s.RequestCount ++
		status := strconv.Itoa(ctx.Response().Status)
		s.Status[status] ++
		return nil
	}
}
func (s *Stats) Handle(ctx echo.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return ctx.JSON(http.StatusOK, s)
}

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx.Response().Header().Set(echo.HeaderServer, "maguey/3.0")
		return next(ctx)
	}
}

func EchoMiddlewareServe() *echo.Echo {
	app := echo.New()
	app.Debug = true

	s := NewStats()
	app.Use(s.Process)
	app.Use(ServerHeader)

	app.GET("/stats", s.Handle)
	app.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "hello world.")
	})
	return app
}

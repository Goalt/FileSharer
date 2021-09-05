package http

import (
	"strconv"

	"github.com/Goalt/FileSharer/internal/interface/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server interface {
	Run() error
	Stop() error
}

type httpServer struct {
	e    *echo.Echo
	port int

	httpController controller.HTTPController
}

func (hs *httpServer) Run() error {
	return hs.e.Start(":" + strconv.Itoa(hs.port))
}

func (hs *httpServer) Stop() error {
	return nil
}

func NewHTTPServer(port int, httpController controller.HTTPController) Server {
	server := &httpServer{
		port:           port,
		httpController: httpController,
	}

	e := echo.New()

	// TODO в константы
	e.Static("/imgs", "html/imgs/")
	e.File("/style.css", "html/style.css")
	e.File("/", "html/index.html")
	e.File("/script.js", "html/script.js")
	e.File("/jquery-3.6.0.min.js", "html/jquery-3.6.0.min.js")

	e.POST("/api/file", server.upload)
	e.GET("/api/file", server.download)

	// Req id
	e.Use(middleware.RequestID())
	e.Use(middleware.LoggerWithConfig(middleware.DefaultLoggerConfig))

	server.e = e

	return server
}

func (hs *httpServer) upload(c echo.Context) error {
	return hs.httpController.Upload(&Context{c: c})
}

func (hs *httpServer) download(c echo.Context) error {
	return hs.httpController.Download(&Context{c: c})
}

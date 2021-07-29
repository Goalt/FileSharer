package http

import (
	"github.com/Goalt/FileSharer/internal/interface/controller"
	"github.com/labstack/echo/v4"
	"strconv"
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
	e := echo.New()

	e.Static("/imgs", "html/imgs/")
	e.File("/style.css", "html/style.css")
	e.File("/", "html/index.html")
	e.File("/script.js", "html/script.js")
	e.File("/jquery-3.6.0.min.js", "html/jquery-3.6.0.min.js")

	return &httpServer{
		e:              e,
		port:           port,
		httpController: httpController,
	}
}

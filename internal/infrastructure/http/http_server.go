package http

import (
	"net/http"
	"strconv"

	"github.com/Goalt/FileSharer/internal/config"
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

	// e.File("/", "docs/index.html")
	// e.Static("/", "docs/")
	e.Static("/", "swagger/")

	// e.Static("/imgs", "html/imgs/")
	// e.File("/style.css", "html/style.css")
	// e.File("/script.js", "html/script.js")
	// e.File("/jquery-3.6.0.min.js", "html/jquery-3.6.0.min.js")

	e.POST(config.FilePath, server.upload)
	e.GET(config.FilePath, server.download)
	e.OPTIONS(config.FilePath, server.handleOptions)

	e.Use(middleware.RequestID())
	e.Use(middleware.LoggerWithConfig(middleware.DefaultLoggerConfig))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	server.e = e

	return server
}

func (hs *httpServer) upload(c echo.Context) error {
	return hs.httpController.Upload(&Context{c: c})
}

func (hs *httpServer) download(c echo.Context) error {
	return hs.httpController.Download(&Context{c: c})
}

func (hs *httpServer) handleOptions(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

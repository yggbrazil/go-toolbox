package api

import (
	"flag"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	port       *string
	debug      *bool
	echoServer *echo.Echo
)

func init() {
	port = flag.String("port", "7700", "port for the service HTTP")
	debug = flag.Bool("debug", false, "mod of the debug")
}

func Make() *echo.Echo {
	flag.Parse()

	echoServer = echo.New()

	// Esconde o cabeçalho do Echo
	echoServer.HideBanner = true

	echoServer.Use(middleware.CORS())
	echoServer.Use(middleware.Recover())
	echoServer.Use(middleware.Gzip())
	echoServer.Use(middleware.RequestID())

	// For Heroku Work
	envPort := os.Getenv("PORT")

	if envPort != "" {
		*port = envPort
	}

	if *debug {
		echoServer.Debug = true
		echoServer.Use(middleware.Logger())
	}

	return echoServer
}

func GetEchoInstance() *echo.Echo {
	return echoServer
}

// Provides the instance of Echo
func ProvideEchoInstance(task func(e *echo.Echo)) {
	task(echoServer)
}

func Run() {
	echoServer.Logger.Fatal(echoServer.Start(":" + *port))
}

func Use(middleware ...echo.MiddlewareFunc) {
	echoServer.Use(middleware...)
}

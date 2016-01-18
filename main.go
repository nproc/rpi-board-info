package main

import (
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nproc/rpi-board-info/gql"
	"github.com/nproc/rpi-board-info/info/providers/native"
	"github.com/nproc/rpi-board-info/info/providers/psutil"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	schema, err := gql.NewSchema(
		psutil.NewBoardInformationProvider(),
		psutil.NewCPUInformationProvider(),
		psutil.NewDiskInformationProvider(),
		psutil.NewLoadAverageInformationProvider(),
		psutil.NewMemoryInformationProvider(),
		native.NewNetworkInformationProvider(),
		native.NewTemperatureInformationProvider(),
	)
	if err != nil {
		return err
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Post("/query", gql.QueryHandler(schema))

	e.Run("0.0.0.0:" + port)

	return nil
}

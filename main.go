package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	controller "github.com/pedroluis02/echo-api-sample1/src/controller"
)

func main() {
	fmt.Println("API with ECHO Framework")

	e := echo.New()
	e.Use(middleware.Logger())

	controller.NewAuthService(e)
	controller.NewUserService(e)

	go func() {
		if err := e.Start(":6000"); err != nil && err != http.ErrServerClosed {
			fmt.Println(err)
			fmt.Println("Shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		fmt.Println(err.Error())
	}
}

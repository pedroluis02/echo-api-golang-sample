package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
	//"strings" 


	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	controller "github.com/pedroluis02/echo-api-sample1/src/controller"
	"github.com/swaggo/echo-swagger"
	_ "github.com/pedroluis02/echo-api-sample1/docs"
)
//	@contact.name	Echo API Example
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@host	localhost:8080
// @BasePath  /

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	fmt.Println("API with ECHO Framework")

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Use(middleware.Logger())

	/*e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			if strings.Contains(c.Request().URL.Path, "swagger") {
				return true
			}
			return false
		},
	}))*/

	controller.NewAuthService(e)
	controller.NewUserService(e)

	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
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

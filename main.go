package main

import (
	"cleanarch/config"
	"log"

	bd "cleanarch/features/book/data"
	bhl "cleanarch/features/book/handler"
	bsrv "cleanarch/features/book/services"
	"cleanarch/features/user/data"
	"cleanarch/features/user/handler"
	"cleanarch/features/user/services"

	// "github.com/go-delve/delve/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)

	bookData := bd.BookIso(db)
	bookSrv := bsrv.BookIso(bookData)
	bookHdl := bhl.BookIso(bookSrv)

	userData := data.Isolation(db)
	userSrv := services.Isolation(userData)
	userHdl := handler.Isolation(userSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	e.POST("/register", userHdl.Register())
	e.POST("/login", userHdl.Login())
	e.GET("/users", userHdl.Profile(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PATCH("/users", userHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/users", userHdl.Deactive(), middleware.JWT([]byte(config.JWT_KEY)))

	e.POST("/books", bookHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PATCH("/books:id", bookHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/books", bookHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))
	// e.GET("/books", bookHdl.MyBook(), middleware.JWT([]byte(config.JWT_KEY)))

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}

//go:generate goagen bootstrap -d github.com/maurodelazeri/slash-auth/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/maurodelazeri/slash-auth/app"
	"github.com/maurodelazeri/slash-auth/helpers"
)

func main() {
	// Create service
	service := goa.New("Slash API")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount security middleware
	jwtMiddleware, err := helpers.NewJWTMiddleware()
	if err != nil {
		panic(err)
	}
	app.UseJWTMiddleware(service, jwtMiddleware)

	// Mount "account" controller
	c := NewAccountController(service)
	app.MountAccountController(service, c)
	// Mount "authentication" controller
	c2 := NewAuthenticationController(service)
	app.MountAuthenticationController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}

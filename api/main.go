//go:generate goagen bootstrap -d github.com/made2591/horizon/api/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/made2591/horizon/api/app"
)

func main() {
	// Create service
	service := goa.New("cellar")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "account" controller
	c := NewAccountController(service)
	app.MountAccountController(service, c)
	// Mount "bottle" controller
	c2 := NewBottleController(service)
	app.MountBottleController(service, c2)
	// Mount "health" controller
	c3 := NewHealthController(service)
	app.MountHealthController(service, c3)
	// Mount "js" controller
	c4 := NewJsController(service)
	app.MountJsController(service, c4)
	// Mount "public" controller
	c5 := NewPublicController(service)
	app.MountPublicController(service, c5)
	// Mount "swagger" controller
	c6 := NewSwaggerController(service)
	app.MountSwaggerController(service, c6)

	// Start service
	if err := service.ListenAndServe(":8081"); err != nil {
		service.LogError("startup", "err", err)
	}

}

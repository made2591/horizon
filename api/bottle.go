package main

import (
	"github.com/goadesign/goa"
	"github.com/made2591/horizon/api/app"
	"golang.org/x/net/websocket"
	"io"
)

// BottleController implements the bottle resource.
type BottleController struct {
	*goa.Controller
}

// NewBottleController creates a bottle controller.
func NewBottleController(service *goa.Service) *BottleController {
	return &BottleController{Controller: service.NewController("BottleController")}
}

// Create runs the create action.
func (c *BottleController) Create(ctx *app.CreateBottleContext) error {
	// BottleController_Create: start_implement

	// Put your logic here

	// BottleController_Create: end_implement
	return nil
}

// Delete runs the delete action.
func (c *BottleController) Delete(ctx *app.DeleteBottleContext) error {
	// BottleController_Delete: start_implement

	// Put your logic here

	// BottleController_Delete: end_implement
	return nil
}

// List runs the list action.
func (c *BottleController) List(ctx *app.ListBottleContext) error {
	// BottleController_List: start_implement

	// Put your logic here

	// BottleController_List: end_implement
	res := app.BottleCollection{}
	return ctx.OK(res)
}

// Rate runs the rate action.
func (c *BottleController) Rate(ctx *app.RateBottleContext) error {
	// BottleController_Rate: start_implement

	// Put your logic here

	// BottleController_Rate: end_implement
	return nil
}

// Show runs the show action.
func (c *BottleController) Show(ctx *app.ShowBottleContext) error {
	// BottleController_Show: start_implement

	// Put your logic here

	// BottleController_Show: end_implement
	res := &app.Bottle{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *BottleController) Update(ctx *app.UpdateBottleContext) error {
	// BottleController_Update: start_implement

	// Put your logic here

	// BottleController_Update: end_implement
	return nil
}

// Watch runs the watch action.
func (c *BottleController) Watch(ctx *app.WatchBottleContext) error {
	c.WatchWSHandler(ctx).ServeHTTP(ctx.ResponseWriter, ctx.Request)
	return nil
}

// WatchWSHandler establishes a websocket connection to run the watch action.
func (c *BottleController) WatchWSHandler(ctx *app.WatchBottleContext) websocket.Handler {
	return func(ws *websocket.Conn) {
		// BottleController_Watch: start_implement

		// Put your logic here

		// BottleController_Watch: end_implement
		ws.Write([]byte("watch bottle"))
		// Dummy echo websocket server
		io.Copy(ws, ws)
	}
}

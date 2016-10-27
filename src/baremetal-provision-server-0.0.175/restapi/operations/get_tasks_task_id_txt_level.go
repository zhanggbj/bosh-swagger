package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetTasksTaskIDTxtLevelHandlerFunc turns a function with the right signature into a get tasks task ID txt level handler
type GetTasksTaskIDTxtLevelHandlerFunc func(GetTasksTaskIDTxtLevelParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTasksTaskIDTxtLevelHandlerFunc) Handle(params GetTasksTaskIDTxtLevelParams) middleware.Responder {
	return fn(params)
}

// GetTasksTaskIDTxtLevelHandler interface for that can handle valid get tasks task ID txt level params
type GetTasksTaskIDTxtLevelHandler interface {
	Handle(GetTasksTaskIDTxtLevelParams) middleware.Responder
}

// NewGetTasksTaskIDTxtLevel creates a new http.Handler for the get tasks task ID txt level operation
func NewGetTasksTaskIDTxtLevel(ctx *middleware.Context, handler GetTasksTaskIDTxtLevelHandler) *GetTasksTaskIDTxtLevel {
	return &GetTasksTaskIDTxtLevel{Context: ctx, Handler: handler}
}

/*GetTasksTaskIDTxtLevel swagger:route GET /tasks/{taskId}/txt/{level} getTasksTaskIdTxtLevel

List baremetal server task output


*/
type GetTasksTaskIDTxtLevel struct {
	Context *middleware.Context
	Handler GetTasksTaskIDTxtLevelHandler
}

func (o *GetTasksTaskIDTxtLevel) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetTasksTaskIDTxtLevelParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
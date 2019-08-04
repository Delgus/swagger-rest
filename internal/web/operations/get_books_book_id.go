// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetBooksBookIDHandlerFunc turns a function with the right signature into a get books book ID handler
type GetBooksBookIDHandlerFunc func(GetBooksBookIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBooksBookIDHandlerFunc) Handle(params GetBooksBookIDParams) middleware.Responder {
	return fn(params)
}

// GetBooksBookIDHandler interface for that can handle valid get books book ID params
type GetBooksBookIDHandler interface {
	Handle(GetBooksBookIDParams) middleware.Responder
}

// NewGetBooksBookID creates a new http.Handler for the get books book ID operation
func NewGetBooksBookID(ctx *middleware.Context, handler GetBooksBookIDHandler) *GetBooksBookID {
	return &GetBooksBookID{Context: ctx, Handler: handler}
}

/*GetBooksBookID swagger:route GET /books/{bookID} getBooksBookId

Get a book by ID

*/
type GetBooksBookID struct {
	Context *middleware.Context
	Handler GetBooksBookIDHandler
}

func (o *GetBooksBookID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetBooksBookIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

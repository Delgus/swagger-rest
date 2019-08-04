// This file is safe to edit. Once it exists it will not be overwritten

package web

import (
	"crypto/tls"
	"github.com/delgus/swagger-rest/internal/books"
	"github.com/delgus/swagger-rest/internal/models"
	"github.com/delgus/swagger-rest/internal/store/memory"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/delgus/swagger-rest/internal/web/operations"
)

var (
	ErrorBookNotFoundMessage  = "Sorry. Book not found"
	ErrorBooksNotFoundMessage = "Sorry. Books not found"

	ErrorBookNotFound  = &models.Error{Message: &ErrorBookNotFoundMessage}
	ErrorBooksNotFound = &models.Error{Message: &ErrorBooksNotFoundMessage}
)

//go:generate swagger generate server --target ../../internal --name Library --spec ../../swagger.yml --server-package web

func configureFlags(api *operations.LibraryAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.LibraryAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	//repo
	bookRepo := new(memory.BooksRepo)
	bookRepo.Fill()
	//service
	bookService := books.NewService(bookRepo)

	api.GetBooksHandler = operations.GetBooksHandlerFunc(func(params operations.GetBooksParams) middleware.Responder {
		bs := bookService.GetAll(params.Author, params.Title)
		if len(bs) == 0 {
			return operations.NewGetBooksNotFound().WithPayload(ErrorBooksNotFound)
		}
		return operations.NewGetBooksOK().WithPayload(bs)
	})

	api.GetBooksBookIDHandler = operations.GetBooksBookIDHandlerFunc(func(params operations.GetBooksBookIDParams) middleware.Responder {
		book := bookService.GetByID(params.BookID)
		if book == nil {
			return operations.NewGetBooksBookIDNotFound().WithPayload(ErrorBookNotFound)
		}
		return operations.NewGetBooksBookIDOK().WithPayload(book)
	})

	api.PostBooksHandler = operations.PostBooksHandlerFunc(func(params operations.PostBooksParams) middleware.Responder {
		if bookService.Save(params.Book) {
			return operations.NewPostBooksCreated()
		}
		return operations.NewPostBooksInternalServerError()
	})

	api.PutBooksHandler = operations.PutBooksHandlerFunc(func(params operations.PutBooksParams) middleware.Responder {
		if bookService.Update(params.Book) {
			return operations.NewPutBooksNoContent()
		}
		return operations.NewPutBooksInternalServerError()
	})

	api.DeleteBooksBookIDHandler = operations.DeleteBooksBookIDHandlerFunc(func(params operations.DeleteBooksBookIDParams) middleware.Responder {
		if bookService.Delete(params.BookID) {
			return operations.NewDeleteBooksBookIDNoContent()
		}
		return operations.NewDeleteBooksBookIDInternalServerError()
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

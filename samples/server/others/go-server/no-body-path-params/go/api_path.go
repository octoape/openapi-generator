// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Simple no path and body param spec
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

package petstoreserver

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// PathAPIController binds http requests to an api service and writes the service results to the http response
type PathAPIController struct {
	service PathAPIServicer
	errorHandler ErrorHandler
}

// PathAPIOption for how the controller is set up.
type PathAPIOption func(*PathAPIController)

// WithPathAPIErrorHandler inject ErrorHandler into controller
func WithPathAPIErrorHandler(h ErrorHandler) PathAPIOption {
	return func(c *PathAPIController) {
		c.errorHandler = h
	}
}

// NewPathAPIController creates a default api controller
func NewPathAPIController(s PathAPIServicer, opts ...PathAPIOption) *PathAPIController {
	controller := &PathAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the PathAPIController
func (c *PathAPIController) Routes() Routes {
	return Routes{
		"Path": Route{
			"Path",
			strings.ToUpper("Get"),
			"/path/endpoint/{pathParam}",
			c.Path,
		},
	}
}

// OrderedRoutes returns all the api routes in a deterministic order for the PathAPIController
func (c *PathAPIController) OrderedRoutes() []Route {
	return []Route{
		Route{
			"Path",
			strings.ToUpper("Get"),
			"/path/endpoint/{pathParam}",
			c.Path,
		},
	}
}



// Path - summary
func (c *PathAPIController) Path(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pathParamParam := params["pathParam"]
	if pathParamParam == "" {
		c.errorHandler(w, r, &RequiredError{"pathParam"}, nil)
		return
	}
	result, err := c.service.Path(r.Context(), pathParamParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

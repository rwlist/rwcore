package utils

import (
	"net/http"

	"github.com/go-chi/render"
)

//--
// Error response payloads & renderers
// Some credits go to https://github.com/go-chi/chi/blob/master/_examples/rest/main.go
//--

// ErrResponse renderer type for handling all sorts of errors.
//
// In the best case scenario, the excellent github.com/pkg/errors package
// helps reveal information on the error, setting it on Err, and in the Render()
// method, using it to set the application-specific error code in AppCode.
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func (e ErrResponse) With(err error) ErrResponse {
	e.Err = err
	e.ErrorText = err.Error()
	return e
}

func ErrByCode(code int) ErrResponse {
	return ErrResponse{
		HTTPStatusCode: code,
		StatusText:     http.StatusText(code),
	}
}

var (
	ErrNotFound       = ErrByCode(http.StatusNotFound)
	ErrUnathorized    = ErrByCode(http.StatusUnauthorized)
	ErrBadRequest     = ErrByCode(http.StatusBadRequest)
	ErrRender         = ErrByCode(422)
	ErrInvalidRequest = ErrByCode(400)
)

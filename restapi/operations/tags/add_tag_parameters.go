package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/cyverse-de/tagger/models"
)

// NewAddTagParams creates a new AddTagParams object
// with the default values initialized.
func NewAddTagParams() AddTagParams {
	var ()
	return AddTagParams{}
}

// AddTagParams contains all the bound params for the add tag operation
// typically these are obtained from a http.Request
//
// swagger:parameters addTag
type AddTagParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*The new tag to index.
	  Required: true
	  In: body
	*/
	TagIn *models.TagIn
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *AddTagParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.TagIn
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("tagIn", "body"))
			} else {
				res = append(res, errors.NewParseError("tagIn", "body", "", err))
			}

		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.TagIn = &body
			}
		}

	} else {
		res = append(res, errors.Required("tagIn", "body"))
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

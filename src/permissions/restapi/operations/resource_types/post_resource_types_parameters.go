package resource_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/httpkit/middleware"

	"permissions/models"
)

// NewPostResourceTypesParams creates a new PostResourceTypesParams object
// with the default values initialized.
func NewPostResourceTypesParams() PostResourceTypesParams {
	var ()
	return PostResourceTypesParams{}
}

// PostResourceTypesParams contains all the bound params for the post resource types operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostResourceTypes
type PostResourceTypesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*The resource type to add.
	  Required: true
	  In: body
	*/
	ResourceTypeIn *models.ResourceTypeIn
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *PostResourceTypesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if httpkit.HasBody(r) {
		defer r.Body.Close()
		var body models.ResourceTypeIn
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("resourceTypeIn", "body"))
			} else {
				res = append(res, errors.NewParseError("resourceTypeIn", "body", "", err))
			}

		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.ResourceTypeIn = &body
			}
		}

	} else {
		res = append(res, errors.Required("resourceTypeIn", "body"))
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

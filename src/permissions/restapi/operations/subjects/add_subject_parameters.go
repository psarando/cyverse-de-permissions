package subjects

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

// NewAddSubjectParams creates a new AddSubjectParams object
// with the default values initialized.
func NewAddSubjectParams() AddSubjectParams {
	var ()
	return AddSubjectParams{}
}

// AddSubjectParams contains all the bound params for the add subject operation
// typically these are obtained from a http.Request
//
// swagger:parameters addSubject
type AddSubjectParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*The new subject to add.
	  Required: true
	  In: body
	*/
	SubjectIn *models.SubjectIn
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *AddSubjectParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if httpkit.HasBody(r) {
		defer r.Body.Close()
		var body models.SubjectIn
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("subjectIn", "body"))
			} else {
				res = append(res, errors.NewParseError("subjectIn", "body", "", err))
			}

		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.SubjectIn = &body
			}
		}

	} else {
		res = append(res, errors.Required("subjectIn", "body"))
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

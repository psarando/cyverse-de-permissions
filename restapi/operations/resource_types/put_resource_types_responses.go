package resource_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"

	"permissions/models"
)

/*PutResourceTypesCreated add resource type

swagger:response putResourceTypesCreated
*/
type PutResourceTypesCreated struct {

	// In: body
	Payload *models.ResourceTypeOut `json:"body,omitempty"`
}

// NewPutResourceTypesCreated creates PutResourceTypesCreated with default headers values
func NewPutResourceTypesCreated() *PutResourceTypesCreated {
	return &PutResourceTypesCreated{}
}

// WithPayload adds the payload to the put resource types created response
func (o *PutResourceTypesCreated) WithPayload(payload *models.ResourceTypeOut) *PutResourceTypesCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put resource types created response
func (o *PutResourceTypesCreated) SetPayload(payload *models.ResourceTypeOut) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutResourceTypesCreated) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PutResourceTypesBadRequest add resource type

swagger:response putResourceTypesBadRequest
*/
type PutResourceTypesBadRequest struct {

	// In: body
	Payload *models.ErrorOut `json:"body,omitempty"`
}

// NewPutResourceTypesBadRequest creates PutResourceTypesBadRequest with default headers values
func NewPutResourceTypesBadRequest() *PutResourceTypesBadRequest {
	return &PutResourceTypesBadRequest{}
}

// WithPayload adds the payload to the put resource types bad request response
func (o *PutResourceTypesBadRequest) WithPayload(payload *models.ErrorOut) *PutResourceTypesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put resource types bad request response
func (o *PutResourceTypesBadRequest) SetPayload(payload *models.ErrorOut) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutResourceTypesBadRequest) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PutResourceTypesInternalServerError add resource type

swagger:response putResourceTypesInternalServerError
*/
type PutResourceTypesInternalServerError struct {

	// In: body
	Payload *models.ErrorOut `json:"body,omitempty"`
}

// NewPutResourceTypesInternalServerError creates PutResourceTypesInternalServerError with default headers values
func NewPutResourceTypesInternalServerError() *PutResourceTypesInternalServerError {
	return &PutResourceTypesInternalServerError{}
}

// WithPayload adds the payload to the put resource types internal server error response
func (o *PutResourceTypesInternalServerError) WithPayload(payload *models.ErrorOut) *PutResourceTypesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put resource types internal server error response
func (o *PutResourceTypesInternalServerError) SetPayload(payload *models.ErrorOut) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutResourceTypesInternalServerError) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

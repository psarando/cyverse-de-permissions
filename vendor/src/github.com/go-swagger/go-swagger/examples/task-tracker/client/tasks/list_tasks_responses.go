package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/swag"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/examples/task-tracker/models"
)

// ListTasksReader is a Reader for the ListTasks structure.
type ListTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *ListTasksReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 422:
		result := NewListTasksUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewListTasksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewListTasksOK creates a ListTasksOK with default headers values
func NewListTasksOK() *ListTasksOK {
	return &ListTasksOK{}
}

/*ListTasksOK handles this case with default header values.

Successful response
*/
type ListTasksOK struct {
	/*the last task id known to the application
	 */
	XLastTaskID int64

	Payload []*models.TaskCard
}

func (o *ListTasksOK) Error() string {
	return fmt.Sprintf("[GET /tasks][%d] listTasksOK  %+v", 200, o.Payload)
}

func (o *ListTasksOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response header X-Last-Task-Id
	xLastTaskId, err := swag.ConvertInt64(response.GetHeader("X-Last-Task-Id"))
	if err != nil {
		return errors.InvalidType("X-Last-Task-Id", "header", "int64", response.GetHeader("X-Last-Task-Id"))
	}
	o.XLastTaskID = xLastTaskId

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTasksUnprocessableEntity creates a ListTasksUnprocessableEntity with default headers values
func NewListTasksUnprocessableEntity() *ListTasksUnprocessableEntity {
	return &ListTasksUnprocessableEntity{}
}

/*ListTasksUnprocessableEntity handles this case with default header values.

Validation error
*/
type ListTasksUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *ListTasksUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /tasks][%d] listTasksUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ListTasksUnprocessableEntity) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTasksDefault creates a ListTasksDefault with default headers values
func NewListTasksDefault(code int) *ListTasksDefault {
	return &ListTasksDefault{
		_statusCode: code,
	}
}

/*ListTasksDefault handles this case with default header values.

Error response
*/
type ListTasksDefault struct {
	_statusCode int

	XErrorCode string

	Payload *models.Error
}

// Code gets the status code for the list tasks default response
func (o *ListTasksDefault) Code() int {
	return o._statusCode
}

func (o *ListTasksDefault) Error() string {
	return fmt.Sprintf("[GET /tasks][%d] listTasks default  %+v", o._statusCode, o.Payload)
}

func (o *ListTasksDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response header X-Error-Code
	o.XErrorCode = response.GetHeader("X-Error-Code")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

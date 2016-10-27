package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"baremetal-provision-server-0.0.177/models"
)

/*GetTasksOK Successful response

swagger:response getTasksOK
*/
type GetTasksOK struct {

	// In: body
	Payload *models.Tasks `json:"body,omitempty"`
}

// NewGetTasksOK creates GetTasksOK with default headers values
func NewGetTasksOK() *GetTasksOK {
	return &GetTasksOK{}
}

// WithPayload adds the payload to the get tasks o k response
func (o *GetTasksOK) WithPayload(payload *models.Tasks) *GetTasksOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tasks o k response
func (o *GetTasksOK) SetPayload(payload *models.Tasks) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTasksOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetTasksDefault Unexpected error

swagger:response getTasksDefault
*/
type GetTasksDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTasksDefault creates GetTasksDefault with default headers values
func NewGetTasksDefault(code int) *GetTasksDefault {
	if code <= 0 {
		code = 500
	}

	return &GetTasksDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get tasks default response
func (o *GetTasksDefault) WithStatusCode(code int) *GetTasksDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get tasks default response
func (o *GetTasksDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get tasks default response
func (o *GetTasksDefault) WithPayload(payload *models.Error) *GetTasksDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tasks default response
func (o *GetTasksDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTasksDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

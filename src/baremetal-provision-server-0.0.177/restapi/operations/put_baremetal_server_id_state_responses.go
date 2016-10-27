package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"baremetal-provision-server-0.0.177/models"
)

/*PutBaremetalServerIDStateOK Successful response

swagger:response putBaremetalServerIdStateOK
*/
type PutBaremetalServerIDStateOK struct {
}

// NewPutBaremetalServerIDStateOK creates PutBaremetalServerIDStateOK with default headers values
func NewPutBaremetalServerIDStateOK() *PutBaremetalServerIDStateOK {
	return &PutBaremetalServerIDStateOK{}
}

// WriteResponse to the client
func (o *PutBaremetalServerIDStateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}

/*PutBaremetalServerIDStateDefault Unexpected error

swagger:response putBaremetalServerIdStateDefault
*/
type PutBaremetalServerIDStateDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutBaremetalServerIDStateDefault creates PutBaremetalServerIDStateDefault with default headers values
func NewPutBaremetalServerIDStateDefault(code int) *PutBaremetalServerIDStateDefault {
	if code <= 0 {
		code = 500
	}

	return &PutBaremetalServerIDStateDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put baremetal server ID state default response
func (o *PutBaremetalServerIDStateDefault) WithStatusCode(code int) *PutBaremetalServerIDStateDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put baremetal server ID state default response
func (o *PutBaremetalServerIDStateDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the put baremetal server ID state default response
func (o *PutBaremetalServerIDStateDefault) WithPayload(payload *models.Error) *PutBaremetalServerIDStateDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put baremetal server ID state default response
func (o *PutBaremetalServerIDStateDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutBaremetalServerIDStateDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

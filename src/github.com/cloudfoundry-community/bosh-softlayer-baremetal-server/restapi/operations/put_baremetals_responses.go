package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cloudfoundry-community/bosh-softlayer-baremetal-server/models"
)

/*PutBaremetalsOK Successful response

swagger:response putBaremetalsOK
*/
type PutBaremetalsOK struct {

	// In: body
	Payload *models.BaremetalsResponse `json:"body,omitempty"`
}

// NewPutBaremetalsOK creates PutBaremetalsOK with default headers values
func NewPutBaremetalsOK() *PutBaremetalsOK {
	return &PutBaremetalsOK{}
}

// WithPayload adds the payload to the put baremetals o k response
func (o *PutBaremetalsOK) WithPayload(payload *models.BaremetalsResponse) *PutBaremetalsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put baremetals o k response
func (o *PutBaremetalsOK) SetPayload(payload *models.BaremetalsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutBaremetalsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PutBaremetalsDefault Unexpected error

swagger:response putBaremetalsDefault
*/
type PutBaremetalsDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutBaremetalsDefault creates PutBaremetalsDefault with default headers values
func NewPutBaremetalsDefault(code int) *PutBaremetalsDefault {
	if code <= 0 {
		code = 500
	}

	return &PutBaremetalsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put baremetals default response
func (o *PutBaremetalsDefault) WithStatusCode(code int) *PutBaremetalsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put baremetals default response
func (o *PutBaremetalsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the put baremetals default response
func (o *PutBaremetalsDefault) WithPayload(payload *models.Error) *PutBaremetalsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put baremetals default response
func (o *PutBaremetalsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutBaremetalsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

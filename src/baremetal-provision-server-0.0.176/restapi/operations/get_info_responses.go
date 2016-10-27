package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"baremetal-provision-server-0.0.176/models"
)

/*GetInfoOK Successful response

swagger:response getInfoOK
*/
type GetInfoOK struct {

	// In: body
	Payload *models.Info `json:"body,omitempty"`
}

// NewGetInfoOK creates GetInfoOK with default headers values
func NewGetInfoOK() *GetInfoOK {
	return &GetInfoOK{}
}

// WithPayload adds the payload to the get info o k response
func (o *GetInfoOK) WithPayload(payload *models.Info) *GetInfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get info o k response
func (o *GetInfoOK) SetPayload(payload *models.Info) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetInfoDefault Unexpected error

swagger:response getInfoDefault
*/
type GetInfoDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetInfoDefault creates GetInfoDefault with default headers values
func NewGetInfoDefault(code int) *GetInfoDefault {
	if code <= 0 {
		code = 500
	}

	return &GetInfoDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get info default response
func (o *GetInfoDefault) WithStatusCode(code int) *GetInfoDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get info default response
func (o *GetInfoDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get info default response
func (o *GetInfoDefault) WithPayload(payload *models.Error) *GetInfoDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get info default response
func (o *GetInfoDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInfoDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"baremetal-provision-server-0.0.172/models"
)

/*GetBmsDeploymentNameOK Successful response

swagger:response getBmsDeploymentNameOK
*/
type GetBmsDeploymentNameOK struct {

	// In: body
	Payload *models.BaremetalServers `json:"body,omitempty"`
}

// NewGetBmsDeploymentNameOK creates GetBmsDeploymentNameOK with default headers values
func NewGetBmsDeploymentNameOK() *GetBmsDeploymentNameOK {
	return &GetBmsDeploymentNameOK{}
}

// WithPayload adds the payload to the get bms deployment name o k response
func (o *GetBmsDeploymentNameOK) WithPayload(payload *models.BaremetalServers) *GetBmsDeploymentNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get bms deployment name o k response
func (o *GetBmsDeploymentNameOK) SetPayload(payload *models.BaremetalServers) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBmsDeploymentNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetBmsDeploymentNameDefault Unexpected error

swagger:response getBmsDeploymentNameDefault
*/
type GetBmsDeploymentNameDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBmsDeploymentNameDefault creates GetBmsDeploymentNameDefault with default headers values
func NewGetBmsDeploymentNameDefault(code int) *GetBmsDeploymentNameDefault {
	if code <= 0 {
		code = 500
	}

	return &GetBmsDeploymentNameDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get bms deployment name default response
func (o *GetBmsDeploymentNameDefault) WithStatusCode(code int) *GetBmsDeploymentNameDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get bms deployment name default response
func (o *GetBmsDeploymentNameDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get bms deployment name default response
func (o *GetBmsDeploymentNameDefault) WithPayload(payload *models.Error) *GetBmsDeploymentNameDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get bms deployment name default response
func (o *GetBmsDeploymentNameDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBmsDeploymentNameDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
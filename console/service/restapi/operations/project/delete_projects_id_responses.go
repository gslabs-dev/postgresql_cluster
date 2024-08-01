// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"postgesql-cluster-console/models"
)

// DeleteProjectsIDNoContentCode is the HTTP code returned for type DeleteProjectsIDNoContent
const DeleteProjectsIDNoContentCode int = 204

/*
DeleteProjectsIDNoContent OK

swagger:response deleteProjectsIdNoContent
*/
type DeleteProjectsIDNoContent struct {
}

// NewDeleteProjectsIDNoContent creates DeleteProjectsIDNoContent with default headers values
func NewDeleteProjectsIDNoContent() *DeleteProjectsIDNoContent {

	return &DeleteProjectsIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteProjectsIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteProjectsIDBadRequestCode is the HTTP code returned for type DeleteProjectsIDBadRequest
const DeleteProjectsIDBadRequestCode int = 400

/*
DeleteProjectsIDBadRequest Error

swagger:response deleteProjectsIdBadRequest
*/
type DeleteProjectsIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ResponseError `json:"body,omitempty"`
}

// NewDeleteProjectsIDBadRequest creates DeleteProjectsIDBadRequest with default headers values
func NewDeleteProjectsIDBadRequest() *DeleteProjectsIDBadRequest {

	return &DeleteProjectsIDBadRequest{}
}

// WithPayload adds the payload to the delete projects Id bad request response
func (o *DeleteProjectsIDBadRequest) WithPayload(payload *models.ResponseError) *DeleteProjectsIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete projects Id bad request response
func (o *DeleteProjectsIDBadRequest) SetPayload(payload *models.ResponseError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteProjectsIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

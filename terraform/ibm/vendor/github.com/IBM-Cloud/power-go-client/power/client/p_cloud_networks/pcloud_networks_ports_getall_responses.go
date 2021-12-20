// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudNetworksPortsGetallReader is a Reader for the PcloudNetworksPortsGetall structure.
type PcloudNetworksPortsGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudNetworksPortsGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudNetworksPortsGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudNetworksPortsGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPcloudNetworksPortsGetallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudNetworksPortsGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudNetworksPortsGetallOK creates a PcloudNetworksPortsGetallOK with default headers values
func NewPcloudNetworksPortsGetallOK() *PcloudNetworksPortsGetallOK {
	return &PcloudNetworksPortsGetallOK{}
}

/*PcloudNetworksPortsGetallOK handles this case with default header values.

OK
*/
type PcloudNetworksPortsGetallOK struct {
	Payload *models.NetworkPorts
}

func (o *PcloudNetworksPortsGetallOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports][%d] pcloudNetworksPortsGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudNetworksPortsGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkPorts)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPortsGetallBadRequest creates a PcloudNetworksPortsGetallBadRequest with default headers values
func NewPcloudNetworksPortsGetallBadRequest() *PcloudNetworksPortsGetallBadRequest {
	return &PcloudNetworksPortsGetallBadRequest{}
}

/*PcloudNetworksPortsGetallBadRequest handles this case with default header values.

Bad Request
*/
type PcloudNetworksPortsGetallBadRequest struct {
	Payload *models.Error
}

func (o *PcloudNetworksPortsGetallBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports][%d] pcloudNetworksPortsGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudNetworksPortsGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPortsGetallUnauthorized creates a PcloudNetworksPortsGetallUnauthorized with default headers values
func NewPcloudNetworksPortsGetallUnauthorized() *PcloudNetworksPortsGetallUnauthorized {
	return &PcloudNetworksPortsGetallUnauthorized{}
}

/*PcloudNetworksPortsGetallUnauthorized handles this case with default header values.

Unauthorized
*/
type PcloudNetworksPortsGetallUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudNetworksPortsGetallUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports][%d] pcloudNetworksPortsGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudNetworksPortsGetallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPortsGetallInternalServerError creates a PcloudNetworksPortsGetallInternalServerError with default headers values
func NewPcloudNetworksPortsGetallInternalServerError() *PcloudNetworksPortsGetallInternalServerError {
	return &PcloudNetworksPortsGetallInternalServerError{}
}

/*PcloudNetworksPortsGetallInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudNetworksPortsGetallInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudNetworksPortsGetallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports][%d] pcloudNetworksPortsGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudNetworksPortsGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

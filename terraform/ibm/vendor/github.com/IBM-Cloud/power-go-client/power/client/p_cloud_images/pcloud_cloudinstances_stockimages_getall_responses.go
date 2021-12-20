// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudCloudinstancesStockimagesGetallReader is a Reader for the PcloudCloudinstancesStockimagesGetall structure.
type PcloudCloudinstancesStockimagesGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudinstancesStockimagesGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudCloudinstancesStockimagesGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudCloudinstancesStockimagesGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPcloudCloudinstancesStockimagesGetallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPcloudCloudinstancesStockimagesGetallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudCloudinstancesStockimagesGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudCloudinstancesStockimagesGetallOK creates a PcloudCloudinstancesStockimagesGetallOK with default headers values
func NewPcloudCloudinstancesStockimagesGetallOK() *PcloudCloudinstancesStockimagesGetallOK {
	return &PcloudCloudinstancesStockimagesGetallOK{}
}

/*PcloudCloudinstancesStockimagesGetallOK handles this case with default header values.

OK
*/
type PcloudCloudinstancesStockimagesGetallOK struct {
	Payload *models.Images
}

func (o *PcloudCloudinstancesStockimagesGetallOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/stock-images][%d] pcloudCloudinstancesStockimagesGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesStockimagesGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Images)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesStockimagesGetallBadRequest creates a PcloudCloudinstancesStockimagesGetallBadRequest with default headers values
func NewPcloudCloudinstancesStockimagesGetallBadRequest() *PcloudCloudinstancesStockimagesGetallBadRequest {
	return &PcloudCloudinstancesStockimagesGetallBadRequest{}
}

/*PcloudCloudinstancesStockimagesGetallBadRequest handles this case with default header values.

Bad Request
*/
type PcloudCloudinstancesStockimagesGetallBadRequest struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesStockimagesGetallBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/stock-images][%d] pcloudCloudinstancesStockimagesGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesStockimagesGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesStockimagesGetallUnauthorized creates a PcloudCloudinstancesStockimagesGetallUnauthorized with default headers values
func NewPcloudCloudinstancesStockimagesGetallUnauthorized() *PcloudCloudinstancesStockimagesGetallUnauthorized {
	return &PcloudCloudinstancesStockimagesGetallUnauthorized{}
}

/*PcloudCloudinstancesStockimagesGetallUnauthorized handles this case with default header values.

Unauthorized
*/
type PcloudCloudinstancesStockimagesGetallUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesStockimagesGetallUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/stock-images][%d] pcloudCloudinstancesStockimagesGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesStockimagesGetallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesStockimagesGetallNotFound creates a PcloudCloudinstancesStockimagesGetallNotFound with default headers values
func NewPcloudCloudinstancesStockimagesGetallNotFound() *PcloudCloudinstancesStockimagesGetallNotFound {
	return &PcloudCloudinstancesStockimagesGetallNotFound{}
}

/*PcloudCloudinstancesStockimagesGetallNotFound handles this case with default header values.

Not Found
*/
type PcloudCloudinstancesStockimagesGetallNotFound struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesStockimagesGetallNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/stock-images][%d] pcloudCloudinstancesStockimagesGetallNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesStockimagesGetallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesStockimagesGetallInternalServerError creates a PcloudCloudinstancesStockimagesGetallInternalServerError with default headers values
func NewPcloudCloudinstancesStockimagesGetallInternalServerError() *PcloudCloudinstancesStockimagesGetallInternalServerError {
	return &PcloudCloudinstancesStockimagesGetallInternalServerError{}
}

/*PcloudCloudinstancesStockimagesGetallInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudCloudinstancesStockimagesGetallInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesStockimagesGetallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/stock-images][%d] pcloudCloudinstancesStockimagesGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesStockimagesGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

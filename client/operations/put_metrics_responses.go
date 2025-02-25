// Code generated by go-swagger; DO NOT EDIT.

// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
// 	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

<<<<<<< HEAD
	models "titan/lib/firecracker/client/models"
=======
	models "github.com/Ian-Kibet/firecracker-go-sdk/client/models"
>>>>>>> b8aa219df3977843c18fb0ce7af8af072b1bf0b8
)

// PutMetricsReader is a Reader for the PutMetrics structure.
type PutMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutMetricsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutMetricsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPutMetricsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutMetricsNoContent creates a PutMetricsNoContent with default headers values
func NewPutMetricsNoContent() *PutMetricsNoContent {
	return &PutMetricsNoContent{}
}

/*PutMetricsNoContent handles this case with default header values.

Metrics system created.
*/
type PutMetricsNoContent struct {
}

func (o *PutMetricsNoContent) Error() string {
	return fmt.Sprintf("[PUT /metrics][%d] putMetricsNoContent ", 204)
}

func (o *PutMetricsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutMetricsBadRequest creates a PutMetricsBadRequest with default headers values
func NewPutMetricsBadRequest() *PutMetricsBadRequest {
	return &PutMetricsBadRequest{}
}

/*PutMetricsBadRequest handles this case with default header values.

Metrics system cannot be initialized due to bad input.
*/
type PutMetricsBadRequest struct {
	Payload *models.Error
}

func (o *PutMetricsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /metrics][%d] putMetricsBadRequest  %+v", 400, o.Payload)
}

func (o *PutMetricsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutMetricsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutMetricsDefault creates a PutMetricsDefault with default headers values
func NewPutMetricsDefault(code int) *PutMetricsDefault {
	return &PutMetricsDefault{
		_statusCode: code,
	}
}

/*PutMetricsDefault handles this case with default header values.

Internal server error.
*/
type PutMetricsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put metrics default response
func (o *PutMetricsDefault) Code() int {
	return o._statusCode
}

func (o *PutMetricsDefault) Error() string {
	return fmt.Sprintf("[PUT /metrics][%d] putMetrics default  %+v", o._statusCode, o.Payload)
}

func (o *PutMetricsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutMetricsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

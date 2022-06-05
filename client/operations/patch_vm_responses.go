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

	models "titan/lib/firecracker/client/models"
)

// PatchVMReader is a Reader for the PatchVM structure.
type PatchVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPatchVMNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchVMBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchVMDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchVMNoContent creates a PatchVMNoContent with default headers values
func NewPatchVMNoContent() *PatchVMNoContent {
	return &PatchVMNoContent{}
}

/*PatchVMNoContent handles this case with default header values.

Vm state updated
*/
type PatchVMNoContent struct {
}

func (o *PatchVMNoContent) Error() string {
	return fmt.Sprintf("[PATCH /vm][%d] patchVmNoContent ", 204)
}

func (o *PatchVMNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchVMBadRequest creates a PatchVMBadRequest with default headers values
func NewPatchVMBadRequest() *PatchVMBadRequest {
	return &PatchVMBadRequest{}
}

/*PatchVMBadRequest handles this case with default header values.

Vm state cannot be updated due to bad input
*/
type PatchVMBadRequest struct {
	Payload *models.Error
}

func (o *PatchVMBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /vm][%d] patchVmBadRequest  %+v", 400, o.Payload)
}

func (o *PatchVMBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchVMBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVMDefault creates a PatchVMDefault with default headers values
func NewPatchVMDefault(code int) *PatchVMDefault {
	return &PatchVMDefault{
		_statusCode: code,
	}
}

/*PatchVMDefault handles this case with default header values.

Internal server error
*/
type PatchVMDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch Vm default response
func (o *PatchVMDefault) Code() int {
	return o._statusCode
}

func (o *PatchVMDefault) Error() string {
	return fmt.Sprintf("[PATCH /vm][%d] patchVm default  %+v", o._statusCode, o.Payload)
}

func (o *PatchVMDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchVMDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

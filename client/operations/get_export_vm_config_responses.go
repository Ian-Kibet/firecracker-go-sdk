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

	models "github.com/Ian-Kibet/firecracker-go-sdk/client/models"
)

// GetExportVMConfigReader is a Reader for the GetExportVMConfig structure.
type GetExportVMConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExportVMConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExportVMConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetExportVMConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetExportVMConfigOK creates a GetExportVMConfigOK with default headers values
func NewGetExportVMConfigOK() *GetExportVMConfigOK {
	return &GetExportVMConfigOK{}
}

/*GetExportVMConfigOK handles this case with default header values.

OK
*/
type GetExportVMConfigOK struct {
	Payload *models.FullVMConfiguration
}

func (o *GetExportVMConfigOK) Error() string {
	return fmt.Sprintf("[GET /vm/config][%d] getExportVmConfigOK  %+v", 200, o.Payload)
}

func (o *GetExportVMConfigOK) GetPayload() *models.FullVMConfiguration {
	return o.Payload
}

func (o *GetExportVMConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FullVMConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExportVMConfigDefault creates a GetExportVMConfigDefault with default headers values
func NewGetExportVMConfigDefault(code int) *GetExportVMConfigDefault {
	return &GetExportVMConfigDefault{
		_statusCode: code,
	}
}

/*GetExportVMConfigDefault handles this case with default header values.

Internal server error
*/
type GetExportVMConfigDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get export Vm config default response
func (o *GetExportVMConfigDefault) Code() int {
	return o._statusCode
}

func (o *GetExportVMConfigDefault) Error() string {
	return fmt.Sprintf("[GET /vm/config][%d] getExportVmConfig default  %+v", o._statusCode, o.Payload)
}

func (o *GetExportVMConfigDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetExportVMConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

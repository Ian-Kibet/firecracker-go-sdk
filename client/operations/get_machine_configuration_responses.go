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

// GetMachineConfigurationReader is a Reader for the GetMachineConfiguration structure.
type GetMachineConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMachineConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMachineConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMachineConfigurationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMachineConfigurationOK creates a GetMachineConfigurationOK with default headers values
func NewGetMachineConfigurationOK() *GetMachineConfigurationOK {
	return &GetMachineConfigurationOK{}
}

/*GetMachineConfigurationOK handles this case with default header values.

OK
*/
type GetMachineConfigurationOK struct {
	Payload *models.MachineConfiguration
}

func (o *GetMachineConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /machine-config][%d] getMachineConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetMachineConfigurationOK) GetPayload() *models.MachineConfiguration {
	return o.Payload
}

func (o *GetMachineConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MachineConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMachineConfigurationDefault creates a GetMachineConfigurationDefault with default headers values
func NewGetMachineConfigurationDefault(code int) *GetMachineConfigurationDefault {
	return &GetMachineConfigurationDefault{
		_statusCode: code,
	}
}

/*GetMachineConfigurationDefault handles this case with default header values.

Internal server error
*/
type GetMachineConfigurationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get machine configuration default response
func (o *GetMachineConfigurationDefault) Code() int {
	return o._statusCode
}

func (o *GetMachineConfigurationDefault) Error() string {
	return fmt.Sprintf("[GET /machine-config][%d] getMachineConfiguration default  %+v", o._statusCode, o.Payload)
}

func (o *GetMachineConfigurationDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetMachineConfigurationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// GetMmdsReader is a Reader for the GetMmds structure.
type GetMmdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMmdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMmdsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetMmdsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetMmdsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMmdsOK creates a GetMmdsOK with default headers values
func NewGetMmdsOK() *GetMmdsOK {
	return &GetMmdsOK{}
}

/*GetMmdsOK handles this case with default header values.

The MMDS data store JSON.
*/
type GetMmdsOK struct {
	Payload interface{}
}

func (o *GetMmdsOK) Error() string {
	return fmt.Sprintf("[GET /mmds][%d] getMmdsOK  %+v", 200, o.Payload)
}

func (o *GetMmdsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetMmdsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMmdsNotFound creates a GetMmdsNotFound with default headers values
func NewGetMmdsNotFound() *GetMmdsNotFound {
	return &GetMmdsNotFound{}
}

/*GetMmdsNotFound handles this case with default header values.

The MMDS data store content can not be found.
*/
type GetMmdsNotFound struct {
	Payload *models.Error
}

func (o *GetMmdsNotFound) Error() string {
	return fmt.Sprintf("[GET /mmds][%d] getMmdsNotFound  %+v", 404, o.Payload)
}

func (o *GetMmdsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetMmdsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMmdsDefault creates a GetMmdsDefault with default headers values
func NewGetMmdsDefault(code int) *GetMmdsDefault {
	return &GetMmdsDefault{
		_statusCode: code,
	}
}

/*GetMmdsDefault handles this case with default header values.

Internal server error
*/
type GetMmdsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get mmds default response
func (o *GetMmdsDefault) Code() int {
	return o._statusCode
}

func (o *GetMmdsDefault) Error() string {
	return fmt.Sprintf("[GET /mmds][%d] getMmds default  %+v", o._statusCode, o.Payload)
}

func (o *GetMmdsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetMmdsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

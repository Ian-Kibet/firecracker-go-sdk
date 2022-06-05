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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "titan/lib/firecracker/client/models"
)

// NewPutGuestVsockParams creates a new PutGuestVsockParams object
// with the default values initialized.
func NewPutGuestVsockParams() *PutGuestVsockParams {
	var ()
	return &PutGuestVsockParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutGuestVsockParamsWithTimeout creates a new PutGuestVsockParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutGuestVsockParamsWithTimeout(timeout time.Duration) *PutGuestVsockParams {
	var ()
	return &PutGuestVsockParams{

		timeout: timeout,
	}
}

// NewPutGuestVsockParamsWithContext creates a new PutGuestVsockParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutGuestVsockParamsWithContext(ctx context.Context) *PutGuestVsockParams {
	var ()
	return &PutGuestVsockParams{

		Context: ctx,
	}
}

// NewPutGuestVsockParamsWithHTTPClient creates a new PutGuestVsockParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutGuestVsockParamsWithHTTPClient(client *http.Client) *PutGuestVsockParams {
	var ()
	return &PutGuestVsockParams{
		HTTPClient: client,
	}
}

/*PutGuestVsockParams contains all the parameters to send to the API endpoint
for the put guest vsock operation typically these are written to a http.Request
*/
type PutGuestVsockParams struct {

	/*Body
	  Guest vsock properties

	*/
	Body *models.Vsock

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put guest vsock params
func (o *PutGuestVsockParams) WithTimeout(timeout time.Duration) *PutGuestVsockParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put guest vsock params
func (o *PutGuestVsockParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put guest vsock params
func (o *PutGuestVsockParams) WithContext(ctx context.Context) *PutGuestVsockParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put guest vsock params
func (o *PutGuestVsockParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put guest vsock params
func (o *PutGuestVsockParams) WithHTTPClient(client *http.Client) *PutGuestVsockParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put guest vsock params
func (o *PutGuestVsockParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put guest vsock params
func (o *PutGuestVsockParams) WithBody(body *models.Vsock) *PutGuestVsockParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put guest vsock params
func (o *PutGuestVsockParams) SetBody(body *models.Vsock) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutGuestVsockParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

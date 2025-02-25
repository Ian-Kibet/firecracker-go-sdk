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

<<<<<<< HEAD
	models "titan/lib/firecracker/client/models"
=======
	models "github.com/Ian-Kibet/firecracker-go-sdk/client/models"
>>>>>>> b8aa219df3977843c18fb0ce7af8af072b1bf0b8
)

// NewCreateSyncActionParams creates a new CreateSyncActionParams object
// with the default values initialized.
func NewCreateSyncActionParams() *CreateSyncActionParams {
	var ()
	return &CreateSyncActionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSyncActionParamsWithTimeout creates a new CreateSyncActionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSyncActionParamsWithTimeout(timeout time.Duration) *CreateSyncActionParams {
	var ()
	return &CreateSyncActionParams{

		timeout: timeout,
	}
}

// NewCreateSyncActionParamsWithContext creates a new CreateSyncActionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSyncActionParamsWithContext(ctx context.Context) *CreateSyncActionParams {
	var ()
	return &CreateSyncActionParams{

		Context: ctx,
	}
}

// NewCreateSyncActionParamsWithHTTPClient creates a new CreateSyncActionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSyncActionParamsWithHTTPClient(client *http.Client) *CreateSyncActionParams {
	var ()
	return &CreateSyncActionParams{
		HTTPClient: client,
	}
}

/*CreateSyncActionParams contains all the parameters to send to the API endpoint
for the create sync action operation typically these are written to a http.Request
*/
type CreateSyncActionParams struct {

	/*Info*/
	Info *models.InstanceActionInfo

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create sync action params
func (o *CreateSyncActionParams) WithTimeout(timeout time.Duration) *CreateSyncActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create sync action params
func (o *CreateSyncActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create sync action params
func (o *CreateSyncActionParams) WithContext(ctx context.Context) *CreateSyncActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create sync action params
func (o *CreateSyncActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create sync action params
func (o *CreateSyncActionParams) WithHTTPClient(client *http.Client) *CreateSyncActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create sync action params
func (o *CreateSyncActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the create sync action params
func (o *CreateSyncActionParams) WithInfo(info *models.InstanceActionInfo) *CreateSyncActionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the create sync action params
func (o *CreateSyncActionParams) SetInfo(info *models.InstanceActionInfo) {
	o.Info = info
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSyncActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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

// NewPatchBalloonStatsIntervalParams creates a new PatchBalloonStatsIntervalParams object
// with the default values initialized.
func NewPatchBalloonStatsIntervalParams() *PatchBalloonStatsIntervalParams {
	var ()
	return &PatchBalloonStatsIntervalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchBalloonStatsIntervalParamsWithTimeout creates a new PatchBalloonStatsIntervalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchBalloonStatsIntervalParamsWithTimeout(timeout time.Duration) *PatchBalloonStatsIntervalParams {
	var ()
	return &PatchBalloonStatsIntervalParams{

		timeout: timeout,
	}
}

// NewPatchBalloonStatsIntervalParamsWithContext creates a new PatchBalloonStatsIntervalParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchBalloonStatsIntervalParamsWithContext(ctx context.Context) *PatchBalloonStatsIntervalParams {
	var ()
	return &PatchBalloonStatsIntervalParams{

		Context: ctx,
	}
}

// NewPatchBalloonStatsIntervalParamsWithHTTPClient creates a new PatchBalloonStatsIntervalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchBalloonStatsIntervalParamsWithHTTPClient(client *http.Client) *PatchBalloonStatsIntervalParams {
	var ()
	return &PatchBalloonStatsIntervalParams{
		HTTPClient: client,
	}
}

/*PatchBalloonStatsIntervalParams contains all the parameters to send to the API endpoint
for the patch balloon stats interval operation typically these are written to a http.Request
*/
type PatchBalloonStatsIntervalParams struct {

	/*Body
	  Balloon properties

	*/
	Body *models.BalloonStatsUpdate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch balloon stats interval params
func (o *PatchBalloonStatsIntervalParams) WithTimeout(timeout time.Duration) *PatchBalloonStatsIntervalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch balloon stats interval params
func (o *PatchBalloonStatsIntervalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch balloon stats interval params
func (o *PatchBalloonStatsIntervalParams) WithContext(ctx context.Context) *PatchBalloonStatsIntervalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch balloon stats interval params
func (o *PatchBalloonStatsIntervalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch balloon stats interval params
func (o *PatchBalloonStatsIntervalParams) WithHTTPClient(client *http.Client) *PatchBalloonStatsIntervalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch balloon stats interval params
func (o *PatchBalloonStatsIntervalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch balloon stats interval params
func (o *PatchBalloonStatsIntervalParams) WithBody(body *models.BalloonStatsUpdate) *PatchBalloonStatsIntervalParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch balloon stats interval params
func (o *PatchBalloonStatsIntervalParams) SetBody(body *models.BalloonStatsUpdate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PatchBalloonStatsIntervalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

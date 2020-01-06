// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetNetworkAclsNetworkACLIDRulesIDParams creates a new GetNetworkAclsNetworkACLIDRulesIDParams object
// with the default values initialized.
func NewGetNetworkAclsNetworkACLIDRulesIDParams() *GetNetworkAclsNetworkACLIDRulesIDParams {
	var ()
	return &GetNetworkAclsNetworkACLIDRulesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkAclsNetworkACLIDRulesIDParamsWithTimeout creates a new GetNetworkAclsNetworkACLIDRulesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkAclsNetworkACLIDRulesIDParamsWithTimeout(timeout time.Duration) *GetNetworkAclsNetworkACLIDRulesIDParams {
	var ()
	return &GetNetworkAclsNetworkACLIDRulesIDParams{

		timeout: timeout,
	}
}

// NewGetNetworkAclsNetworkACLIDRulesIDParamsWithContext creates a new GetNetworkAclsNetworkACLIDRulesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkAclsNetworkACLIDRulesIDParamsWithContext(ctx context.Context) *GetNetworkAclsNetworkACLIDRulesIDParams {
	var ()
	return &GetNetworkAclsNetworkACLIDRulesIDParams{

		Context: ctx,
	}
}

// NewGetNetworkAclsNetworkACLIDRulesIDParamsWithHTTPClient creates a new GetNetworkAclsNetworkACLIDRulesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkAclsNetworkACLIDRulesIDParamsWithHTTPClient(client *http.Client) *GetNetworkAclsNetworkACLIDRulesIDParams {
	var ()
	return &GetNetworkAclsNetworkACLIDRulesIDParams{
		HTTPClient: client,
	}
}

/*GetNetworkAclsNetworkACLIDRulesIDParams contains all the parameters to send to the API endpoint
for the get network acls network ACL ID rules ID operation typically these are written to a http.Request
*/
type GetNetworkAclsNetworkACLIDRulesIDParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*ID
	  The rule identifier

	*/
	ID string
	/*NetworkACLID
	  The network ACL identifier

	*/
	NetworkACLID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network acls network ACL ID rules ID params
func (o *GetNetworkAclsNetworkACLIDRulesIDParams) WithTimeout(timeout time.Duration) *GetNetworkAclsNetworkACLIDRulesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network acls network ACL ID rules ID params
func (o *GetNetworkAclsNetworkACLIDRulesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network acls network ACL ID rules ID params
func (o *GetNetworkAclsNetworkACLIDRulesIDParams) WithContext(ctx context.Context) *GetNetworkAclsNetworkACLIDRulesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network acls network ACL ID rules ID params
func (o *GetNetworkAclsNetworkACLIDRulesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network acls network ACL ID rules ID params
func (o *GetNetworkAclsNetworkACLIDRulesIDParams) WithHTTPClient(client *http.Client) *GetNetworkAclsNetworkACLIDRulesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network acls network ACL ID rules ID params
func (o *GetNetworkAclsNetworkACLIDRulesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the get network acls network ACL ID rules ID params
func (o *GetNetworkAclsNetworkACLIDRulesIDParams) WithGeneration(generation int64) *GetNetworkAclsNetworkACLIDRulesIDParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the get network acls network ACL ID rules ID params
func (o *GetNetworkAclsNetworkACLIDRulesIDParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the get network acls network ACL ID rules ID params
func (o *GetNetworkAclsNetworkACLIDRulesIDParams) WithID(id string) *GetNetworkAclsNetworkACLIDRulesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get network acls network ACL ID rules ID params
func (o *GetNetworkAclsNetworkACLIDRulesIDParams) SetID(id string) {
	o.ID = id
}

// WithNetworkACLID adds the networkACLID to the get network acls network ACL ID rules ID params
func (o *GetNetworkAclsNetworkACLIDRulesIDParams) WithNetworkACLID(networkACLID string) *GetNetworkAclsNetworkACLIDRulesIDParams {
	o.SetNetworkACLID(networkACLID)
	return o
}

// SetNetworkACLID adds the networkAclId to the get network acls network ACL ID rules ID params
func (o *GetNetworkAclsNetworkACLIDRulesIDParams) SetNetworkACLID(networkACLID string) {
	o.NetworkACLID = networkACLID
}

// WithVersion adds the version to the get network acls network ACL ID rules ID params
func (o *GetNetworkAclsNetworkACLIDRulesIDParams) WithVersion(version string) *GetNetworkAclsNetworkACLIDRulesIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get network acls network ACL ID rules ID params
func (o *GetNetworkAclsNetworkACLIDRulesIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkAclsNetworkACLIDRulesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param network_acl_id
	if err := r.SetPathParam("network_acl_id", o.NetworkACLID); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
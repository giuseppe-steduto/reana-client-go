// Code generated by go-swagger; DO NOT EDIT.

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
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetWorkflowsParams creates a new GetWorkflowsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWorkflowsParams() *GetWorkflowsParams {
	return &GetWorkflowsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkflowsParamsWithTimeout creates a new GetWorkflowsParams object
// with the ability to set a timeout on a request.
func NewGetWorkflowsParamsWithTimeout(timeout time.Duration) *GetWorkflowsParams {
	return &GetWorkflowsParams{
		timeout: timeout,
	}
}

// NewGetWorkflowsParamsWithContext creates a new GetWorkflowsParams object
// with the ability to set a context for a request.
func NewGetWorkflowsParamsWithContext(ctx context.Context) *GetWorkflowsParams {
	return &GetWorkflowsParams{
		Context: ctx,
	}
}

// NewGetWorkflowsParamsWithHTTPClient creates a new GetWorkflowsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWorkflowsParamsWithHTTPClient(client *http.Client) *GetWorkflowsParams {
	return &GetWorkflowsParams{
		HTTPClient: client,
	}
}

/*
GetWorkflowsParams contains all the parameters to send to the API endpoint

	for the get workflows operation.

	Typically these are written to a http.Request.
*/
type GetWorkflowsParams struct {

	/* AccessToken.

	   The API access_token of workflow owner.
	*/
	AccessToken *string

	/* IncludeLastCommand.

	   Include information about the current command.
	*/
	IncludeLastCommand *bool

	/* IncludeProgress.

	   Include progress information of the workflows.
	*/
	IncludeProgress *bool

	/* IncludeWorkspaceSize.

	   Include size information of the workspace.
	*/
	IncludeWorkspaceSize *bool

	/* Page.

	   Results page number (pagination).
	*/
	Page *int64

	/* Search.

	   Filter workflows by name.
	*/
	Search *string

	/* Size.

	   Number of results per page (pagination).
	*/
	Size *int64

	/* Sort.

	   Sort workflows by creation date (asc, desc).
	*/
	Sort *string

	/* Status.

	   Filter workflows by list of statuses.
	*/
	Status []string

	/* Type.

	   Required. Type of workflows.
	*/
	Type string

	/* Verbose.

	   Optional flag to show more information.
	*/
	Verbose *bool

	/* WorkflowIDOrName.

	   Optional analysis UUID or name to filter.
	*/
	WorkflowIDOrName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get workflows params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkflowsParams) WithDefaults() *GetWorkflowsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get workflows params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkflowsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get workflows params
func (o *GetWorkflowsParams) WithTimeout(timeout time.Duration) *GetWorkflowsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workflows params
func (o *GetWorkflowsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workflows params
func (o *GetWorkflowsParams) WithContext(ctx context.Context) *GetWorkflowsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workflows params
func (o *GetWorkflowsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workflows params
func (o *GetWorkflowsParams) WithHTTPClient(client *http.Client) *GetWorkflowsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workflows params
func (o *GetWorkflowsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessToken adds the accessToken to the get workflows params
func (o *GetWorkflowsParams) WithAccessToken(accessToken *string) *GetWorkflowsParams {
	o.SetAccessToken(accessToken)
	return o
}

// SetAccessToken adds the accessToken to the get workflows params
func (o *GetWorkflowsParams) SetAccessToken(accessToken *string) {
	o.AccessToken = accessToken
}

// WithIncludeLastCommand adds the includeLastCommand to the get workflows params
func (o *GetWorkflowsParams) WithIncludeLastCommand(includeLastCommand *bool) *GetWorkflowsParams {
	o.SetIncludeLastCommand(includeLastCommand)
	return o
}

// SetIncludeLastCommand adds the includeLastCommand to the get workflows params
func (o *GetWorkflowsParams) SetIncludeLastCommand(includeLastCommand *bool) {
	o.IncludeLastCommand = includeLastCommand
}

// WithIncludeProgress adds the includeProgress to the get workflows params
func (o *GetWorkflowsParams) WithIncludeProgress(includeProgress *bool) *GetWorkflowsParams {
	o.SetIncludeProgress(includeProgress)
	return o
}

// SetIncludeProgress adds the includeProgress to the get workflows params
func (o *GetWorkflowsParams) SetIncludeProgress(includeProgress *bool) {
	o.IncludeProgress = includeProgress
}

// WithIncludeWorkspaceSize adds the includeWorkspaceSize to the get workflows params
func (o *GetWorkflowsParams) WithIncludeWorkspaceSize(includeWorkspaceSize *bool) *GetWorkflowsParams {
	o.SetIncludeWorkspaceSize(includeWorkspaceSize)
	return o
}

// SetIncludeWorkspaceSize adds the includeWorkspaceSize to the get workflows params
func (o *GetWorkflowsParams) SetIncludeWorkspaceSize(includeWorkspaceSize *bool) {
	o.IncludeWorkspaceSize = includeWorkspaceSize
}

// WithPage adds the page to the get workflows params
func (o *GetWorkflowsParams) WithPage(page *int64) *GetWorkflowsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get workflows params
func (o *GetWorkflowsParams) SetPage(page *int64) {
	o.Page = page
}

// WithSearch adds the search to the get workflows params
func (o *GetWorkflowsParams) WithSearch(search *string) *GetWorkflowsParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get workflows params
func (o *GetWorkflowsParams) SetSearch(search *string) {
	o.Search = search
}

// WithSize adds the size to the get workflows params
func (o *GetWorkflowsParams) WithSize(size *int64) *GetWorkflowsParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get workflows params
func (o *GetWorkflowsParams) SetSize(size *int64) {
	o.Size = size
}

// WithSort adds the sort to the get workflows params
func (o *GetWorkflowsParams) WithSort(sort *string) *GetWorkflowsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get workflows params
func (o *GetWorkflowsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithStatus adds the status to the get workflows params
func (o *GetWorkflowsParams) WithStatus(status []string) *GetWorkflowsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get workflows params
func (o *GetWorkflowsParams) SetStatus(status []string) {
	o.Status = status
}

// WithType adds the typeVar to the get workflows params
func (o *GetWorkflowsParams) WithType(typeVar string) *GetWorkflowsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get workflows params
func (o *GetWorkflowsParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WithVerbose adds the verbose to the get workflows params
func (o *GetWorkflowsParams) WithVerbose(verbose *bool) *GetWorkflowsParams {
	o.SetVerbose(verbose)
	return o
}

// SetVerbose adds the verbose to the get workflows params
func (o *GetWorkflowsParams) SetVerbose(verbose *bool) {
	o.Verbose = verbose
}

// WithWorkflowIDOrName adds the workflowIDOrName to the get workflows params
func (o *GetWorkflowsParams) WithWorkflowIDOrName(workflowIDOrName *string) *GetWorkflowsParams {
	o.SetWorkflowIDOrName(workflowIDOrName)
	return o
}

// SetWorkflowIDOrName adds the workflowIdOrName to the get workflows params
func (o *GetWorkflowsParams) SetWorkflowIDOrName(workflowIDOrName *string) {
	o.WorkflowIDOrName = workflowIDOrName
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkflowsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccessToken != nil {

		// query param access_token
		var qrAccessToken string

		if o.AccessToken != nil {
			qrAccessToken = *o.AccessToken
		}
		qAccessToken := qrAccessToken
		if qAccessToken != "" {

			if err := r.SetQueryParam("access_token", qAccessToken); err != nil {
				return err
			}
		}
	}

	if o.IncludeLastCommand != nil {

		// query param include_last_command
		var qrIncludeLastCommand bool

		if o.IncludeLastCommand != nil {
			qrIncludeLastCommand = *o.IncludeLastCommand
		}
		qIncludeLastCommand := swag.FormatBool(qrIncludeLastCommand)
		if qIncludeLastCommand != "" {

			if err := r.SetQueryParam("include_last_command", qIncludeLastCommand); err != nil {
				return err
			}
		}
	}

	if o.IncludeProgress != nil {

		// query param include_progress
		var qrIncludeProgress bool

		if o.IncludeProgress != nil {
			qrIncludeProgress = *o.IncludeProgress
		}
		qIncludeProgress := swag.FormatBool(qrIncludeProgress)
		if qIncludeProgress != "" {

			if err := r.SetQueryParam("include_progress", qIncludeProgress); err != nil {
				return err
			}
		}
	}

	if o.IncludeWorkspaceSize != nil {

		// query param include_workspace_size
		var qrIncludeWorkspaceSize bool

		if o.IncludeWorkspaceSize != nil {
			qrIncludeWorkspaceSize = *o.IncludeWorkspaceSize
		}
		qIncludeWorkspaceSize := swag.FormatBool(qrIncludeWorkspaceSize)
		if qIncludeWorkspaceSize != "" {

			if err := r.SetQueryParam("include_workspace_size", qIncludeWorkspaceSize); err != nil {
				return err
			}
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.Search != nil {

		// query param search
		var qrSearch string

		if o.Search != nil {
			qrSearch = *o.Search
		}
		qSearch := qrSearch
		if qSearch != "" {

			if err := r.SetQueryParam("search", qSearch); err != nil {
				return err
			}
		}
	}

	if o.Size != nil {

		// query param size
		var qrSize int64

		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt64(qrSize)
		if qSize != "" {

			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if o.Status != nil {

		// binding items for status
		joinedStatus := o.bindParamStatus(reg)

		// query array param status
		if err := r.SetQueryParam("status", joinedStatus...); err != nil {
			return err
		}
	}

	// query param type
	qrType := o.Type
	qType := qrType
	if qType != "" {

		if err := r.SetQueryParam("type", qType); err != nil {
			return err
		}
	}

	if o.Verbose != nil {

		// query param verbose
		var qrVerbose bool

		if o.Verbose != nil {
			qrVerbose = *o.Verbose
		}
		qVerbose := swag.FormatBool(qrVerbose)
		if qVerbose != "" {

			if err := r.SetQueryParam("verbose", qVerbose); err != nil {
				return err
			}
		}
	}

	if o.WorkflowIDOrName != nil {

		// query param workflow_id_or_name
		var qrWorkflowIDOrName string

		if o.WorkflowIDOrName != nil {
			qrWorkflowIDOrName = *o.WorkflowIDOrName
		}
		qWorkflowIDOrName := qrWorkflowIDOrName
		if qWorkflowIDOrName != "" {

			if err := r.SetQueryParam("workflow_id_or_name", qWorkflowIDOrName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetWorkflows binds the parameter status
func (o *GetWorkflowsParams) bindParamStatus(formats strfmt.Registry) []string {
	statusIR := o.Status

	var statusIC []string
	for _, statusIIR := range statusIR { // explode []string

		statusIIV := statusIIR // string as string
		statusIC = append(statusIC, statusIIV)
	}

	// items.CollectionFormat: ""
	statusIS := swag.JoinByFormat(statusIC, "")

	return statusIS
}

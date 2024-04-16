// Code generated by go-swagger; DO NOT EDIT.

package recurring_run_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new recurring run service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for recurring run service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
RecurringRunServiceCreateRecurringRun creates a new recurring run in an experiment given the experiment ID
*/
func (a *Client) RecurringRunServiceCreateRecurringRun(params *RecurringRunServiceCreateRecurringRunParams) (*RecurringRunServiceCreateRecurringRunOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRecurringRunServiceCreateRecurringRunParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RecurringRunService_CreateRecurringRun",
		Method:             "POST",
		PathPattern:        "/apis/v2beta1/recurringruns",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RecurringRunServiceCreateRecurringRunReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RecurringRunServiceCreateRecurringRunOK), nil

}

/*
RecurringRunServiceDeleteRecurringRun deletes a recurring run
*/
func (a *Client) RecurringRunServiceDeleteRecurringRun(params *RecurringRunServiceDeleteRecurringRunParams) (*RecurringRunServiceDeleteRecurringRunOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRecurringRunServiceDeleteRecurringRunParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RecurringRunService_DeleteRecurringRun",
		Method:             "DELETE",
		PathPattern:        "/apis/v2beta1/recurringruns/{recurring_run_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RecurringRunServiceDeleteRecurringRunReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RecurringRunServiceDeleteRecurringRunOK), nil

}

/*
RecurringRunServiceDisableRecurringRun stops a recurring run and all its associated runs the recurring run is not deleted
*/
func (a *Client) RecurringRunServiceDisableRecurringRun(params *RecurringRunServiceDisableRecurringRunParams) (*RecurringRunServiceDisableRecurringRunOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRecurringRunServiceDisableRecurringRunParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RecurringRunService_DisableRecurringRun",
		Method:             "POST",
		PathPattern:        "/apis/v2beta1/recurringruns/{recurring_run_id}:disable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RecurringRunServiceDisableRecurringRunReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RecurringRunServiceDisableRecurringRunOK), nil

}

/*
RecurringRunServiceEnableRecurringRun restarts a recurring run that was previously stopped all runs associated with the recurring run will continue
*/
func (a *Client) RecurringRunServiceEnableRecurringRun(params *RecurringRunServiceEnableRecurringRunParams) (*RecurringRunServiceEnableRecurringRunOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRecurringRunServiceEnableRecurringRunParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RecurringRunService_EnableRecurringRun",
		Method:             "POST",
		PathPattern:        "/apis/v2beta1/recurringruns/{recurring_run_id}:enable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RecurringRunServiceEnableRecurringRunReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RecurringRunServiceEnableRecurringRunOK), nil

}

/*
RecurringRunServiceGetRecurringRun finds a specific recurring run by ID
*/
func (a *Client) RecurringRunServiceGetRecurringRun(params *RecurringRunServiceGetRecurringRunParams) (*RecurringRunServiceGetRecurringRunOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRecurringRunServiceGetRecurringRunParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RecurringRunService_GetRecurringRun",
		Method:             "GET",
		PathPattern:        "/apis/v2beta1/recurringruns/{recurring_run_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RecurringRunServiceGetRecurringRunReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RecurringRunServiceGetRecurringRunOK), nil

}

/*
RecurringRunServiceListRecurringRuns finds all recurring runs given experiment and namespace if experiment ID is not specified find all recurring runs across all experiments
*/
func (a *Client) RecurringRunServiceListRecurringRuns(params *RecurringRunServiceListRecurringRunsParams) (*RecurringRunServiceListRecurringRunsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRecurringRunServiceListRecurringRunsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RecurringRunService_ListRecurringRuns",
		Method:             "GET",
		PathPattern:        "/apis/v2beta1/recurringruns",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RecurringRunServiceListRecurringRunsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RecurringRunServiceListRecurringRunsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

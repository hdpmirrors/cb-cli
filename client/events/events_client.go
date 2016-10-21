package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new events API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for events API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*
GetEvents retrieves events by timestamp long

Events are used to track stack creation initiated by cloudbreak users. Events are generated by the backend when resources requested by the user become available or unavailable
*/
func (a *Client) GetEvents(params *GetEventsParams) (*GetEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEventsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "GetEvents",
		Method:             "GET",
		PathPattern:        "/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEventsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEventsOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport client.Transport) {
	a.transport = transport
}

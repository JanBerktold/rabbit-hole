package rabbithole

import (
	"encoding/json"
	"net/http"
	"fmt"
)

// Federation definition: additional arguments
// added to the entities (queues, exchanges or both)
// that match a policy.
type FederationDefinition struct {
	Uri            string `json:"uri"`
	Expires        int    `json:"expires"`
	MessageTTL     int32  `json:"message-ttl"`
	MaxHops        int    `json:"max-hops"`
	PrefetchCount  int    `json:"prefetch-count"`
	ReconnectDelay int    `json:"reconnect-delay"`
	AckMode        string `json:"ack-mode,omitempty"`
	TrustUserId    bool   `json:"trust-user-id"`
	Exchange       string `json:"exchange"`
	Queue          string `json:"queue"`
}

// Represents a configured Federation upstream.
type FederationUpstream struct {
	Name string `json:"name,omitempty"`
	VHost string `json:"vhost,omitempty"`
	Definition FederationDefinition `json:"value"`
}

//
// GET /api/parameters/federation-upstream/{vhost}/{name}
//

// Gets a federation upstream.
func (c * Client) GetFederationUpstream(vhost string, name string) (*FederationUpstream, error) {
	req, err := newRequestWithBody(c, "GET", fmt.Sprintf("parameters/federation-upstream/%s/%s", vhost, name) , nil)
	if err != nil {
		return nil, err
	}

	result := FederationUpstream{}

	err = executeAndParseRequest(c, req, &result)
	if err != nil {
		return nil, err
	}

	return &result, err
}

//
// GET /api/parameters/federation-upstream
//

// Lists all configured federation upstreams.
func (c * Client) ListFederationUpstreams() ([]FederationUpstream, error) {
	req, err := newRequestWithBody(c, "GET", "parameters/federation-upstream", nil)
	if err != nil {
		return nil, err
	}

	result := make([]FederationUpstream, 0)

	err = executeAndParseRequest(c, req, &result)
	if err != nil {
		return nil, err
	}

	return result, err
}

//
// PUT /api/parameters/federation-upstream/{vhost}/{upstream}
//

// Updates a federation upstream
func (c *Client) PutFederationUpstream(vhost string, upstreamName string, fDef FederationDefinition) (res *http.Response, err error) {
	fedUp := FederationUpstream{
		Definition: fDef,
	}
	body, err := json.Marshal(fedUp)
	if err != nil {
		return nil, err
	}

	req, err := newRequestWithBody(c, "PUT", "parameters/federation-upstream/"+PathEscape(vhost)+"/"+PathEscape(upstreamName), body)
	if err != nil {
		return nil, err
	}

	res, err = executeRequest(c, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//
// DELETE /api/parameters/federation-upstream/{vhost}/{name}
//

// Deletes a federation upstream.
func (c *Client) DeleteFederationUpstream(vhost, upstreamName string) (res *http.Response, err error) {
	req, err := newRequestWithBody(c, "DELETE", "parameters/federation-upstream/"+PathEscape(vhost)+"/"+PathEscape(upstreamName), nil)
	if err != nil {
		return nil, err
	}

	res, err = executeRequest(c, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

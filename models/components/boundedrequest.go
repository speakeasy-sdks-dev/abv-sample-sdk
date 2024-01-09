// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks-dev/abv-sample-sdk/internal/utils"
	"time"
)

// A BoundedRequest is a request that has been logged by the Speakeasy without the contents of the request.
type BoundedRequest struct {
	// The ID of the ApiEndpoint this request was made to.
	APIEndpointID string `json:"api_endpoint_id"`
	// The ID of the Api this request was made to.
	APIID string `json:"api_id"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at"`
	// The ID of the customer that made this request.
	CustomerID string `json:"customer_id"`
	// The latency of the request.
	Latency int64 `json:"latency"`
	// A set of values associated with a metadata key.
	Metadata map[string][]string `json:"metadata,omitempty"`
	// HTTP verb.
	Method string `json:"method"`
	// The path of the request.
	Path string `json:"path"`
	// The time the request finished.
	RequestFinishTime time.Time `json:"request_finish_time"`
	// The ID of this request.
	RequestID string `json:"request_id"`
	// The time the request was made.
	RequestStartTime time.Time `json:"request_start_time"`
	// The status code of the request.
	Status int64 `json:"status"`
	// The version ID of the Api this request was made to.
	VersionID string `json:"version_id"`
	// The workspace ID this request was made to.
	WorkspaceID string `json:"workspace_id"`
}

func (b BoundedRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BoundedRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *BoundedRequest) GetAPIEndpointID() string {
	if o == nil {
		return ""
	}
	return o.APIEndpointID
}

func (o *BoundedRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *BoundedRequest) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *BoundedRequest) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

func (o *BoundedRequest) GetLatency() int64 {
	if o == nil {
		return 0
	}
	return o.Latency
}

func (o *BoundedRequest) GetMetadata() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *BoundedRequest) GetMethod() string {
	if o == nil {
		return ""
	}
	return o.Method
}

func (o *BoundedRequest) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

func (o *BoundedRequest) GetRequestFinishTime() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.RequestFinishTime
}

func (o *BoundedRequest) GetRequestID() string {
	if o == nil {
		return ""
	}
	return o.RequestID
}

func (o *BoundedRequest) GetRequestStartTime() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.RequestStartTime
}

func (o *BoundedRequest) GetStatus() int64 {
	if o == nil {
		return 0
	}
	return o.Status
}

func (o *BoundedRequest) GetVersionID() string {
	if o == nil {
		return ""
	}
	return o.VersionID
}

func (o *BoundedRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

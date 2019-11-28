package byob_sdk_go

import uuid "github.com/satori/go.uuid"

// RequestContext stores the context of a single byob request
type RequestContext struct {
	Flaggable
	ID              uuid.UUID              `json:"id"`
	User            RequestUser            `json:"user"`
	Session         Session                `json:"session"`
	EnvironmentData map[string]interface{} `json:"envData"`
	Text            string                 `json:"text"`
	OriginPlatform  string                 `json:"originPlatform"`
	OriginalRequest interface{}            `json:"originalRequest"`
	IsStart         bool                   `json:"isStart"`
	IsBroadcast     bool                   `json:"isBroadcast"`
	Errors          []ExecError            `json:"errors,omitempty"`
	Response        *Response              `json:"response,omitempty"`

	// LastError is the error that ocurred during execution of the last node
	LastError *ExecError `json:"lastError,omitempty"`
}

// Request user is holds the information for a user making a bot request
type RequestUser struct {
	Flaggable
	ID         uuid.UUID `json:"id"`
	PlatformID string    `json:"platformId"`
	Name       string    `json:"name"`
}

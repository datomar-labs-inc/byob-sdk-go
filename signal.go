package byob

import (
	uuid "github.com/satori/go.uuid"
)

// BrainSignal is the collection of information that the brain requires to process a request
type BrainSignal struct {
	EnvironmentID   uuid.UUID       `json:"environmentId"`
	ContextModifier ContextModifier `json:"contextModifier"`
	PlatformID      string          `json:"platformId"`
	OriginPlatform  string          `json:"originPlatform"`
	RawRequest      interface{}     `json:"rawRequest"`
}

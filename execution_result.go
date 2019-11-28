package byob

import (
	"time"
)

type ExecutionResult struct {
	FinalContext  RequestContext `json:"finalContext"`
	ExecutionLogs []ExecutionLog `json:"execLogs"`
	ExecutionTime time.Duration  `json:"execTime"`
}

// ExecutionLog logs a single execution of a node or link
type ExecutionLog struct {
	GraphID         int64            `json:"gid"`
	NodeID          *int64           `json:"nid,omitempty"`
	LinkSourceID    *int64           `json:"lsid,omitempty"`
	LinkDestID      *int64           `json:"ldid,omitempty"`
	ExecutionType   int16            `json:"ety"`
	ExecutionTime   int64            `json:"eti"`
	ContextModifier *ContextModifier `json:"cmo,omitempty"`
}

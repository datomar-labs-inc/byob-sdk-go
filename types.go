package byob

type ReachableUserResult struct {
	Count uint64 `json:"count"`
}

type BroadcastInput struct {
	ContextModifier *ContextModifier `json:"contextModifier,omitempty"`
	Platform        string           `json:"platform" binding:"required"`
	UserQuery       UserQuery        `json:"userQuery" binding:"required"`
}

type BroadcastResult struct {
	Status string `json:"status"`
	Users  int    `json:"users"`
}

const (
	UQEquals = iota
	UQExists
	UQNotExists
	UQNotEquals
	UQStartsWith
	UQGreaterThan
	UQLessThan
)

const (
	UQMAny = iota
	UQMAll
	UQMNone
)

type QueryCheck struct {
	Field     string   `json:"field"`
	Operation int64    `json:"operation"`
	Values    []string `json:"values"`
}

type UserQuery struct {
	Checks []QueryCheck `json:"checks"`
	Mode   int64        `json:"mode"`
}


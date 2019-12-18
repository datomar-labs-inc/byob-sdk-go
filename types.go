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

type UpdateUserDataInput struct {
	// A list of keys and their values to be set on the user data
	// If the key exists it will be overwritten
	Set    map[string]interface{} `json:"set"`

	// A list of keys to be remove from the user data, if the keys did not exist, nothing happens
	Delete []string               `json:"delete"`
}

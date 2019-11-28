package byob_sdk_go

type WebhookRequest struct {
	Name    string          `json:"name"`
	Context *RequestContext `json:"ctx"`
}

type WebhookHandler func(name string, ctx *RequestContext) (*ContextModifier, error)

package byob

type Response struct {
	Messages []Message `json:"messages" mapstructure:"messages"`
}

type ResponseConfig struct {
	Basic           []MessageConfig            `json:"messages" mapstructure:"messages" msgpack:"messages"`
	SendNow         bool                       `json:"sendNow" mapstructure:"sendNow"`
	CustomResponses map[string][]MessageConfig `json:"customResponses" mapstructure:"customResponses"`
}

type Message struct {
	Text        string          `json:"text" mapstructure:"text"`
	TypingTime  float64         `json:"typingTime" mapstructure:"typingTime"`
	ShouldBatch bool            `json:"shouldBatch" mapstructure:"shouldBatch"`
	GraphID     *int64          `json:"graphId,omitempty" mapstructure:"graphId"`
	NodeID      *int64          `json:"nodeId,omitempty" mapstructure:"nodeId"`
	Blocks      []ResponseBlock `json:"blocks,omitempty" mapstructure:"blocks"`
}

type MessageConfig struct {
	Text       string  `json:"text" mapstructure:"text"`
	TypingTime float64 `json:"typingTime" mapstructure:"typingTime"`
}


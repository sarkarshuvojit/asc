package types

type BlockingRequestDto struct {
	RequestTopic  string      `json:"requestTopic"`
	ResponseTopic string      `json:"responseTopic"`
	Payload       interface{} `json:"payload"`
	Brokers       []string    `json:"brokers"`
}

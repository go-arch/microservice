package models


//rebbitmq data send recieve model
type MQPublish struct {
	EventType string `json:"event_type"`
	Content interface{} `json:"content"`
}


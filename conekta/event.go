package conekta

// Event conekta Event callback
type Event struct {
	ID            string      `json:"id,omitempty"`
	Object        string      `json:"object,omitempty"`
	Livemode      *bool       `json:"livemode,omitempty"`
	CreatedAt     int64       `json:"created_at,omitempty"`
	Type          string      `json:"type,omitempty"`
	Data          EventData   `json:"data,omitempty"`
	WebhookStatus string      `json:"webhook_status,omitempty"`
	WebhookLogs   WebhookLogs `json:"webhook_logs,omitempty"`
}

// EventData conekta event callback
type EventData struct {
	ConektaResponse
	Description   string        `json:"description,omitempty"`
	PaymentMethod PaymentMethod `json:"payment_method,omitempty"`
}

// WebhookLogs conekta webhooklogs
type WebhookLogs struct {
	ID                     string `json:"id,omitempty"`
	Url                    string `json:"url,omitempty"`
	FailedAttempts         int    `json:"failed_attempts,omitempty"`
	LastHttpResponseStatus int    `json:"last_http_response_status,omitempty"`
	Object                 string `json:"object,omitempty"`
	LastAttemptedAt        int64  `json:"last_attempted_at,omitempty"`
}

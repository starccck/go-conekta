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

// EventType conekta webhook event type
type EventType = string

// various kinds of conekta webhook event type
const (
	OrderCreated      EventType = "order.created"
	OrderPaid                   = "order.paid"
	OrderExpiredOxxO            = "order.expired"
	ChargeCreatedOxxO           = "charge.created(oxxo)"
	ChargeCreatedCard           = "charge.created"
	ChargePaidOxxO              = "charge.paid"
	ChargePaidCard              = "charge.paid(oxxo)"
	ChargeRefunded              = "charge.refunded"
	ChargebackCreated           = "charge.chargeback.created"
	PlanCreated                 = "plan.created"
	SubsCreated                 = "subscription.created"
	SubsPaid                    = "subscription.paid"
	SubsCanceled                = "subscription.canceled"
	SubsPaymentFailed           = "subscription.payment_failed"
)

// ConektaEventMap conekta webhook event callback type
var ConektaEventMap = map[string]EventType{
	"order.created":               OrderCreated,
	"order.paid":                  OrderPaid,
	"order.expired":               OrderExpiredOxxO,
	"charge.created(oxxo)":        ChargeCreatedOxxO,
	"charge.created":              ChargeCreatedCard,
	"charge.paid":                 ChargePaidOxxO,
	"charge.paid(oxxo)":           ChargePaidCard,
	"charge.refunded":             ChargeRefunded,
	"charge.chargeback.created":   ChargebackCreated,
	"plan.created":                PlanCreated,
	"subscription.created":        SubsCreated,
	"subscription.paid":           SubsPaid,
	"subscription.canceled":       SubsCanceled,
	"subscription.payment_failed": SubsPaymentFailed,
}

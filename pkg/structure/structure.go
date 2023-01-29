package structure

type SMSData struct {
	Country      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

type VoiceData struct {
	Country      string
	Bandwidth    int
	ResponseTime int
	Provider     string
	Stability    float32
	TTFB         int
	Purity       int
	Duration     int
}

type EmailData struct {
	Country      string
	Provider     string
	DeliveryTime int
}

type BillingData struct {
	CreateCustomer bool
	Purchase       bool
	Payout         bool
	Recurring      bool
	FraudControl   bool
	CheckoutPage   bool
}

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"` // возможные статусы active и closed
}
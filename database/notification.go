package database

type Notification struct {
	Type     string        `json:"type"`
	EventID  string        `json:"eventId"`
	DialogID string        `json:"dialogId,omitempty"`
	Message  DialogMessage `json:"message,omitempty"`
}

const (
	SoldOffer string = "RagfairOfferSold"
	New       string = "new_message"
	Ping      string = "ping"
)

func CreateNotification(message *DialogMessage) *Notification {
	return &Notification{
		Type:     New,
		EventID:  message.ID,
		DialogID: message.UID,
		Message:  *message,
	}
}

func (n *Notification) SendNotification(recipient string) {}

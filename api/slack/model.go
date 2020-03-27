package slack

// ResponseType is a type of slack message type
type ResponseType string

const (
	// InChannel = everyone in the channel can view the message
	InChannel ResponseType = "in_channel"
	// Ephemeral = only the person who triggered the command can view the message
	Ephemeral ResponseType = "ephemeral"
)

// Response is a Slack Response struct
type Response struct {
	Text         string       `json:"text"`
	ResponseType ResponseType `json:"response_type,omitempty"`
}

// EventChallenge is the first event sent when registering the app
//
// The app should return challenge right away
type EventChallenge struct {
	Token     string `json:"token"`
	Challenge string `json:"challenge"`
	Type      string `json:"type"`
}

// EventMessage is a struct sent by Slack
type EventMessage struct {
	Token       string           `json:"token"`
	TeamID      string           `json:"team_id"`
	APIAppID    string           `json:"api_app_id"`
	Event       EventMessageType `json:"event"`
	Type        string           `json:"type"`
	EventID     string           `json:"event_id"`
	EventTime   int              `json:"event_time"`
	AuthedUsers []string         `json:"authed_users"`
}

// EventMessageType holds the information of message type
type EventMessageType struct {
	ClientMsgID string `json:"client_msg_id"`
	Type        string `json:"type"`

	// (e.g. "bot_message")
	SubType EventMessageSubType `json:"subtype,omitempty"`
	Text    string              `json:"text"`
	User    string              `json:"user"`
	Ts      string              `json:"ts"`

	// (optional) if this field exists, then it's a thread reply
	ThreadTs string `json:"thread_ts,omitempty"`

	// (optional) if this field exists, then it's a thread reply
	ParentUserID string `json:"parent_user_id,omitempty"`
	Channel      string `json:"channel"`
	EventTs      string `json:"event_ts"`
	ChannelType  string `json:"channel_type"`
}

// EventMessageSubType represents a sub type of the message.
type EventMessageSubType string

const (
	// EventMessageSubTypeBotMessage is If the message was created by a bot, it
	// contains a subtype bot_message.
	EventMessageSubTypeBotMessage EventMessageSubType = "bot_message"
)

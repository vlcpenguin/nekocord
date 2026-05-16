// internal/models/structs.go
package structs

type User struct {
	UserID        string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Global_Name   string `json:"global_name"`
	Avatar_Hash   string `json:"avatar"`
	Is_Bot        bool   `json:"bot"`
}

type Message struct {
	Message_ID     uint64 `json:"id"`
	Channel_ID     uint64 `json:"channel_id"`
	Message_Author *User  `json:"author"`
	Content        string `json:"content"`
}

type EventData struct {
	Type               int           `json:"type"`
	TTS                bool          `json:"tts"`
	Timestamp          string        `json:"timestamp"`
	Pinned             bool          `json:"pinned"`
	Nonce              string        `json:"nonce"`
	Mentions           []interface{} `json:"mentions"`
	MentionRoles       []interface{} `json:"mention_roles"`
	MentionEveryone    bool          `json:"mention_everyone"`
	ID                 string        `json:"id"`
	Flags              int           `json:"flags"`
	Embeds             []interface{} `json:"embeds"`
	EditedTimestamp    interface{}   `json:"edited_timestamp"`
	Content            string        `json:"content"`
	Components         []interface{} `json:"components"`
	ChannelType        int           `json:"channel_type"`
	ChannelID          string        `json:"channel_id"`
	Author             *User         `json:"author"`
	Attachments        []interface{} `json:"attachments"`
	Heartbeat_Interval int           `json:"heartbeat_interval"`
}

// {"t":null,"s":null,"op":10,"d":{"heartbeat_interval":41250,"_trace":["[\"gateway-prd-arm-us-east1-b-qbf9\",{\"micros\":0.0}]"]}}

type DiscordGatewayMessage struct {
	GatewayOpcode  int         `json:"op"`
	EventData      interface{} `json:"d"`
	SequenceNumber int         `json:"s"`
	EventName      string      `json:"t"`
}

type HeartbeatPayload struct {
	Op int  `json:"op"`
	D  *int `json:"d"`
}

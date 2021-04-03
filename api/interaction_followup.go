package api

// FollowupMessage is used to add additional messages to an Interaction after you've responded initially
type FollowupMessage struct {
	Content         string          `json:"content,omitempty"`
	Username        string          `json:"username,omitempty"`
	AvatarURL       string          `json:"avatar_url,omitempty"`
	TTS             bool             `json:"tts,omitempty"`
	Embeds          []Embed          `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
	//PayloadJSON     string           `json:"payload_json"`
	//File          FileContents     `json:"file"`
}

// FollowupMessageBuilder allows you to create an FollowupMessage with ease
type FollowupMessageBuilder struct {
	FollowupMessage
}

// NewFollowupMessageBuilder returns a new FollowupMessageBuilder
func NewFollowupMessageBuilder() *FollowupMessageBuilder {
	return &FollowupMessageBuilder{
		FollowupMessage{},
	}
}

// SetTTS sets if the FollowupMessage is a tts message
func (b *FollowupMessageBuilder) SetTTS(tts bool) *FollowupMessageBuilder {
	b.TTS = tts
	return b
}

// SetContent sets the content of the FollowupMessage
func (b *FollowupMessageBuilder) SetContent(content string) *FollowupMessageBuilder {
	b.Content = content
	return b
}

// SetEmbeds sets the embeds of the FollowupMessage
func (b *FollowupMessageBuilder) SetEmbeds(embeds ...Embed) *FollowupMessageBuilder {
	b.Embeds = embeds
	return b
}

// AddEmbeds adds multiple embeds to the FollowupMessage
func (b *FollowupMessageBuilder) AddEmbeds(embeds ...Embed) *FollowupMessageBuilder {
	b.Embeds = append(b.Embeds, embeds...)
	return b
}

// ClearEmbeds removes all of the embeds from the FollowupMessage
func (b *FollowupMessageBuilder) ClearEmbeds() *FollowupMessageBuilder {
	if b.Embeds != nil {
		b.Embeds = []Embed{}
	}
	return b
}

// RemoveEmbed removes an embed from the FollowupMessage
func (b *FollowupMessageBuilder) RemoveEmbed(index int) *FollowupMessageBuilder {
	if b != nil && len(b.Embeds) > index {
		b.Embeds = append(b.Embeds[:index], b.Embeds[index+1:]...)
	}
	return b
}

// SetAllowedMentions sets the allowed mentions of the FollowupMessage
func (b *FollowupMessageBuilder) SetAllowedMentions(allowedMentions *AllowedMentions) *FollowupMessageBuilder {
	b.AllowedMentions = allowedMentions
	return b
}

// SetAllowedMentionsEmpty sets the allowed mentions of the FollowupMessage to nothing
func (b *FollowupMessageBuilder) SetAllowedMentionsEmpty() *FollowupMessageBuilder {
	return b.SetAllowedMentions(&AllowedMentions{})
}

// Build returns your built FollowupMessage
func (b *FollowupMessageBuilder) Build() FollowupMessage {
	return b.FollowupMessage
}

package api

import "fmt"

// FollowupMessage is used to add additional messages to an Interaction after you've responded initially
type FollowupMessage struct {
	TTS             bool             `json:"tts,omitempty"`
	Content         string           `json:"content,omitempty"`
	Embeds          []*Embed         `json:"embeds,omitempty"`
	Components      []Component      `json:"components,omitempty"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
	Flags           MessageFlags     `json:"flags,omitempty"`
}

// FollowupMessageBuilder allows you to create an FollowupMessage with ease
type FollowupMessageBuilder struct {
	FollowupMessage
}

// NewFollowupMessageBuilder returns a new FollowupMessageBuilder
func NewFollowupMessageBuilder() *FollowupMessageBuilder {
	return &FollowupMessageBuilder{
		FollowupMessage: FollowupMessage{
			AllowedMentions: &DefaultInteractionAllowedMentions,
		},
	}
}

// NewFollowupMessageBuilderByMessage returns a new FollowupMessageBuilder and takes an existing Message
func NewFollowupMessageBuilderByMessage(message *Message) *FollowupMessageBuilder {
	msg := FollowupMessage{
		TTS:             message.TTS,
		Embeds:          message.Embeds,
		Components:      message.Components,
		AllowedMentions: &DefaultInteractionAllowedMentions,
		Flags:           message.Flags,
	}
	if message.Content != nil {
		msg.Content = *message.Content
	}
	return &FollowupMessageBuilder{
		FollowupMessage: msg,
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

// SetContentf sets the content of the FollowupMessage with format
func (b *FollowupMessageBuilder) SetContentf(content string, a ...interface{}) *FollowupMessageBuilder {
	b.Content = fmt.Sprintf(content, a...)
	return b
}

// SetEmbeds sets the embeds of the FollowupMessage
func (b *FollowupMessageBuilder) SetEmbeds(embeds ...*Embed) *FollowupMessageBuilder {
	b.Embeds = embeds
	return b
}

// AddEmbeds adds multiple embeds to the FollowupMessage
func (b *FollowupMessageBuilder) AddEmbeds(embeds ...*Embed) *FollowupMessageBuilder {
	b.Embeds = append(b.Embeds, embeds...)
	return b
}

// ClearEmbeds removes all of the embeds from the FollowupMessage
func (b *FollowupMessageBuilder) ClearEmbeds() *FollowupMessageBuilder {
	b.Embeds = []*Embed{}
	return b
}

// RemoveEmbed removes an embed from the FollowupMessage
func (b *FollowupMessageBuilder) RemoveEmbed(index int) *FollowupMessageBuilder {
	if b != nil && len(b.Embeds) > index {
		b.Embeds = append(b.Embeds[:index], b.Embeds[index+1:]...)
	}
	return b
}

// SetComponents sets the Component(s) of the FollowupMessage
func (b *FollowupMessageBuilder) SetComponents(components ...Component) *FollowupMessageBuilder {
	b.Components = components
	return b
}

// AddComponents adds the Component(s) to the FollowupMessage
func (b *FollowupMessageBuilder) AddComponents(components ...Component) *FollowupMessageBuilder {
	b.Components = append(b.Components, components...)
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

// SetFlags sets the message flags of the FollowupMessage
func (b *FollowupMessageBuilder) SetFlags(flags MessageFlags) *FollowupMessageBuilder {
	b.Flags = flags
	return b
}

// SetEphemeral adds/removes MessageFlagEphemeral to the message flags
func (b *FollowupMessageBuilder) SetEphemeral(ephemeral bool) *FollowupMessageBuilder {
	if ephemeral {
		b.Flags = b.Flags.Add(MessageFlagEphemeral)
	} else {
		b.Flags = b.Flags.Remove(MessageFlagEphemeral)
	}
	return b
}

// Build returns your built FollowupMessage
func (b *FollowupMessageBuilder) Build() *FollowupMessage {
	return &b.FollowupMessage
}

package api

import "errors"

// ChannelType for interacting with discord's channels
type ChannelType int

// Channel constants
const (
	ChannelTypeText ChannelType = iota
	ChannelTypeDM
	ChannelTypeVoice
	ChannelTypeGroupDM
	ChannelTypeCategory
	ChannelTypeNews
	ChannelTypeStore
)

// Channel is a generic discord channel object
type Channel struct {
	Disgo            Disgo
	ID               Snowflake    `json:"id"`
	Name             *string      `json:"name,omitempty"`
	Type             ChannelType  `json:"type"`
	LastMessageID    *Snowflake   `json:"last_message_id,omitempty"`
	GuildID          *Snowflake   `json:"guild_id,omitempty"`
	Position         *int         `json:"position,omitempty"`
	Topic            *string      `json:"topic,omitempty"`
	NSFW             *bool        `json:"nsfw,omitempty"`
	Bitrate          *int         `json:"bitrate,omitempty"`
	UserLimit        *int         `json:"user_limit,omitempty"`
	RateLimitPerUser *int         `json:"rate_limit_per_user,omitempty"`
	Recipients       []*User      `json:"recipients,omitempty"`
	Icon             *string      `json:"icon,omitempty"`
	OwnerID          *Snowflake   `json:"owner_id,omitempty"`
	ApplicationID    *Snowflake   `json:"application_id,omitempty"`
	ParentID         *Snowflake   `json:"parent_id,omitempty"`
	Permissions      *Permissions `json:"permissions,omitempty"`
	//LastPinTimestamp *time.Time  `json:"last_pin_timestamp,omitempty"`
}

// MessageChannel is used for sending Message(s) to User(s)
type MessageChannel struct {
	Channel
}

// SendMessage sends a Message to a TextChannel
func (c MessageChannel) SendMessage(message MessageCreate) (*Message, error) {
	// Todo: attachments
	return c.Disgo.RestClient().SendMessage(c.ID, message)
}


// EditMessage edits a Message in this TextChannel
func (c MessageChannel) EditMessage(messageID Snowflake, message MessageUpdate) (*Message, error) {
	return c.Disgo.RestClient().EditMessage(c.ID, messageID, message)
}

// DeleteMessage allows you to edit an existing Message sent by you
func (c MessageChannel) DeleteMessage(messageID Snowflake) error {
	return c.Disgo.RestClient().DeleteMessage(c.ID, messageID)
}

// BulkDeleteMessages allows you bulk delete Message(s)
func (c MessageChannel) BulkDeleteMessages(messageIDs ...Snowflake) error {
	return c.Disgo.RestClient().BulkDeleteMessages(c.ID, messageIDs...)
}

// CrosspostMessage crossposts an existing Message
func (c MessageChannel) CrosspostMessage(messageID Snowflake) (*Message, error) {
	if c.Type != ChannelTypeNews {
		return nil, errors.New("channel type is not NEWS")
	}
	return c.Disgo.RestClient().CrosspostMessage(c.ID, messageID)
}

// DMChannel is used for interacting in private Message(s) with users
type DMChannel struct {
	MessageChannel
	Users []User `json:"recipients"`
}

// GuildChannel is a generic type for all server channels
type GuildChannel struct {
	Channel
	GuildID Snowflake `json:"guild_id"`
}

// Guild returns the channel's Guild
func (c GuildChannel) Guild() *Guild {
	return c.Disgo.Cache().Guild(c.GuildID)
}

// Category groups text & voice channels in servers together
type Category struct {
	GuildChannel
}

// VoiceChannel adds methods specifically for interacting with discord's voice
type VoiceChannel struct {
	GuildChannel
}

// TextChannel allows you to interact with discord's text channels
type TextChannel struct {
	GuildChannel
	MessageChannel
}

// StoreChannel allows you to interact with discord's store channels
type StoreChannel struct {
	GuildChannel
}

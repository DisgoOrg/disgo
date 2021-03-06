package api

// CommandInteraction is a specific Interaction when using Command(s)
type CommandInteraction struct {
	*Interaction
	Data *CommandInteractionData `json:"data,omitempty"`
}

// CommandInteractionData is the command data payload
type CommandInteractionData struct {
	ID       Snowflake     `json:"id"`
	Name     string        `json:"name"`
	Resolved *Resolved     `json:"resolved,omitempty"`
	Options  []*OptionData `json:"options,omitempty"`
}

// Resolved contains resolved mention data
type Resolved struct {
	Users    map[Snowflake]*User    `json:"users,omitempty"`
	Members  map[Snowflake]*Member  `json:"members,omitempty"`
	Roles    map[Snowflake]*Role    `json:"roles,omitempty"`
	Channels map[Snowflake]*Channel `json:"channels,omitempty"`
}

// OptionData is used for options or subcommands in your slash commands
type OptionData struct {
	Name    string            `json:"name"`
	Type    CommandOptionType `json:"type"`
	Value   interface{}       `json:"value,omitempty"`
	Options []*OptionData     `json:"options,omitempty"`
}

// Option holds info about an Option.Value
type Option struct {
	Resolved *Resolved
	Name     string
	Type     CommandOptionType
	Value    interface{}
}

// String returns the Option.Value as string
func (o Option) String() string {
	return o.Value.(string)
}

// Integer returns the Option.Value as int
func (o Option) Integer() int {
	return o.Value.(int)
}

// Bool returns the Option.Value as bool
func (o Option) Bool() bool {
	return o.Value.(bool)
}

// Snowflake returns the Option.Value as Snowflake
func (o Option) Snowflake() Snowflake {
	return Snowflake(o.String())
}

// User returns the Option.Value as User
func (o Option) User() *User {
	return o.Resolved.Users[o.Snowflake()]
}

// Member returns the Option.Value as Member
func (o Option) Member() *Member {
	return o.Resolved.Members[o.Snowflake()]
}

// Role returns the Option.Value as Role
func (o Option) Role() *Role {
	return o.Resolved.Roles[o.Snowflake()]
}

// Channel returns the Option.Value as Channel
func (o Option) Channel() *Channel {
	return o.Resolved.Channels[o.Snowflake()]
}

// MessageChannel returns the Option.Value as MessageChannel
func (o Option) MessageChannel() *MessageChannel {
	channel := o.Channel()
	if channel == nil || (channel.Type != ChannelTypeText && channel.Type != ChannelTypeNews) {
		return nil
	}
	return &MessageChannel{Channel: *channel}
}

// GuildChannel returns the Option.Value as GuildChannel
func (o Option) GuildChannel() *GuildChannel {
	channel := o.Channel()
	if channel == nil || (channel.Type != ChannelTypeText && channel.Type != ChannelTypeNews && channel.Type != ChannelTypeCategory && channel.Type != ChannelTypeStore && channel.Type != ChannelTypeVoice) {
		return nil
	}
	return &GuildChannel{Channel: *channel}
}

// VoiceChannel returns the Option.Value as VoiceChannel
func (o Option) VoiceChannel() *VoiceChannel {
	channel := o.Channel()
	if channel == nil || channel.Type != ChannelTypeVoice {
		return nil
	}
	return &VoiceChannel{GuildChannel: GuildChannel{Channel: *channel}}
}

// TextChannel returns the Option.Value as TextChannel
func (o Option) TextChannel() *TextChannel {
	channel := o.Channel()
	if channel == nil || (channel.Type != ChannelTypeText && channel.Type != ChannelTypeNews) {
		return nil
	}
	return &TextChannel{GuildChannel: GuildChannel{Channel: *channel}, MessageChannel: MessageChannel{Channel: *channel}}
}

// Category returns the Option.Value as Category
func (o Option) Category() *Category {
	channel := o.Channel()
	if channel == nil || channel.Type != ChannelTypeCategory {
		return nil
	}
	return &Category{GuildChannel: GuildChannel{Channel: *channel}}
}

// StoreChannel returns the Option.Value as StoreChannel
func (o Option) StoreChannel() *StoreChannel {
	channel := o.Channel()
	if channel == nil || channel.Type != ChannelTypeStore {
		return nil
	}
	return &StoreChannel{GuildChannel: GuildChannel{Channel: *channel}}
}

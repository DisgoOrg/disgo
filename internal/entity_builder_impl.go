package internal

import (
	"github.com/DisgoOrg/disgo/api"
)

func newEntityBuilderImpl(disgo api.Disgo) api.EntityBuilder {
	return &EntityBuilderImpl{disgo: disgo}
}

// EntityBuilderImpl is used for creating structs used by Disgo
type EntityBuilderImpl struct {
	disgo api.Disgo
}

// Disgo returns the api.Disgo client
func (b EntityBuilderImpl) Disgo() api.Disgo {
	return b.disgo
}

// CreateGlobalCommand returns a new api.Command entity
func (b EntityBuilderImpl) CreateGlobalCommand(command *api.Command, updateCache api.CacheStrategy) *api.Command {
	command.Disgo = b.Disgo()
	if updateCache(b.Disgo()) {
		return b.Disgo().Cache().CacheGlobalCommand(command)
	}
	return command
}

// CreateUser returns a new api.User entity
func (b EntityBuilderImpl) CreateUser(user *api.User, updateCache api.CacheStrategy) *api.User {
	user.Disgo = b.Disgo()
	if updateCache(b.Disgo()) {
		return b.Disgo().Cache().CacheUser(user)
	}
	return user
}

// CreateMessage returns a new api.Message entity
func (b EntityBuilderImpl) CreateMessage(message *api.Message, updateCache api.CacheStrategy) *api.Message {
	message.Disgo = b.Disgo()
	if message.Member != nil {
		message.Member = b.CreateMember(*message.GuildID, message.Member, updateCache)
	}
	if message.Author != nil {
		message.Author = b.CreateUser(message.Author, updateCache)
	}
	// TODO: should we cache mentioned users, members, etc?
	if updateCache(b.Disgo()) {
		return b.Disgo().Cache().CacheMessage(message)
	}
	return message
}

// CreateGuild returns a new api.Guild entity
func (b EntityBuilderImpl) CreateGuild(guild *api.Guild, updateCache api.CacheStrategy) *api.Guild {
	guild.Disgo = b.Disgo()
	if updateCache(b.Disgo()) {
		return b.Disgo().Cache().CacheGuild(guild)
	}
	return guild
}

// CreateMember returns a new api.Member entity
func (b EntityBuilderImpl) CreateMember(guildID api.Snowflake, member *api.Member, updateCache api.CacheStrategy) *api.Member {
	member.Disgo = b.Disgo()
	member.GuildID = guildID
	member.User = b.CreateUser(member.User, updateCache)
	if updateCache(b.Disgo()) {
		return b.Disgo().Cache().CacheMember(member)
	}
	return member
}

// CreateVoiceState returns a new api.VoiceState entity
func (b EntityBuilderImpl) CreateVoiceState(voiceState *api.VoiceState, updateCache api.CacheStrategy) *api.VoiceState {
	voiceState.Disgo = b.Disgo()
	if updateCache(b.Disgo()) {
		return b.Disgo().Cache().CacheVoiceState(voiceState)
	}
	return voiceState
}

// CreateGuildCommand returns a new api.Command entity
func (b EntityBuilderImpl) CreateGuildCommand(guildID api.Snowflake, command *api.Command, updateCache api.CacheStrategy) *api.Command {
	command.Disgo = b.Disgo()
	command.GuildID = &guildID
	if updateCache(b.Disgo()) {
		return b.Disgo().Cache().CacheGuildCommand(command)
	}
	return command
}

// CreateGuildCommandPermissions returns a new api.GuildCommandPermissions entity
func (b EntityBuilderImpl) CreateGuildCommandPermissions(guildCommandPermissions *api.GuildCommandPermissions, updateCache api.CacheStrategy) *api.GuildCommandPermissions {
	guildCommandPermissions.Disgo = b.Disgo()
	if updateCache(b.Disgo()) && b.Disgo().Cache().CacheFlags().Has(api.CacheFlagCommandPermissions) {
		if cmd := b.Disgo().Cache().Command(guildCommandPermissions.ID); cmd != nil {
			cmd.GuildPermissions[guildCommandPermissions.GuildID] = guildCommandPermissions
		}
	}
	return guildCommandPermissions
}

// CreateRole returns a new api.Role entity
func (b EntityBuilderImpl) CreateRole(guildID api.Snowflake, role *api.Role, updateCache api.CacheStrategy) *api.Role {
	role.Disgo = b.Disgo()
	role.GuildID = guildID
	if updateCache(b.Disgo()) {
		return b.Disgo().Cache().CacheRole(role)
	}
	return role
}

// CreateTextChannel returns a new api.TextChannel entity
func (b EntityBuilderImpl) CreateTextChannel(channel *api.Channel, updateCache api.CacheStrategy) *api.TextChannel {
	channel.Disgo = b.Disgo()
	textChannel := &api.TextChannel{
		MessageChannel: api.MessageChannel{
			Channel: *channel,
		},
		GuildChannel: api.GuildChannel{
			Channel: *channel,
		},
	}
	if updateCache(b.Disgo()) {
		return b.Disgo().Cache().CacheTextChannel(textChannel)
	}
	return textChannel
}

// CreateVoiceChannel returns a new api.VoiceChannel entity
func (b EntityBuilderImpl) CreateVoiceChannel(channel *api.Channel, updateCache api.CacheStrategy) *api.VoiceChannel {
	channel.Disgo = b.Disgo()
	voiceChannel := &api.VoiceChannel{
		GuildChannel: api.GuildChannel{
			Channel: *channel,
		},
	}
	if updateCache(b.Disgo()) {
		return b.Disgo().Cache().CacheVoiceChannel(voiceChannel)
	}
	return voiceChannel
}

// CreateStoreChannel returns a new api.StoreChannel entity
func (b EntityBuilderImpl) CreateStoreChannel(channel *api.Channel, updateCache api.CacheStrategy) *api.StoreChannel {
	channel.Disgo = b.Disgo()
	storeChannel := &api.StoreChannel{
		GuildChannel: api.GuildChannel{
			Channel: *channel,
		},
	}
	if updateCache(b.Disgo()) {
		return b.Disgo().Cache().CacheStoreChannel(storeChannel)
	}
	return storeChannel
}

// CreateCategory returns a new api.Category entity
func (b EntityBuilderImpl) CreateCategory(channel *api.Channel, updateCache api.CacheStrategy) *api.Category {
	channel.Disgo = b.Disgo()
	category := &api.Category{
		GuildChannel: api.GuildChannel{
			Channel: *channel,
		},
	}
	if updateCache(b.Disgo()) {
		return b.Disgo().Cache().CacheCategory(category)
	}
	return category
}

// CreateDMChannel returns a new api.DMChannel entity
func (b EntityBuilderImpl) CreateDMChannel(channel *api.Channel, updateCache api.CacheStrategy) *api.DMChannel {
	channel.Disgo = b.Disgo()
	dmChannel := &api.DMChannel{
		MessageChannel: api.MessageChannel{
			Channel: *channel,
		},
	}
	if updateCache(b.Disgo()) {
		return b.Disgo().Cache().CacheDMChannel(dmChannel)
	}
	return dmChannel
}

// CreateEmote returns a new api.Emote entity
func (b EntityBuilderImpl) CreateEmote(guildId api.Snowflake, emote *api.Emote, updateCache api.CacheStrategy) *api.Emote {
	emote.Disgo = b.Disgo()
	emote.GuildID = guildId
	if updateCache(b.Disgo()) {
		return b.Disgo().Cache().CacheEmote(emote)
	}
	return emote
}
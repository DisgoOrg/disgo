package disgo

type Intent int

const (
	IntentsGuilds Intent = 1 << iota
	IntentsGuildMembers
	IntentsGuildBans
	IntentsGuildEmojis
	IntentsGuildIntegrations
	IntentsGuildWebhooks
	IntentsGuildInvites
	IntentsGuildVoiceStates
	IntentsGuildPresences
	IntentsGuildMessages
	IntentsGuildMessageReactions
	IntentsGuildMessageTyping
	IntentsDirectMessages
	IntentsDirectMessageReactions
	IntentsDirectMessageTyping

	IntentsAllWithoutPrivileged = IntentsGuilds |
		IntentsGuildBans |
		IntentsGuildEmojis |
		IntentsGuildIntegrations |
		IntentsGuildWebhooks |
		IntentsGuildInvites |
		IntentsGuildVoiceStates |
		IntentsGuildMessages |
		IntentsGuildMessageReactions |
		IntentsGuildMessageTyping |
		IntentsDirectMessages |
		IntentsDirectMessageReactions |
		IntentsDirectMessageTyping
	IntentsAll = IntentsAllWithoutPrivileged |
		IntentsGuildMembers |
		IntentsGuildPresences
	IntentsNone Intent = 0
)

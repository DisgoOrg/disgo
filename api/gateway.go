package api

import (
	"time"

	"github.com/gorilla/websocket"
)

// GatewayStatus is the state that the client is currently in
type GatewayStatus int

// Indicates how far along the client is to connecting
const (
	Ready GatewayStatus = iota
	Unconnected
	Connecting
	Reconnecting
	WaitingForHello
	WaitingForReady
	Disconnected
	WaitingForGuilds
	Identifying
	Resuming
)

// Gateway is what is used to connect to discord
type Gateway interface {
	Disgo() Disgo
	Open() error
	Status() GatewayStatus
	Close()
	Conn() *websocket.Conn
	Latency() time.Duration
}

// GatewayOp are opcodes used by discord
type GatewayOp int

// Constants for the gateway opcodes
const (
	OpDispatch GatewayOp = iota
	OpHeartbeat
	OpIdentify
	OpPresenceUpdate
	OpVoiceStateUpdate
	_
	OpResume
	OpReconnect
	OpRequestGuildMembers
	OpInvalidSession
	OpHello
	OpHeartbeatACK
)

// GatewayEventType wraps all GatewayEventType types
type GatewayEventType string

// Constants for the gateway events
const (
	GatewayEventHello                      GatewayEventType = "HELLO"
	GatewayEventReady                      GatewayEventType = "READY"
	GatewayEventResumed                    GatewayEventType = "RESUMED"
	GatewayEventReconnect                  GatewayEventType = "RECONNECT"
	GatewayEventInvalidSession             GatewayEventType = "INVALID_SESSION"
	GatewayEventApplicationCommandCreate   GatewayEventType = "APPLICATION_COMMAND_CREATE"
	GatewayEventApplicationCommandUpdate   GatewayEventType = "APPLICATION_COMMAND_UPDATE"
	GatewayEventApplicationCommandDelete   GatewayEventType = "APPLICATION_COMMAND_DELETE"
	GatewayEventChannelCreate              GatewayEventType = "CHANNEL_CREATE"
	GatewayEventChannelUpdate              GatewayEventType = "CHANNEL_UPDATE"
	GatewayEventChannelDelete              GatewayEventType = "CHANNEL_DELETE"
	GatewayEventChannelPinsUpdate          GatewayEventType = "CHANNEL_PINS_UPDATE"
	GatewayEventThreadCreate               GatewayEventType = "THREAD_CREATE"
	GatewayEventThreadUpdate               GatewayEventType = "THREAD_UPDATE"
	GatewayEventThreadDelete               GatewayEventType = "THREAD_DELETE"
	GatewayEventThreadListSync             GatewayEventType = "THREAD_LIST_sync"
	GatewayEventThreadMemberUpdate         GatewayEventType = "THREAD_MEMBER_UPDATE"
	GatewayEventThreadMembersUpdate        GatewayEventType = "THREAD_MEMBERS_UPDATE"
	GatewayEventGuildCreate                GatewayEventType = "GUILD_CREATE"
	GatewayEventGuildUpdate                GatewayEventType = "GUILD_UPDATE"
	GatewayEventGuildDelete                GatewayEventType = "GUILD_DELETE"
	GatewayEventGuildBanAdd                GatewayEventType = "GUILD_BAN_ADD"
	GatewayEventGuildBanRemove             GatewayEventType = "GUILD_BAN_REMOVE"
	GatewayEventGuildEmojisUpdate          GatewayEventType = "GUILD_EMOJIS_UPDATE"
	GatewayEventGuildIntegrationsUpdate    GatewayEventType = "GUILD_INTEGRATIONS_UPDATE"
	GatewayEventGuildMemberAdd             GatewayEventType = "GUILD_MEMBER_ADD"
	GatewayEventGuildMemberRemove          GatewayEventType = "GUILD_MEMBER_REMOVE"
	GatewayEventGuildMemberUpdate          GatewayEventType = "GUILD_MEMBER_UPDATE"
	GatewayEventGuildMembersChunk          GatewayEventType = "GUILD_MEMBERS_CHUNK"
	GatewayEventGuildRoleCreate            GatewayEventType = "GUILD_ROLE_CREATE"
	GatewayEventGuildRoleUpdate            GatewayEventType = "GUILD_ROLE_UPDATE"
	GatewayEventGuildRoleDelete            GatewayEventType = "GUILD_ROLE_DELETE"
	GatewayEventIntegrationCreate          GatewayEventType = "INTEGRATION_CREATE"
	GatewayEventIntegrationUpdate          GatewayEventType = "INTEGRATION_UPDATE"
	GatewayEventIntegrationDelete          GatewayEventType = "INTEGRATION_DELETE"
	GatewayEventInteractionCreate          GatewayEventType = "INTERACTION_CREATE"
	WebhookEventInteractionCreate          GatewayEventType = "WEBHOOK_INTERACTION_CREATE"
	GatewayEventInviteCreate               GatewayEventType = "INVITE_CREATE"
	GatewayEventInviteDelete               GatewayEventType = "INVITE_DELETE"
	GatewayEventMessageCreate              GatewayEventType = "MESSAGE_CREATE"
	GatewayEventMessageUpdate              GatewayEventType = "MESSAGE_UPDATE"
	GatewayEventMessageDelete              GatewayEventType = "MESSAGE_DELETE"
	GatewayEventMessageDeleteBulk          GatewayEventType = "MESSAGE_DELETE_BULK"
	GatewayEventMessageReactionAdd         GatewayEventType = "MESSAGE_REACTION_ADD"
	GatewayEventMessageReactionRemove      GatewayEventType = "MESSAGE_REACTION_REMOVE"
	GatewayEventMessageReactionRemoveAll   GatewayEventType = "MESSAGE_REACTION_REMOVE_ALL"
	GatewayEventMessageReactionRemoveEmoji GatewayEventType = "MESSAGE_REACTION_REMOVE_EMOJI"
	GatewayEventPresenceUpdate             GatewayEventType = "PRESENCE_UPDATE"
	GatewayEventTypingStart                GatewayEventType = "TYPING_START"
	GatewayEventUserUpdate                 GatewayEventType = "USER_UPDATE"
	GatewayEventVoiceStateUpdate           GatewayEventType = "VOICE_STATE_UPDATE"
	GatewayEventVoiceServerUpdate          GatewayEventType = "VOICE_SERVER_UPDATE"
	GatewayEventWebhooksUpdate             GatewayEventType = "WEBHOOKS_UPDATE"
)

// GatewayRs contains the response for GET /gateway
type GatewayRs struct {
	URL string `json:"url"`
}

// GatewayBotRs contains the response for GET /gateway/bot
type GatewayBotRs struct {
	URL               string `json:"url"`
	Shards            int    `json:"shards"`
	SessionStartLimit struct {
		Total          int `json:"total"`
		Remaining      int `json:"remaining"`
		ResetAfter     int `json:"reset_after"`
		MaxConcurrency int `json:"max_concurrency"`
	} `json:"session_start_limit"`
}

package api

import (
	"encoding/json"
	"time"
)

// GatewayPacket raw GatewayEvent type
type GatewayPacket struct {
	Op GatewayOp         `json:"op"`
	S  *int              `json:"s,omitempty"`
	T  *GatewayEventType `json:"t,omitempty"`
}

// RawGatewayEvent specifies the data for the GatewayCommand payload that is being sent
type RawGatewayEvent struct {
	GatewayPacket
	D json.RawMessage `json:"d"`
}

// ReadyGatewayEvent is the event sent by discord when you successfully Identify
type ReadyGatewayEvent struct {
	Version   int      `json:"v"`
	SelfUser  User     `json:"user"`
	Guilds    []*Guild `json:"guilds"`
	SessionID string   `json:"session_id"`
	Shard     *[2]int  `json:"shard,omitempty"`
}

// HelloGatewayEventData is sent when we connect to the gateway
type HelloGatewayEventData struct {
	HeartbeatInterval time.Duration `json:"heartbeat_interval"`
}

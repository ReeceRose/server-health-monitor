package types

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// StandardResponse is used to standardize the returns for all HTTP endpoint
type StandardResponse struct {
	Data       interface{}
	StatusCode int
	Error      string
	Success    bool
}

// Health contains all information realted to an agents health
type Health struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	AgentID    string             `json:"agentID" bson:"agentID"`
	CreateTime int64              `json:"createTime" bson:"createTime"`
	Uptime     uint64             `json:"uptime" bson:"uptime"`
}

// Host contains all information about the agent/host
type Host struct {
	ID                   primitive.ObjectID `json:"_id" bson:"_id"`
	AgentID              string             `json:"agentID" bson:"agentID"` // NOTE: Host.HostID is an alternative solution to AgentID
	CreateTime           int64              `json:"createTime" bson:"createTime"`
	UpdateTime           int64              `json:"updateTime" bson:"updateTime"`
	Hostname             string             `json:"hostname" bson:"hostname"`
	Uptime               uint64             `json:"uptime" bson:"uptime"`
	BootTime             uint64             `json:"bootTime" bson:"bootTime"`
	Procs                uint64             `json:"procs" bson:"procs"`
	OS                   string             `json:"os" bson:"os"`
	Platform             string             `json:"platform" bson:"platform"`
	PlatformFamily       string             `json:"platformFamily" bson:"platformFamily"`
	PlatformVersion      string             `json:"platformVersion" bson:"platformVersion"`
	KernelVersion        string             `json:"kernelVersion" bson:"kernelVersion"`
	KernelArch           string             `json:"kernelArch" bson:"kernelArch"`
	VirtualizationSystem string             `json:"virtualizationSystem" bson:"virtualizationSystem"`
	VirtualizationRole   string             `json:"virtualizationRole" bson:"virtualizationRole"`
	HostID               string             `json:"hostId" bson:"hostId"`
	Online               bool               `json:"online" bson:",omitempty"`
	LastConnected        int64              `json:"lastConnected" bson:",omitempty"`
	Health               []Health           `json:"health" bson:",omitempty"`
}

// AgentInformation contains an ID which is used to differentiate between different agents
type AgentInformation struct {
	ID uuid.UUID
}

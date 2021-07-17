package types

import (
	"github.com/google/uuid"
)

type StandardResponse struct {
	Data       interface{}
	StatusCode int
	Error      string
	Success    bool
}

type Health struct {
	// TODO: decide if host should be apart of 'health' or just apart of 'host'
	// (host information should mostly stay the same, and maybe we'll just send it at startup,
	// we can either re-send non-static items like uptime or maybe calculate that based off of boottime?)
	Host    Host   `json:"host" bson:"host"`
	AgentID string `json:"agentID" bson:"agentID"`
}

type Host struct {
	Hostname             string `json:"hostname" bson:"hostname"`
	Uptime               uint64 `json:"uptime" bson:"uptime"`
	BootTime             uint64 `json:"bootTime" bson:"bootTime"`
	Procs                uint64 `json:"procs" bson:"procs"`
	OS                   string `json:"os" bson:"os"`
	Platform             string `json:"platform" bson:"platform"`
	PlatformFamily       string `json:"platformFamily" bson:"platformFamily"`
	PlatformVersion      string `json:"platformVersion" bson:"platformVersion"`
	KernelVersion        string `json:"kernelVersion" bson:"kernelVersion"`
	KernelArch           string `json:"kernelArch" bson:"kernelArch"`
	VirtualizationSystem string `json:"virtualizationSystem" bson:"virtualizationSystem"`
	VirtualizationRole   string `json:"virtualizationRole" bson:"virtualizationRole"`
	HostID               string `json:"hostId" bson:"hostId"`
}

type AgentInformation struct {
	ID uuid.UUID
}

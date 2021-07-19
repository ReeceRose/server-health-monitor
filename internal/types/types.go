package types

import (
	"io/fs"
	"os"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	Host       Host               `json:"host" bson:"host"`
	AgentID    string             `json:"agentID" bson:"agentID"` // NOTE: Host.HostID is an alternative solution to AgentID
	CreateTime int64              `json:"createTime" bson:"createTime"`
	UpdateTime int64              `json:"updateTime" bson:"updateTime"`
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

type OperatingSystem interface {
	OpenFile(string, int, fs.FileMode) (*os.File, error)
	ReadFile(string) ([]byte, error)
	WriteFile(string, []byte, os.FileMode) error
	IsNotExist(error) bool
	Stat(string) (os.FileInfo, error)
	Remove(string) error
}

type HostInformation interface {
	Info() (*Host, error)
}

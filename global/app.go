package global

import "time"

const (
	// Debug 开发
	Debug = "Debug"
	// ReleaseDebug 外测
	ReleaseDebug = "ReleaseDebug"
	// Release 正式
	Release = "Release"
)

type app struct {
	Name       string
	Version    string
	Date       time.Time
	Copyright  string
	LaunchTime time.Time
	Uptime     time.Duration
	Env        string
	Host       string
	Port       string
	BaseURL    string
	Domain     string
}

// App 初始化
var App = &app{}

func init() {
	App.Name = "weibo"
	App.Version = "0.1"
	App.LaunchTime = time.Now()
	App.Env = Debug
}

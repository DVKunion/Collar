package config

import (
	"io"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

const CfgPath = ".config.yml"

var Version = ""

var SingleConfig = &Config{}

type Config struct {
	Token    string    `json:"token" yaml:"token"`
	UpdateAt time.Time `json:"update_at" yaml:"update_at"`
	HostList []Host    `json:"host_list" yaml:"host_list"`
}

type Host struct {
	Id         string      `json:"id" yaml:"id"`
	Name       string      `json:"name" yaml:"name"`
	OsType     string      `json:"os_type" yaml:"os_type"`
	OsName     string      `json:"os_name" yaml:"os_name"`
	OrgId      string      `json:"org_id" yaml:"org_id"`
	UserId     string      `json:"user_id" yaml:"user_id"`
	Cpu        int         `json:"cpu" yaml:"cpu"`
	Memory     int64       `json:"memory" yaml:"memory"`
	InternalIp string      `json:"internal_ip" yaml:"internal_ip"`
	ExternalIp string      `json:"external_ip" yaml:"external_ip"`
	Comment    interface{} `json:"comment" yaml:"comment"`
	Arch       string      `json:"arch" yaml:"arch"`
	Status     string      `json:"status" yaml:"status"`
	Uptime     int         `json:"uptime" yaml:"uptime"`
	Ipinfo     struct {
		Country  string `json:"country" yaml:"country"`
		Province string `json:"province" yaml:"province"`
		City     string `json:"city" yaml:"city"`
	} `json:"ipinfo" yaml:"ipinfo"`
	LatestTime    int    `json:"latest_time" yaml:"latest_time"`
	EngineVersion string `json:"engine_version" yaml:"engine_version"`
}

func (c *Config) Load() error {
	f, err := os.Open(CfgPath)
	if err != nil {
		return err
	}
	t, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(t, &c)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) GetOnline() []Host {
	var online []Host
	for _, h := range c.HostList {
		if h.Status == "online" {
			online = append(online, h)
		}
	}
	return online
}

func (c *Config) Save() error {
	data, err := yaml.Marshal(&c)

	f, err := os.Create(CfgPath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}

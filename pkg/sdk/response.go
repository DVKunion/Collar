package sdk

import "github.com/DVKunion/collar/pkg/config"

type Response interface {
	Success() bool
	Body() interface{}
}

type HostResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Edges       []config.Host `json:"edges"`
		HasNextPage bool          `json:"has_next_page"`
		TotalCount  int           `json:"total_count"`
	}
}

func (h *HostResponse) Success() bool {
	if h.Code != 200 || h.Message != "ok" {
		return false
	}
	return true
}

func (h *HostResponse) Body() interface{} {
	return h.Data.Edges
}

type ProcessResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		DataTime  int       `json:"data_time"`
		Processes []Process `json:"processes"`
	}
}

type Process struct {
	Name        string  `json:"name"`
	Pid         int     `json:"pid"`
	User        string  `json:"user"`
	Uid         int     `json:"uid"`
	StartTime   int     `json:"start_time"`
	CpuUsage    float64 `json:"cpu_usage"`
	Memory      int     `json:"memory"`
	MemoryUsage float64 `json:"memory_usage"`
	Cmd         string  `json:"cmd"`
}

func (p *ProcessResponse) Success() bool {
	if p.Code != 200 || p.Message != "ok" {
		return false
	}
	return true
}

func (p *ProcessResponse) Body() interface{} {
	return p.Data.Processes
}

type AutoLoginResponse struct {
	Message string `json:"message"`
	Data    struct {
		Id     int    `json:"id"`
		HostId string `json:"host_id"`
		Name   string `json:"name"`
	} `json:"data"`
	Code int `json:"code"`
}

func (a *AutoLoginResponse) Success() bool {
	if a.Code != 200 || a.Message != "ok" {
		return false
	}
	return true
}

func (a *AutoLoginResponse) Body() interface{} {
	return a.Data.Name
}

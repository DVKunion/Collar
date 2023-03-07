package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/DVKunion/collar/pkg/terminal"
	"github.com/DVKunion/collar/pkg/utils"
)

const (
	Host          = "collie.rivers.chaitin.cn"
	HostList      = "api/v1/host/list?page=1&size=1000"
	ProcessList   = "api/v1/host/%s/processes?id=%s&field=cpu_usage&order_by=desc&type=instant"
	LoginList     = "api/v1/host/%s/login_history?id=%s&type=instant"
	ImageList     = "api/v1/docker/%s/images?id=%s"
	ContainerList = "api/v1/docker/%s/containers?id=%s"

	AutoLoginUser   = "api/v1/host/%s/auto_login_user?id=%s"
	TerminalWss     = "api/v1/ws/terminal?id=%s"
	TerminalWssAuto = "api/v1/ws/terminal?id=%s&user=%s"
)

func GetHostList() (Response, error) {
	originResp, err := utils.Get(strings.Join([]string{"https:/", Host, HostList}, "/"))
	if err != nil {
		return nil, err
	}
	resp := &HostResponse{}
	err = json.Unmarshal(originResp.Body(), &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetProcessList(hostId string) (Response, error) {
	originResp, err := utils.Get(strings.Join([]string{"https:/", Host, fmt.Sprintf(ProcessList, hostId, hostId)}, "/"))
	if err != nil {
		return nil, err
	}
	resp := &ProcessResponse{}
	err = json.Unmarshal(originResp.Body(), &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetLoginList(hostId string) (Response, error) {
	originResp, err := utils.Get(strings.Join([]string{"https:/", Host, fmt.Sprintf(LoginList, hostId, hostId)}, "/"))
	if err != nil {
		return nil, err
	}
	resp := &HostResponse{}
	err = json.Unmarshal(originResp.Body(), &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetImageList(hostId string) (Response, error) {
	originResp, err := utils.Get(strings.Join([]string{"https:/", Host, fmt.Sprintf(ImageList, hostId, hostId)}, "/"))
	if err != nil {
		return nil, err
	}
	resp := &HostResponse{}
	err = json.Unmarshal(originResp.Body(), &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetContainerList(hostId string) (Response, error) {
	originResp, err := utils.Get(strings.Join([]string{"https:/", Host, fmt.Sprintf(ContainerList, hostId, hostId)}, "/"))
	if err != nil {
		return nil, err
	}
	resp := &HostResponse{}
	err = json.Unmarshal(originResp.Body(), &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetAutoLoginUser(hostId string) (Response, error) {
	originResp, err := utils.Get(strings.Join([]string{"https:/", Host, fmt.Sprintf(AutoLoginUser, hostId, hostId)}, "/"))
	if err != nil {
		return nil, err
	}
	resp := &AutoLoginResponse{}
	err = json.Unmarshal(originResp.Body(), &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetTerminal(ctx context.Context, hostId string, user string) (*terminal.Terminal, error) {
	tm := terminal.NewTerminal(ctx)
	if user == "" {
		return tm, tm.Connect(strings.Join([]string{"wss:/", Host, fmt.Sprintf(TerminalWss, hostId)}, "/"))
	}
	return tm, tm.Connect(strings.Join([]string{"wss:/", Host, fmt.Sprintf(TerminalWssAuto, hostId, user)}, "/"))
}

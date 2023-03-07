package utils

import "github.com/DVKunion/collar/pkg/config"

// Trans2HostId transfer users input 2 hostId
func Trans2HostId(arg string) string {
	// Just Get Online host
	hostList := config.SingleConfig.GetOnline()
	for _, h := range hostList {
		// check id/name/internal_ip/external_ip
		if h.Id == arg || h.Name == arg || h.ExternalIp == arg || h.InternalIp == arg {
			return h.Id
		}
	}
	return ""
}

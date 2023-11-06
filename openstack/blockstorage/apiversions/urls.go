package apiversions

import (
	"strings"

	gophercloud "github.com/nhn/nhncloud.gophercloud"
	"github.com/nhn/nhncloud.gophercloud/openstack/utils"
)

func listURL(c *gophercloud.ServiceClient) string {
	baseEndpoint, _ := utils.BaseEndpoint(c.Endpoint)
	endpoint := strings.TrimRight(baseEndpoint, "/") + "/"
	return endpoint
}

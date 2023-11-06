package limits

import (
	gophercloud "github.com/nhn/nhncloud.gophercloud"
)

const resourcePath = "limits"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

package buildinfo

import gophercloud "github.com/nhn/nhncloud.gophercloud"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("build_info")
}

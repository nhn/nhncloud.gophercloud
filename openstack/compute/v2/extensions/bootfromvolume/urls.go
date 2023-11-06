package bootfromvolume

import gophercloud "github.com/nhn/nhncloud.gophercloud"

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("servers")
}

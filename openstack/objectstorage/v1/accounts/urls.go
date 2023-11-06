package accounts

import gophercloud "github.com/nhn/nhncloud.gophercloud"

func getURL(c *gophercloud.ServiceClient) string {
	return c.Endpoint
}

func updateURL(c *gophercloud.ServiceClient) string {
	return getURL(c)
}

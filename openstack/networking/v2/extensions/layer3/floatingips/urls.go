package floatingips

import gophercloud "github.com/nhn/nhncloud.gophercloud"

const resourcePath = "floatingips"

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id)
}

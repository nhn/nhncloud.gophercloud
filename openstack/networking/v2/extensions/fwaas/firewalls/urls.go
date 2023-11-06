package firewalls

import gophercloud "github.com/nhn/nhncloud.gophercloud"

const (
	rootPath     = "fw"
	resourcePath = "firewalls"
)

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}

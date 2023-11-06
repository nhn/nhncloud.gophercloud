package quotas

import gophercloud "github.com/nhn/nhncloud.gophercloud"

const resourcePath = "quotas"

func resourceURL(c *gophercloud.ServiceClient, projectID string) string {
	return c.ServiceURL(resourcePath, projectID)
}

func getURL(c *gophercloud.ServiceClient, projectID string) string {
	return resourceURL(c, projectID)
}

func updateURL(c *gophercloud.ServiceClient, projectID string) string {
	return resourceURL(c, projectID)
}

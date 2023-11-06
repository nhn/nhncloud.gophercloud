package volumeactions

import gophercloud "github.com/nhn/nhncloud.gophercloud"

func actionURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("volumes", id, "action")
}

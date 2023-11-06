package schedulerstats

import gophercloud "github.com/nhn/nhncloud.gophercloud"

func storagePoolsListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("scheduler-stats", "get_pools")
}

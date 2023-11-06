package serviceassets

import gophercloud "github.com/nhn/nhncloud.gophercloud"

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("services", id, "assets")
}

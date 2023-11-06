package services

import gophercloud "github.com/nhn/nhncloud.gophercloud"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("services")
}

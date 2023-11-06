package services

import gophercloud "github.com/nhn/nhncloud.gophercloud"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("services")
}

func createURL(c *gophercloud.ServiceClient) string {
	return listURL(c)
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("services", id)
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	return getURL(c, id)
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return getURL(c, id)
}

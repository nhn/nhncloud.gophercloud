package databases

import gophercloud "github.com/nhn/nhncloud.gophercloud"

func baseURL(c *gophercloud.ServiceClient, instanceID string) string {
	return c.ServiceURL("instances", instanceID, "databases")
}

func dbURL(c *gophercloud.ServiceClient, instanceID, dbName string) string {
	return c.ServiceURL("instances", instanceID, "databases", dbName)
}

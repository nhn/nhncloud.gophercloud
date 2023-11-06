package extraroutes

import gophercloud "github.com/nhn/nhncloud.gophercloud"

const resourcePath = "routers"

func addExtraRoutesURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id, "add_extraroutes")
}

func removeExtraRoutesURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id, "remove_extraroutes")
}

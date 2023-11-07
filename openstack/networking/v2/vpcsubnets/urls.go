package vpcsubnets

import "github.com/gophercloud/gophercloud"

func resourceURL(c *gophercloud.ServiceClient, id string, action ...string) string {
	// @tc-iaas-compute/1452
	// action is a variadic parameter, but only the first element is considered,
	// and this is only to ensure backward compatibility.
	if len(action) == 1 {
		return c.ServiceURL("vpcsubnets", id, action[0])
	}
	return c.ServiceURL("vpcsubnets", id)
}

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("vpcsubnets")
}

func listURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func createURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func attachURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id, "attach_routingtable")
}

func detachURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id, "detach_routingtable")
}

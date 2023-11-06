package extensions

import gophercloud "github.com/nhn/nhncloud.gophercloud"

func ActionURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "action")
}

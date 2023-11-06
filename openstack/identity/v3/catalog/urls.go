package catalog

import gophercloud "github.com/nhn/nhncloud.gophercloud"

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("auth", "catalog")
}

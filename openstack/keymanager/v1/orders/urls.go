package orders

import gophercloud "github.com/nhn/nhncloud.gophercloud"

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("orders")
}

func getURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("orders", id)
}

func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("orders")
}

func deleteURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("orders", id)
}

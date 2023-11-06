package drivers

import gophercloud "github.com/nhn/nhncloud.gophercloud"

func driversURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("drivers")
}

func driverDetailsURL(client *gophercloud.ServiceClient, driverName string) string {
	return client.ServiceURL("drivers", driverName)
}

func driverPropertiesURL(client *gophercloud.ServiceClient, driverName string) string {
	return client.ServiceURL("drivers", driverName, "properties")
}

func driverDiskPropertiesURL(client *gophercloud.ServiceClient, driverName string) string {
	return client.ServiceURL("drivers", driverName, "raid", "logical_disk_properties")
}

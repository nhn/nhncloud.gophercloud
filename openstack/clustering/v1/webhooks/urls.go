package webhooks

import gophercloud "github.com/nhn/nhncloud.gophercloud"

func triggerURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("v1", "webhooks", id, "trigger")
}

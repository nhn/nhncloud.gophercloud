package resetnetwork

import (
	gophercloud "github.com/nhn/nhncloud.gophercloud"
	"github.com/nhn/nhncloud.gophercloud/openstack/compute/v2/extensions"
)

// ResetNetwork will reset the network of a server
func ResetNetwork(client *gophercloud.ServiceClient, id string) (r ResetResult) {
	b := map[string]interface{}{
		"resetNetwork": nil,
	}
	resp, err := client.Post(extensions.ActionURL(client, id), b, nil, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

package injectnetworkinfo

import (
	gophercloud "github.com/nhn/nhncloud.gophercloud"
	"github.com/nhn/nhncloud.gophercloud/openstack/compute/v2/extensions"
)

// InjectNetworkInfo will inject the network info into a server
func InjectNetworkInfo(client *gophercloud.ServiceClient, id string) (r InjectNetworkResult) {
	b := map[string]interface{}{
		"injectNetworkInfo": nil,
	}
	resp, err := client.Post(extensions.ActionURL(client, id), b, nil, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

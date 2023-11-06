package common

import (
	gophercloud "github.com/nhn/nhncloud.gophercloud"
	"github.com/nhn/nhncloud.gophercloud/testhelper/client"
)

const TokenID = client.TokenID

func ServiceClient() *gophercloud.ServiceClient {
	sc := client.ServiceClient()
	sc.ResourceBase = sc.Endpoint + "v2.0/"
	return sc
}

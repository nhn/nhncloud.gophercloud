package swauth

import gophercloud "github.com/nhn/nhncloud.gophercloud"

func getURL(c *gophercloud.ProviderClient) string {
	return c.IdentityBase + "auth/v1.0"
}

package tokens

import gophercloud "github.com/nhn/nhncloud.gophercloud"

func tokenURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("auth", "tokens")
}

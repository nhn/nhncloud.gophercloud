package buildinfo

import gophercloud "github.com/nhn/nhncloud.gophercloud"

// Get retreives data for the given stack template.
func Get(c *gophercloud.ServiceClient) (r GetResult) {
	resp, err := c.Get(getURL(c), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

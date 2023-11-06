package apiversions

import (
	gophercloud "github.com/nhn/nhncloud.gophercloud"
	"github.com/nhn/nhncloud.gophercloud/pagination"
)

// List lists all the Cinder API versions available to end-users.
func List(c *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(c, listURL(c), func(r pagination.PageResult) pagination.Page {
		return APIVersionPage{pagination.SinglePageBase(r)}
	})
}

package injectnetworkinfo

import (
	gophercloud "github.com/nhn/nhncloud.gophercloud"
)

// InjectNetworkResult is the response of a InjectNetworkInfo operation. Call
// its ExtractErr method to determine if the request suceeded or failed.
type InjectNetworkResult struct {
	gophercloud.ErrResult
}

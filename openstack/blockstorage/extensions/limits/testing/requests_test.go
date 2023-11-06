package testing

import (
	"testing"

	"github.com/nhn/nhncloud.gophercloud/openstack/blockstorage/extensions/limits"
	th "github.com/nhn/nhncloud.gophercloud/testhelper"
	"github.com/nhn/nhncloud.gophercloud/testhelper/client"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	actual, err := limits.Get(client.ServiceClient()).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &LimitsResult, actual)
}

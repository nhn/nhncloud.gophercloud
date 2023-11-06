//go:build acceptance
// +build acceptance

package v2

import (
	"testing"

	"github.com/nhn/nhncloud.gophercloud/acceptance/clients"
	"github.com/nhn/nhncloud.gophercloud/acceptance/tools"
	"github.com/nhn/nhncloud.gophercloud/openstack/sharedfilesystems/v2/schedulerstats"
	th "github.com/nhn/nhncloud.gophercloud/testhelper"
)

func TestSchedulerStatsList(t *testing.T) {
	client, err := clients.NewSharedFileSystemV2Client()
	client.Microversion = "2.23"
	th.AssertNoErr(t, err)

	allPages, err := schedulerstats.List(client, nil).AllPages()
	th.AssertNoErr(t, err)

	allPools, err := schedulerstats.ExtractPools(allPages)
	th.AssertNoErr(t, err)

	for _, recordset := range allPools {
		tools.PrintResource(t, &recordset)
	}
}

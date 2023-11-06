//go:build acceptance || clustering || actions
// +build acceptance clustering actions

package v1

import (
	"testing"

	"github.com/nhn/nhncloud.gophercloud/acceptance/clients"
	"github.com/nhn/nhncloud.gophercloud/acceptance/tools"
	"github.com/nhn/nhncloud.gophercloud/openstack/clustering/v1/actions"
	th "github.com/nhn/nhncloud.gophercloud/testhelper"
)

func TestActionsList(t *testing.T) {
	client, err := clients.NewClusteringV1Client()
	th.AssertNoErr(t, err)

	opts := actions.ListOpts{
		Limit: 200,
	}

	allPages, err := actions.List(client, opts).AllPages()
	th.AssertNoErr(t, err)

	allActions, err := actions.ExtractActions(allPages)
	th.AssertNoErr(t, err)

	for _, action := range allActions {
		tools.PrintResource(t, action)
	}
}

package vpcsubnets

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type ListOptsBuilder interface {
	ToSubnetListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the subnet attributes you want to see returned. SortKey allows you to sort
// by a particular subnet attribute. SortDir sets the direction, and is either
// `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
	TenantID string `q:"tenant_id"`
	ID       string `q:"id"`
	Name     string `q:"name"`
	Shared   *bool  `q:"shared"`
	Limit    int    `q:"limit"`
	Marker   string `q:"marker"`
	SortKey  string `q:"sort_key"`
	SortDir  string `q:"sort_dir"`
}

func (opts ListOpts) ToSubnetListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// subnets. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
//
// Default policy settings return only those subnets that are owned by the tenant
// who submits the request, unless the request is submitted by a user with
// administrative rights.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToSubnetListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return SubnetPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
	resp, err := c.Get(getURL(c, id), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

type CreateOptsBuilder interface {
	ToSubnetCreateMap() (map[string]interface{}, error)
}

type CreateOpts struct {
	// VPC ID to which subnet is assigned
	VpcId string `json:"vpc_id" required:"true"`

	// CIDR of subnet
	CIDR string `json:"cidr" required:"true"`

	// Subnet name
	Name string `json:"name, required:"true""`

	// Tenant ID to which subnet is assigned
	TenantID string `json:"tenant_id,omitempty"`

	// Routingtalbe ID to be attached to or detached from the subnet
}

func (opts CreateOpts) ToSubnetCreateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "vpcsubnet")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToSubnetCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Post(createURL(c), b, &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

type UpdateOptsBuilder interface {
	ToSubnetUpdateMap() (map[string]interface{}, error)
}

type UpdateOpts struct {
	Name *string `json:"name" required:"true"`

	// RevisionNumber implements extension:standard-attr-revisions. If != "" it
	// will set revision_number=%s. If the revision number does not match, the
	// update will fail.
	RevisionNumber *int `json:"-" h:"If-Match"`
}

func (opts UpdateOpts) ToSubnetUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "vpcsubnet")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Update(c *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToSubnetUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	h, err := gophercloud.BuildHeaders(opts)
	if err != nil {
		r.Err = err
		return
	}
	for k := range h {
		if k == "If-Match" {
			h[k] = fmt.Sprintf("revision_number=%s", h[k])
		}
	}

	resp, err := c.Put(updateURL(c, id), b, &r.Body, &gophercloud.RequestOpts{
		MoreHeaders: h,
		OkCodes:     []int{200, 201},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

func Delete(c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	resp, err := c.Delete(deleteURL(c, id), nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

type AttachOptsBuilder interface {
	ToSubnetAttachMap() (map[string]interface{}, error)
}

type AttachOpts struct {
	RoutingtableID string `json:"routingtable_id" required:"true"`
}

func (opts AttachOpts) ToSubnetAttachMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Attach(c *gophercloud.ServiceClient, id string, opts AttachOptsBuilder) (r AttachResult) {
	b, err := opts.ToSubnetAttachMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Put(attachURL(c, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

func Detach(c *gophercloud.ServiceClient, id string) (r DetachResult) {
	resp, err := c.Put(detachURL(c, id), nil, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

package vpcsubnets

import (
	gophercloud "github.com/nhn/nhncloud.gophercloud"
	"github.com/nhn/nhncloud.gophercloud/pagination"
)

type commonResult struct {
	gophercloud.Result
}

// @tc-iaas-compute/1452
// LIST is only pre-flight, as we can take the result in more detail from GET
func (r commonResult) Extract() (*SubnetDetail, error) {
	var s struct {
		SubnetDetail *SubnetDetail `json:"vpcsubnet"`
	}
	err := r.ExtractInto(&s)
	return s.SubnetDetail, err
}

type CreateResult struct {
	commonResult
}

type GetResult struct {
	commonResult
}

type UpdateResult struct {
	commonResult
}

type DeleteResult struct {
	gophercloud.ErrResult
}

type AttachResult struct {
	gophercloud.ErrResult
}

type DetachResult struct {
	gophercloud.ErrResult
}

type Subnet struct {
	// Subnet name
	Name string `json:"name"`

	// Tenant ID that subnet belongs to
	TenantID string `json:"tenant_id"`

	// Subnet ID
	ID string `json:"id"`

	// Subnet status
	State string `json:"state"`

	// Created time for subnet
	CreateTime string `json:"create_time"`

	// ID of VPC that subnet belongs to
	VpcId string `json:"vpc_id"`

	// Whether to share subnet
	Shared bool `json:"shared"`

	// CIDR of subnet
	CIDR string `json:"cidr"`

	// Gateway IP of subnet
	Gateway string `json:"gateway:`
}

type SubnetDetail struct {
	External     bool   `json:"router:external,omitempty"`
	Name         string `json:"name,omitempty"`
	TenantID     string `json:"tenant_id,omitempty"`
	State        string `json:"state,omitempty"`
	ID           string `json:"id,omitempty"`
	Routingtable struct {
		GatewayID    string `json:"gateway_id,omitempty"`
		DefaultTable bool   `json:"default_table,omitempty"`
		Explicit     bool   `json:"explicit,omitempty"`
		ID           string `json:"id,omitempty"`
		Name         string `json:"name,omitempty"`
	} `json:"routingtable,omitempty"`
	CreateTime       string `json:"create_time,omitempty"`
	AvailableIPCount int    `json:"available_ip_count,omitempty"`
	VPC              struct {
		Shared bool   `json:"shared,omitempty"`
		State  string `json:"state,omitempty"`
		ID     string `json:"id,omitempty"`
		Cidrv4 string `json:"cidrv4,omitempty"`
		Name   string `json:"name,omitempty"`
	} `json:"vpc,omitempty"`
	VpcId  string `json:"vpc_id,omitempty"`
	Routes []struct {
		SubnetID string `json:"subnet_id,omitempty"`
		TenantID string `json:"tenant_id,omitempty"`
		Mask     int    `json:"mask,omitempty"`
		Gateway  string `json:"gateway,omitempty"`
		CIDR     string `json:"cidr,omitempty"`
		ID       string `json:"id,omitempty"`
	} `json:"routes,omitempty"`
	Shared  bool   `json:"shared,omitempty"`
	CIDR    string `json:"cidr,omitempty"`
	Gateway string `json:"gateway,omitempty"`
}

type SubnetPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of subnets has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r SubnetPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"subnets_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a SubnetPage struct is empty.
func (r SubnetPage) IsEmpty() (bool, error) {
	if r.StatusCode == 204 {
		return true, nil
	}

	is, err := ExtractSubnets(r)
	return len(is) == 0, err
}

// ExtractSubnets accepts a Page struct, specifically a SubnetPage struct,
// and extracts the elements into a slice of Subnet structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractSubnets(r pagination.Page) ([]Subnet, error) {
	var s struct {
		Subnets []Subnet `json:"vpcsubnets"`
	}
	err := (r.(SubnetPage)).ExtractInto(&s)
	return s.Subnets, err
}

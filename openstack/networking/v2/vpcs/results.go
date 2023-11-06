package vpcs

import (
	"encoding/json"
	"time"

	gophercloud "github.com/nhn/nhncloud.gophercloud"
	"github.com/nhn/nhncloud.gophercloud/pagination"
)

type commonResult struct {
	gophercloud.Result
}

func (r commonResult) Extract() (*Network, error) {
	var s Network
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "vpc")
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

type Network struct {
	// Name of VPC to query
	Name string `json:"name"`

	// Tenant ID to which VPC to query is included
	TenantID string `json:"tenant_id"`

	// Status of VPC to query
	State string `json:"state"`

	// Created time of VPC to query
	CreateTime string `json:"create_time"`

	// CIDR of VPC to query
	Cidrv4 string `json:"cidrv4"`

	// Whether to share VPC to query
	Shared bool `json:"shared"`

	// VPC ID to query
	ID string `json:"id"`

	// Whether network is externally connected
	External bool `json:"router:external"`

	// UpdatedAt and CreatedAt contain ISO-8601 timestamps of when the state of the
	// network last changed, and when it was created.
	UpdatedAt time.Time `json:"-"`
	CreatedAt time.Time `json:"-"`
}

type NetworkDetail struct {
	External      bool `json:"router:external,omitempty"`
	Routingtables []struct {
		DefaultTable bool   `json:"default_table,omitempty"`
		ID           string `json:"id,omitempty"`
		Name         string `json:"name,omitempty"`
	} `json:"routingtables,omitempty"`
	Name    string `json:"name,omitempty"`
	Subnets []struct {
		External   bool   `json:"router:external,omitempty"`
		Name       string `json:"name,omitempty"`
		EnableDHCP bool   `json:"enable_dhcp,omitempty"`
		TenantID   string `json:"tenant_id,omitempty"`
		Gateway    string `json:"gateway,omitempty"`
		Routes     []struct {
			SubnetID string `json:"subnet_id,omitempty"`
			TenantID string `json:"tenant_id,omitempty"`
			Mask     int    `json:"mask,omitempty"`
			Gateway  string `json:"gateway,omitempty"`
			CIDR     string `json:"cidr,omitempty"`
			ID       string `json:"id,omitempty"`
		} `json:"routes,omitempty"`
		State            string `json:"state,omitempty"`
		CreateTime       string `json:"create_time,omitempty"`
		AvailableIPCount int    `json:"available_ip_count,omitempty"`
		Routingtable     struct {
			GatewayID    string `json:"gateway_id,omitempty"`
			DefaultTable bool   `json:"default_table,omitempty"`
			Explicit     bool   `json:"explicit,omitempty"`
			ID           string `json:"id,omitempty"`
			Name         string `json:"name,omitempty"`
		} `json:"routingtable,omitempty"`
		VPC struct {
			Shared bool   `json:"shared,omitempty"`
			State  string `json:"state,omitempty"`
			ID     string `json:"id,omitempty"`
			Cidrv4 string `json:"cidrv4,omitempty"`
			Name   string `json:"name,omitempty"`
		} `json:"vpc,omitempty"`
		Shared bool   `json:"shared,omitempty"`
		ID     string `json:"id,omitempty"`
		VpcId  string `json:"vpc_id,omitempty"`
		CIDR   string `json:"cidr,omitempty"`
	} `json:"subnets,omitempty"`
	TenantID   string `json:"tenant_id,omitempty"`
	State      string `json:"state,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
	Cidrv4     string `json:"cidrv4,omitempty"`
	Shared     bool   `json:"shared,omitempty"`
	ID         string `json:"id,omitempty"`
}

func (r *Network) UnmarshalJSON(b []byte) error {
	type tmp Network

	// Support for older neutron time format
	var s1 struct {
		tmp
		CreatedAt gophercloud.JSONRFC3339NoZ `json:"created_at"`
		UpdatedAt gophercloud.JSONRFC3339NoZ `json:"updated_at"`
	}

	err := json.Unmarshal(b, &s1)
	if err == nil {
		*r = Network(s1.tmp)
		r.CreatedAt = time.Time(s1.CreatedAt)
		r.UpdatedAt = time.Time(s1.UpdatedAt)

		return nil
	}

	// Support for newer neutron time format
	var s2 struct {
		tmp
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	err = json.Unmarshal(b, &s2)
	if err != nil {
		return err
	}

	*r = Network(s2.tmp)
	r.CreatedAt = time.Time(s2.CreatedAt)
	r.UpdatedAt = time.Time(s2.UpdatedAt)

	return nil
}

// NetworkPage is the page returned by a pager when traversing over a
// collection of networks.
type NetworkPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of networks has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r NetworkPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"networks_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a NetworkPage struct is empty.
func (r NetworkPage) IsEmpty() (bool, error) {
	if r.StatusCode == 204 {
		return true, nil
	}

	is, err := ExtractNetworks(r)
	return len(is) == 0, err
}

// ExtractNetworks accepts a Page struct, specifically a NetworkPage struct,
// and extracts the elements into a slice of Network structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractNetworks(r pagination.Page) ([]Network, error) {
	var s []Network
	err := ExtractNetworksInto(r, &s)
	return s, err
}

func ExtractNetworksInto(r pagination.Page, v interface{}) error {
	return r.(NetworkPage).Result.ExtractIntoSlicePtr(v, "vpcs")
}

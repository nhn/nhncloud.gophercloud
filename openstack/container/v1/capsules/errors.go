package capsules

import (
	"fmt"

	gophercloud "github.com/nhn/nhncloud.gophercloud"
)

type ErrInvalidDataFormat struct {
	gophercloud.BaseError
}

func (e ErrInvalidDataFormat) Error() string {
	return fmt.Sprintf("Data in neither json nor yaml format.")
}

package pauseunpause

import gophercloud "github.com/nhn/nhncloud.gophercloud"

// PauseResult is the response from a Pause operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type PauseResult struct {
	gophercloud.ErrResult
}

// UnpauseResult is the response from an Unpause operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type UnpauseResult struct {
	gophercloud.ErrResult
}

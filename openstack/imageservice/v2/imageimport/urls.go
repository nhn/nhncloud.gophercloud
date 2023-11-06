package imageimport

import gophercloud "github.com/nhn/nhncloud.gophercloud"

const (
	rootPath     = "images"
	infoPath     = "info"
	resourcePath = "import"
)

func infoURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(infoPath, resourcePath)
}

func importURL(c *gophercloud.ServiceClient, imageID string) string {
	return c.ServiceURL(rootPath, imageID, resourcePath)
}

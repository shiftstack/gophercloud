package groups

import "github.com/gophercloud/gophercloud/v2"

const rootPath = "security-groups"

func rootURL(c gophercloud.Client) string {
	return c.ServiceURL(rootPath)
}

func resourceURL(c gophercloud.Client, id string) string {
	return c.ServiceURL(rootPath, id)
}

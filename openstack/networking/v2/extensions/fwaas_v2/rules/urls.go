package rules

import "github.com/gophercloud/gophercloud/v2"

const (
	rootPath     = "fwaas"
	resourcePath = "firewall_rules"
)

func rootURL(c gophercloud.Client) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func resourceURL(c gophercloud.Client, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}

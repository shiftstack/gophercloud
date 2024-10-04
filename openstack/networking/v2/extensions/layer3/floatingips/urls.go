package floatingips

import "github.com/gophercloud/gophercloud/v2"

const resourcePath = "floatingips"

func rootURL(c gophercloud.Client) string {
	return c.ServiceURL(resourcePath)
}

func resourceURL(c gophercloud.Client, id string) string {
	return c.ServiceURL(resourcePath, id)
}

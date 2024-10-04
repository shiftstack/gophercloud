package loadbalancers

import "github.com/gophercloud/gophercloud/v2"

const (
	rootPath       = "lbaas"
	resourcePath   = "loadbalancers"
	statusPath     = "status"
	statisticsPath = "stats"
	failoverPath   = "failover"
)

func rootURL(c gophercloud.Client) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func resourceURL(c gophercloud.Client, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}

func statusRootURL(c gophercloud.Client, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id, statusPath)
}

func statisticsRootURL(c gophercloud.Client, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id, statisticsPath)
}

func failoverRootURL(c gophercloud.Client, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id, failoverPath)
}

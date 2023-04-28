package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/identity/v3/osinherit"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

func TestAssign(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleAssignSuccessfully(t)

	err := osinherit.Assign(client.ServiceClient(), "{role_id}", osinherit.AssignOpts{
		UserID:    "{user_id}",
		ProjectID: "{project_id}",
	}).ExtractErr()
	th.AssertNoErr(t, err)

	err = osinherit.Assign(client.ServiceClient(), "{role_id}", osinherit.AssignOpts{
		UserID:   "{user_id}",
		DomainID: "{domain_id}",
	}).ExtractErr()
	th.AssertNoErr(t, err)

	err = osinherit.Assign(client.ServiceClient(), "{role_id}", osinherit.AssignOpts{
		GroupID:   "{group_id}",
		ProjectID: "{project_id}",
	}).ExtractErr()
	th.AssertNoErr(t, err)

	err = osinherit.Assign(client.ServiceClient(), "{role_id}", osinherit.AssignOpts{
		GroupID:  "{group_id}",
		DomainID: "{domain_id}",
	}).ExtractErr()
	th.AssertNoErr(t, err)
}

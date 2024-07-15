package subnets

import (
	"context"
	"fmt"
	"os"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack"
)

func getNetworkClient() *gophercloud.ServiceClient {
	ctx := context.Background()

	opts, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		panic(err)
	}

	providerClient, err := openstack.AuthenticatedClient(ctx, opts)
	if err != nil {
		panic(err)
	}

	networkClient, err := openstack.NewComputeV2(providerClient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if err != nil {
		panic(err)
	}
	return networkClient
}

func ExampleList() {
	networkClient := getNetworkClient()
	listOpts := ListOpts{
		IPVersion: 4,
	}

	allPages, err := List(networkClient, listOpts).AllPages(context.TODO())
	if err != nil {
		panic(err)
	}

	allSubnets, err := ExtractSubnets(allPages)
	if err != nil {
		panic(err)
	}

	for _, subnet := range allSubnets {
		fmt.Printf("%+v\n", subnet)
	}
}

func ExampleCreate() {
	networkClient := getNetworkClient()
	var gatewayIP = "192.168.199.1"
	createOpts := CreateOpts{
		NetworkID: "d32019d3-bc6e-4319-9c1d-6722fc136a22",
		IPVersion: 4,
		CIDR:      "192.168.199.0/24",
		GatewayIP: &gatewayIP,
		AllocationPools: []AllocationPool{
			{
				Start: "192.168.199.2",
				End:   "192.168.199.254",
			},
		},
		DNSNameservers: []string{"foo"},
		ServiceTypes:   []string{"network:floatingip"},
	}

	subnet, err := Create(context.TODO(), networkClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", subnet)
}

func ExampleCreate_noGateway() {
	networkClient := getNetworkClient()
	var noGateway = ""
	createOpts := CreateOpts{
		NetworkID: "d32019d3-bc6e-4319-9c1d-6722fc136a23",
		IPVersion: 4,
		CIDR:      "192.168.1.0/24",
		GatewayIP: &noGateway,
		AllocationPools: []AllocationPool{
			{
				Start: "192.168.1.2",
				End:   "192.168.1.254",
			},
		},
		DNSNameservers: []string{},
	}

	subnet, err := Create(context.TODO(), networkClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", subnet)
}

func ExampleCreate_defaultGateway() {
	networkClient := getNetworkClient()
	createOpts := CreateOpts{
		NetworkID: "d32019d3-bc6e-4319-9c1d-6722fc136a23",
		IPVersion: 4,
		CIDR:      "192.168.1.0/24",
		AllocationPools: []AllocationPool{
			{
				Start: "192.168.1.2",
				End:   "192.168.1.254",
			},
		},
		DNSNameservers: []string{},
	}

	subnet, err := Create(context.TODO(), networkClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", subnet)
}

func ExampleUpdate() {
	networkClient := getNetworkClient()
	subnetID := "db77d064-e34f-4d06-b060-f21e28a61c23"
	dnsNameservers := []string{"8.8.8.8"}
	serviceTypes := []string{"network:floatingip", "network:routed"}
	name := "new_name"

	updateOpts := UpdateOpts{
		Name:           &name,
		DNSNameservers: &dnsNameservers,
		ServiceTypes:   &serviceTypes,
	}

	subnet, err := Update(context.TODO(), networkClient, subnetID, updateOpts).Extract()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", subnet)
}

func ExampleUpdate_removeGateway() {
	networkClient := getNetworkClient()
	var noGateway = ""
	subnetID := "db77d064-e34f-4d06-b060-f21e28a61c23"

	updateOpts := UpdateOpts{
		GatewayIP: &noGateway,
	}

	subnet, err := Update(context.TODO(), networkClient, subnetID, updateOpts).Extract()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", subnet)
}

func ExampleDelete() {
	networkClient := getNetworkClient()
	subnetID := "db77d064-e34f-4d06-b060-f21e28a61c23"
	err := Delete(context.TODO(), networkClient, subnetID).ExtractErr()
	if err != nil {
		panic(err)
	}
}

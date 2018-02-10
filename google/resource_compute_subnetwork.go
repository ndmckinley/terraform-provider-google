// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by terraform-codegen and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in README.md and
//     CONTRIBUTING.md located at the root of this package.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	compute "google.golang.org/api/compute/v1"
)

func resourceComputeSubnetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeSubnetworkCreate,
		Read:   resourceComputeSubnetworkRead,
		Update: resourceComputeSubnetworkUpdate,
		Delete: resourceComputeSubnetworkDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeSubnetworkImport,
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"gateway_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_cidr_range": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"private_ip_google_access": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"creation_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"project": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceComputeSubnetworkCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := map[string]interface{}{
		"description":           expandComputeSubnetworkDescription(d.Get("description")),
		"gatewayAddress":        expandComputeSubnetworkGatewayAddress(d.Get("gateway_address")),
		"ipCidrRange":           expandComputeSubnetworkIpCidrRange(d.Get("ip_cidr_range")),
		"name":                  expandComputeSubnetworkName(d.Get("name")),
		"network":               expandComputeSubnetworkNetwork(d.Get("network")),
		"privateIpGoogleAccess": expandComputeSubnetworkPrivateIpGoogleAccess(d.Get("private_ip_google_access")),
		"region":                expandComputeSubnetworkRegion(d.Get("region")),
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/regions/{{region}}/subnetworks")
	if err != nil {
		return err
	}
	res, err := Post(config, url, obj)
	if err != nil {
		return fmt.Errorf("Error creating Subnetwork: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{region}}/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	waitErr := computeOperationWait(config.clientCompute, op, project, "Creating Subnetwork")
	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return waitErr
	}

	return resourceComputeSubnetworkRead(d, meta)
}

func resourceComputeSubnetworkRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/regions/{{region}}/subnetworks/{{name}}")
	if err != nil {
		return err
	}
	res, err := Get(config, url)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeSubnetwork %q", d.Id()))
	}

	d.Set("creation_timestamp", flattenComputeSubnetworkCreationTimestamp(res["creationTimestamp"]))
	d.Set("description", flattenComputeSubnetworkDescription(res["description"]))
	d.Set("gateway_address", flattenComputeSubnetworkGatewayAddress(res["gatewayAddress"]))
	d.Set("ip_cidr_range", flattenComputeSubnetworkIpCidrRange(res["ipCidrRange"]))
	d.Set("name", flattenComputeSubnetworkName(res["name"]))
	d.Set("network", flattenComputeSubnetworkNetwork(res["network"]))
	d.Set("private_ip_google_access", flattenComputeSubnetworkPrivateIpGoogleAccess(res["privateIpGoogleAccess"]))
	d.Set("region", flattenComputeSubnetworkRegion(res["region"]))
	d.Set("self_link", res["selfLink"])
	d.Set("project", project)

	return nil
}

func resourceComputeSubnetworkUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := map[string]interface{}{
		"description":           expandComputeSubnetworkDescription(d.Get("description")),
		"gatewayAddress":        expandComputeSubnetworkGatewayAddress(d.Get("gateway_address")),
		"ipCidrRange":           expandComputeSubnetworkIpCidrRange(d.Get("ip_cidr_range")),
		"name":                  expandComputeSubnetworkName(d.Get("name")),
		"network":               expandComputeSubnetworkNetwork(d.Get("network")),
		"privateIpGoogleAccess": expandComputeSubnetworkPrivateIpGoogleAccess(d.Get("private_ip_google_access")),
		"region":                expandComputeSubnetworkRegion(d.Get("region")),
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/regions/{{region}}/subnetworks/{{name}}")
	if err != nil {
		return err
	}
	res, err := Put(config, url, obj)
	if err != nil {
		return fmt.Errorf("Error updating Subnetwork %q: %s", d.Id(), err)
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWait(config.clientCompute, op, project, "Updating Subnetwork")
	if err != nil {
		return err
	}

	return resourceComputeSubnetworkRead(d, meta)
}

func resourceComputeSubnetworkDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/regions/{{region}}/subnetworks/{{name}}")
	if err != nil {
		return err
	}
	res, err := Delete(config, url)
	if err != nil {
		return fmt.Errorf("Error deleting Subnetwork %q: %s", d.Id(), err)
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWait(config.clientCompute, op, project, "Updating Subnetwork")
	if err != nil {
		return err
	}

	return nil
}

func resourceComputeSubnetworkImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("{{region}}/{{name}}", d.Id())
	return []*schema.ResourceData{d}, nil
}

func flattenComputeSubnetworkCreationTimestamp(v interface{}) interface{} {
	return v
}

func flattenComputeSubnetworkDescription(v interface{}) interface{} {
	return v
}

func flattenComputeSubnetworkGatewayAddress(v interface{}) interface{} {
	return v
}

func flattenComputeSubnetworkIpCidrRange(v interface{}) interface{} {
	return v
}

func flattenComputeSubnetworkName(v interface{}) interface{} {
	return v
}

func flattenComputeSubnetworkNetwork(v interface{}) interface{} {
	return v
}

func flattenComputeSubnetworkPrivateIpGoogleAccess(v interface{}) interface{} {
	return v
}

func flattenComputeSubnetworkRegion(v interface{}) interface{} {
	return v
}

func expandComputeSubnetworkDescription(v interface{}) interface{} {
	return v
}

func expandComputeSubnetworkGatewayAddress(v interface{}) interface{} {
	return v
}

func expandComputeSubnetworkIpCidrRange(v interface{}) interface{} {
	return v
}

func expandComputeSubnetworkName(v interface{}) interface{} {
	return v
}

func expandComputeSubnetworkNetwork(v interface{}) interface{} {
	return v
}

func expandComputeSubnetworkPrivateIpGoogleAccess(v interface{}) interface{} {
	return v
}

func expandComputeSubnetworkRegion(v interface{}) interface{} {
	return v
}

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

func resourceComputeRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeRouteCreate,
		Read:   resourceComputeRouteRead,
		Update: resourceComputeRouteUpdate,
		Delete: resourceComputeRouteDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeRouteImport,
		},

		Schema: map[string]*schema.Schema{
			"dest_range": {
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
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"next_hop_gateway": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"next_hop_instance": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"next_hop_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"next_hop_vpn_tunnel": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"project": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceComputeRouteCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := map[string]interface{}{
		"destRange":        expandComputeRouteDestRange(d.Get("dest_range")),
		"name":             expandComputeRouteName(d.Get("name")),
		"network":          expandComputeRouteNetwork(d.Get("network")),
		"priority":         expandComputeRoutePriority(d.Get("priority")),
		"tags":             expandComputeRouteTags(d.Get("tags")),
		"nextHopGateway":   expandComputeRouteNextHopGateway(d.Get("next_hop_gateway")),
		"nextHopInstance":  expandComputeRouteNextHopInstance(d.Get("next_hop_instance")),
		"nextHopIp":        expandComputeRouteNextHopIp(d.Get("next_hop_ip")),
		"nextHopVpnTunnel": expandComputeRouteNextHopVpnTunnel(d.Get("next_hop_vpn_tunnel")),
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/routes")
	if err != nil {
		return err
	}
	res, err := Post(config, url, obj)
	if err != nil {
		return fmt.Errorf("Error creating Route: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	waitErr := computeOperationWait(config.clientCompute, op, project, "Creating Route")
	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return waitErr
	}

	return resourceComputeRouteRead(d, meta)
}

func resourceComputeRouteRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/routes/{{name}}")
	if err != nil {
		return err
	}
	res, err := Get(config, url)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeRoute %q", d.Id()))
	}

	d.Set("dest_range", flattenComputeRouteDestRange(res["destRange"]))
	d.Set("name", flattenComputeRouteName(res["name"]))
	d.Set("network", flattenComputeRouteNetwork(res["network"]))
	d.Set("priority", flattenComputeRoutePriority(res["priority"]))
	d.Set("tags", flattenComputeRouteTags(res["tags"]))
	d.Set("next_hop_gateway", flattenComputeRouteNextHopGateway(res["nextHopGateway"]))
	d.Set("next_hop_instance", flattenComputeRouteNextHopInstance(res["nextHopInstance"]))
	d.Set("next_hop_ip", flattenComputeRouteNextHopIp(res["nextHopIp"]))
	d.Set("next_hop_vpn_tunnel", flattenComputeRouteNextHopVpnTunnel(res["nextHopVpnTunnel"]))
	d.Set("self_link", res["selfLink"])
	d.Set("project", project)

	return nil
}

func resourceComputeRouteUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := map[string]interface{}{
		"destRange":        expandComputeRouteDestRange(d.Get("dest_range")),
		"name":             expandComputeRouteName(d.Get("name")),
		"network":          expandComputeRouteNetwork(d.Get("network")),
		"priority":         expandComputeRoutePriority(d.Get("priority")),
		"tags":             expandComputeRouteTags(d.Get("tags")),
		"nextHopGateway":   expandComputeRouteNextHopGateway(d.Get("next_hop_gateway")),
		"nextHopInstance":  expandComputeRouteNextHopInstance(d.Get("next_hop_instance")),
		"nextHopIp":        expandComputeRouteNextHopIp(d.Get("next_hop_ip")),
		"nextHopVpnTunnel": expandComputeRouteNextHopVpnTunnel(d.Get("next_hop_vpn_tunnel")),
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/routes/{{name}}")
	if err != nil {
		return err
	}
	res, err := Put(config, url, obj)
	if err != nil {
		return fmt.Errorf("Error updating Route %q: %s", d.Id(), err)
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWait(config.clientCompute, op, project, "Updating Route")
	if err != nil {
		return err
	}

	return resourceComputeRouteRead(d, meta)
}

func resourceComputeRouteDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/routes/{{name}}")
	if err != nil {
		return err
	}
	res, err := Delete(config, url)
	if err != nil {
		return fmt.Errorf("Error deleting Route %q: %s", d.Id(), err)
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWait(config.clientCompute, op, project, "Updating Route")
	if err != nil {
		return err
	}

	return nil
}

func resourceComputeRouteImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("{{name}}", d.Id())
	return []*schema.ResourceData{d}, nil
}

func flattenComputeRouteDestRange(v interface{}) interface{} {
	return v
}

func flattenComputeRouteName(v interface{}) interface{} {
	return v
}

func flattenComputeRouteNetwork(v interface{}) interface{} {
	return v
}

func flattenComputeRoutePriority(v interface{}) interface{} {
	return v
}

func flattenComputeRouteTags(v interface{}) interface{} {
	return v
}

func flattenComputeRouteNextHopGateway(v interface{}) interface{} {
	return v
}

func flattenComputeRouteNextHopInstance(v interface{}) interface{} {
	return v
}

func flattenComputeRouteNextHopIp(v interface{}) interface{} {
	return v
}

func flattenComputeRouteNextHopVpnTunnel(v interface{}) interface{} {
	return v
}

func expandComputeRouteDestRange(v interface{}) interface{} {
	return v
}

func expandComputeRouteName(v interface{}) interface{} {
	return v
}

func expandComputeRouteNetwork(v interface{}) interface{} {
	return v
}

func expandComputeRoutePriority(v interface{}) interface{} {
	return v
}

func expandComputeRouteTags(v interface{}) interface{} {
	return v
}

func expandComputeRouteNextHopGateway(v interface{}) interface{} {
	return v
}

func expandComputeRouteNextHopInstance(v interface{}) interface{} {
	return v
}

func expandComputeRouteNextHopIp(v interface{}) interface{} {
	return v
}

func expandComputeRouteNextHopVpnTunnel(v interface{}) interface{} {
	return v
}

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated and manual changes will be
//     clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	compute "google.golang.org/api/compute/v1"
)

func resourceComputeGlobalForwardingRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeGlobalForwardingRuleCreate,
		Read:   resourceComputeGlobalForwardingRuleRead,
		Update: resourceComputeGlobalForwardingRuleUpdate,
		Delete: resourceComputeGlobalForwardingRuleDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeGlobalForwardingRuleImport,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_protocol": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"backend_service": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"ip_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"load_balancing_scheme": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"port_range": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ports": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"subnetwork": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"target": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"creation_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
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

func resourceComputeGlobalForwardingRuleCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := map[string]interface{}{
		"description":         expandComputeGlobalForwardingRuleDescription(d.Get("description")),
		"IPAddress":           expandComputeGlobalForwardingRuleIPAddress(d.Get("ip_address")),
		"IPProtocol":          expandComputeGlobalForwardingRuleIPProtocol(d.Get("ip_protocol")),
		"backendService":      expandComputeGlobalForwardingRuleBackendService(d.Get("backend_service")),
		"ipVersion":           expandComputeGlobalForwardingRuleIpVersion(d.Get("ip_version")),
		"loadBalancingScheme": expandComputeGlobalForwardingRuleLoadBalancingScheme(d.Get("load_balancing_scheme")),
		"name":                expandComputeGlobalForwardingRuleName(d.Get("name")),
		"network":             expandComputeGlobalForwardingRuleNetwork(d.Get("network")),
		"portRange":           expandComputeGlobalForwardingRulePortRange(d.Get("port_range")),
		"ports":               expandComputeGlobalForwardingRulePorts(d.Get("ports")),
		"subnetwork":          expandComputeGlobalForwardingRuleSubnetwork(d.Get("subnetwork")),
		"target":              expandComputeGlobalForwardingRuleTarget(d.Get("target")),
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/forwardingRules")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new GlobalForwardingRule: %#v", obj)
	res, err := Post(config, url, obj)
	if err != nil {
		return fmt.Errorf("Error creating GlobalForwardingRule: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	waitErr := computeOperationWait(config.clientCompute, op, project, "Creating GlobalForwardingRule")
	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return waitErr
	}

	return resourceComputeGlobalForwardingRuleRead(d, meta)
}

func resourceComputeGlobalForwardingRuleRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/forwardingRules/{{name}}")
	if err != nil {
		return err
	}

	res, err := Get(config, url)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeGlobalForwardingRule %q", d.Id()))
	}

	d.Set("creation_timestamp", flattenComputeGlobalForwardingRuleCreationTimestamp(res["creationTimestamp"]))
	d.Set("description", flattenComputeGlobalForwardingRuleDescription(res["description"]))
	d.Set("id", flattenComputeGlobalForwardingRuleId(res["id"]))
	d.Set("ip_address", flattenComputeGlobalForwardingRuleIPAddress(res["IPAddress"]))
	d.Set("ip_protocol", flattenComputeGlobalForwardingRuleIPProtocol(res["IPProtocol"]))
	d.Set("backend_service", flattenComputeGlobalForwardingRuleBackendService(res["backendService"]))
	d.Set("ip_version", flattenComputeGlobalForwardingRuleIpVersion(res["ipVersion"]))
	d.Set("load_balancing_scheme", flattenComputeGlobalForwardingRuleLoadBalancingScheme(res["loadBalancingScheme"]))
	d.Set("name", flattenComputeGlobalForwardingRuleName(res["name"]))
	d.Set("network", flattenComputeGlobalForwardingRuleNetwork(res["network"]))
	d.Set("port_range", flattenComputeGlobalForwardingRulePortRange(res["portRange"]))
	d.Set("ports", flattenComputeGlobalForwardingRulePorts(res["ports"]))
	d.Set("subnetwork", flattenComputeGlobalForwardingRuleSubnetwork(res["subnetwork"]))
	d.Set("region", flattenComputeGlobalForwardingRuleRegion(res["region"]))
	d.Set("target", flattenComputeGlobalForwardingRuleTarget(res["target"]))
	d.Set("self_link", res["selfLink"])
	d.Set("project", project)

	return nil
}

func resourceComputeGlobalForwardingRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := map[string]interface{}{
		"description":         expandComputeGlobalForwardingRuleDescription(d.Get("description")),
		"IPAddress":           expandComputeGlobalForwardingRuleIPAddress(d.Get("ip_address")),
		"IPProtocol":          expandComputeGlobalForwardingRuleIPProtocol(d.Get("ip_protocol")),
		"backendService":      expandComputeGlobalForwardingRuleBackendService(d.Get("backend_service")),
		"ipVersion":           expandComputeGlobalForwardingRuleIpVersion(d.Get("ip_version")),
		"loadBalancingScheme": expandComputeGlobalForwardingRuleLoadBalancingScheme(d.Get("load_balancing_scheme")),
		"name":                expandComputeGlobalForwardingRuleName(d.Get("name")),
		"network":             expandComputeGlobalForwardingRuleNetwork(d.Get("network")),
		"portRange":           expandComputeGlobalForwardingRulePortRange(d.Get("port_range")),
		"ports":               expandComputeGlobalForwardingRulePorts(d.Get("ports")),
		"subnetwork":          expandComputeGlobalForwardingRuleSubnetwork(d.Get("subnetwork")),
		"target":              expandComputeGlobalForwardingRuleTarget(d.Get("target")),
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/forwardingRules/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating GlobalForwardingRule %q: %#v", d.Id(), obj)
	res, err := sendRequest(config, "PUT", url, obj)

	if err != nil {
		return fmt.Errorf("Error updating GlobalForwardingRule %q: %s", d.Id(), err)
	}
	err = Convert(res, &op)
	if err != nil {
		return err
	}

	err = computeOperationWait(config.clientCompute, &op, project, "Updating GlobalForwardingRule")
	if err != nil {
		return err
	}

	return resourceComputeGlobalForwardingRuleRead(d, meta)
}

func resourceComputeGlobalForwardingRuleDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/forwardingRules/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting GlobalForwardingRule %q", d.Id())
	res, err := Delete(config, url)
	if err != nil {
		return fmt.Errorf("Error deleting GlobalForwardingRule %q: %s", d.Id(), err)
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWait(config.clientCompute, op, project, "Deleting GlobalForwardingRule")
	if err != nil {
		return err
	}

	return nil
}

func resourceComputeGlobalForwardingRuleImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("", d.Id())
	return []*schema.ResourceData{d}, nil
}

func flattenComputeGlobalForwardingRuleCreationTimestamp(v interface{}) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleDescription(v interface{}) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleId(v interface{}) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleIPAddress(v interface{}) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleIPProtocol(v interface{}) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleBackendService(v interface{}) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleIpVersion(v interface{}) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleLoadBalancingScheme(v interface{}) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleName(v interface{}) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleNetwork(v interface{}) interface{} {
	return v
}

func flattenComputeGlobalForwardingRulePortRange(v interface{}) interface{} {
	return v
}

func flattenComputeGlobalForwardingRulePorts(v interface{}) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleSubnetwork(v interface{}) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleRegion(v interface{}) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleTarget(v interface{}) interface{} {
	return v
}

func expandComputeGlobalForwardingRuleDescription(v interface{}) interface{} {
	return v
}

func expandComputeGlobalForwardingRuleIPAddress(v interface{}) interface{} {
	return v
}

func expandComputeGlobalForwardingRuleIPProtocol(v interface{}) interface{} {
	return v
}

func expandComputeGlobalForwardingRuleBackendService(v interface{}) interface{} {
	return v
}

func expandComputeGlobalForwardingRuleIpVersion(v interface{}) interface{} {
	return v
}

func expandComputeGlobalForwardingRuleLoadBalancingScheme(v interface{}) interface{} {
	return v
}

func expandComputeGlobalForwardingRuleName(v interface{}) interface{} {
	return v
}

func expandComputeGlobalForwardingRuleNetwork(v interface{}) interface{} {
	return v
}

func expandComputeGlobalForwardingRulePortRange(v interface{}) interface{} {
	return v
}

func expandComputeGlobalForwardingRulePorts(v interface{}) interface{} {
	return v
}

func expandComputeGlobalForwardingRuleSubnetwork(v interface{}) interface{} {
	return v
}

func expandComputeGlobalForwardingRuleTarget(v interface{}) interface{} {
	return v
}

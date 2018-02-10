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

func resourceComputeInstanceGroupManager() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeInstanceGroupManagerCreate,
		Read:   resourceComputeInstanceGroupManagerRead,
		Update: resourceComputeInstanceGroupManagerUpdate,
		Delete: resourceComputeInstanceGroupManagerDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeInstanceGroupManagerImport,
		},

		Schema: map[string]*schema.Schema{
			"base_instance_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"instance_template": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"zone": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"named_ports": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"target_pools": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					DiffSuppressFunc: compareSelfLinkOrResourceName,
				},
			},
			"target_size": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"creation_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"current_actions": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"abandoning": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"creating": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"creating_without_retries": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"deleting": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"none": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"recreating": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"refreshing": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"restarting": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"instance_group": {
				Type:             schema.TypeString,
				Computed:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
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

func resourceComputeInstanceGroupManagerCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := map[string]interface{}{
		"baseInstanceName": expandComputeInstanceGroupManagerBaseInstanceName(d.Get("base_instance_name")),
		"description":      expandComputeInstanceGroupManagerDescription(d.Get("description")),
		"instanceTemplate": expandComputeInstanceGroupManagerInstanceTemplate(d.Get("instance_template")),
		"name":             expandComputeInstanceGroupManagerName(d.Get("name")),
		"namedPorts":       expandComputeInstanceGroupManagerNamedPorts(d.Get("named_ports")),
		"targetPools":      expandComputeInstanceGroupManagerTargetPools(d.Get("target_pools")),
		"targetSize":       expandComputeInstanceGroupManagerTargetSize(d.Get("target_size")),
		"zone":             expandComputeInstanceGroupManagerZone(d.Get("zone")),
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/zones/{{zone}}/instanceGroupManagers")
	if err != nil {
		return err
	}
	res, err := Post(config, url, obj)
	if err != nil {
		return fmt.Errorf("Error creating InstanceGroupManager: %s", err)
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

	waitErr := computeOperationWait(config.clientCompute, op, project, "Creating InstanceGroupManager")
	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return waitErr
	}

	return resourceComputeInstanceGroupManagerRead(d, meta)
}

func resourceComputeInstanceGroupManagerRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/zones/{{zone}}/instanceGroupManagers/{{name}}")
	if err != nil {
		return err
	}
	res, err := Get(config, url)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeInstanceGroupManager %q", d.Id()))
	}

	d.Set("base_instance_name", flattenComputeInstanceGroupManagerBaseInstanceName(res["baseInstanceName"]))
	d.Set("creation_timestamp", flattenComputeInstanceGroupManagerCreationTimestamp(res["creationTimestamp"]))
	d.Set("current_actions", flattenComputeInstanceGroupManagerCurrentActions(res["currentActions"]))
	d.Set("description", flattenComputeInstanceGroupManagerDescription(res["description"]))
	d.Set("id", flattenComputeInstanceGroupManagerId(res["id"]))
	d.Set("instance_group", flattenComputeInstanceGroupManagerInstanceGroup(res["instanceGroup"]))
	d.Set("instance_template", flattenComputeInstanceGroupManagerInstanceTemplate(res["instanceTemplate"]))
	d.Set("name", flattenComputeInstanceGroupManagerName(res["name"]))
	d.Set("named_ports", flattenComputeInstanceGroupManagerNamedPorts(res["namedPorts"]))
	d.Set("region", flattenComputeInstanceGroupManagerRegion(res["region"]))
	d.Set("target_pools", flattenComputeInstanceGroupManagerTargetPools(res["targetPools"]))
	d.Set("target_size", flattenComputeInstanceGroupManagerTargetSize(res["targetSize"]))
	d.Set("zone", flattenComputeInstanceGroupManagerZone(res["zone"]))
	d.Set("self_link", res["selfLink"])
	d.Set("project", project)

	return nil
}

func resourceComputeInstanceGroupManagerUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := map[string]interface{}{
		"baseInstanceName": expandComputeInstanceGroupManagerBaseInstanceName(d.Get("base_instance_name")),
		"description":      expandComputeInstanceGroupManagerDescription(d.Get("description")),
		"instanceTemplate": expandComputeInstanceGroupManagerInstanceTemplate(d.Get("instance_template")),
		"name":             expandComputeInstanceGroupManagerName(d.Get("name")),
		"namedPorts":       expandComputeInstanceGroupManagerNamedPorts(d.Get("named_ports")),
		"targetPools":      expandComputeInstanceGroupManagerTargetPools(d.Get("target_pools")),
		"targetSize":       expandComputeInstanceGroupManagerTargetSize(d.Get("target_size")),
		"zone":             expandComputeInstanceGroupManagerZone(d.Get("zone")),
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/zones/{{zone}}/instanceGroupManagers/{{name}}")
	if err != nil {
		return err
	}
	res, err := Put(config, url, obj)
	if err != nil {
		return fmt.Errorf("Error updating InstanceGroupManager %q: %s", d.Id(), err)
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWait(config.clientCompute, op, project, "Updating InstanceGroupManager")
	if err != nil {
		return err
	}

	return resourceComputeInstanceGroupManagerRead(d, meta)
}

func resourceComputeInstanceGroupManagerDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/zones/{{zone}}/instanceGroupManagers/{{name}}")
	if err != nil {
		return err
	}
	res, err := Delete(config, url)
	if err != nil {
		return fmt.Errorf("Error deleting InstanceGroupManager %q: %s", d.Id(), err)
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWait(config.clientCompute, op, project, "Updating InstanceGroupManager")
	if err != nil {
		return err
	}

	return nil
}

func resourceComputeInstanceGroupManagerImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("", d.Id())
	return []*schema.ResourceData{d}, nil
}

func flattenComputeInstanceGroupManagerBaseInstanceName(v interface{}) interface{} {
	return v
}

func flattenComputeInstanceGroupManagerCreationTimestamp(v interface{}) interface{} {
	return v
}

func flattenComputeInstanceGroupManagerCurrentActions(v interface{}) interface{} {
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})
	transformed["abandoning"] =
		flattenComputeInstanceGroupManagerCurrentActionsAbandoning(original["abandoning"])
	transformed["creating"] =
		flattenComputeInstanceGroupManagerCurrentActionsCreating(original["creating"])
	transformed["creating_without_retries"] =
		flattenComputeInstanceGroupManagerCurrentActionsCreatingWithoutRetries(original["creatingWithoutRetries"])
	transformed["deleting"] =
		flattenComputeInstanceGroupManagerCurrentActionsDeleting(original["deleting"])
	transformed["none"] =
		flattenComputeInstanceGroupManagerCurrentActionsNone(original["none"])
	transformed["recreating"] =
		flattenComputeInstanceGroupManagerCurrentActionsRecreating(original["recreating"])
	transformed["refreshing"] =
		flattenComputeInstanceGroupManagerCurrentActionsRefreshing(original["refreshing"])
	transformed["restarting"] =
		flattenComputeInstanceGroupManagerCurrentActionsRestarting(original["restarting"])
	return []interface{}{transformed}
}
func flattenComputeInstanceGroupManagerCurrentActionsAbandoning(v interface{}) interface{} {
	return v
}
func flattenComputeInstanceGroupManagerCurrentActionsCreating(v interface{}) interface{} {
	return v
}
func flattenComputeInstanceGroupManagerCurrentActionsCreatingWithoutRetries(v interface{}) interface{} {
	return v
}
func flattenComputeInstanceGroupManagerCurrentActionsDeleting(v interface{}) interface{} {
	return v
}
func flattenComputeInstanceGroupManagerCurrentActionsNone(v interface{}) interface{} {
	return v
}
func flattenComputeInstanceGroupManagerCurrentActionsRecreating(v interface{}) interface{} {
	return v
}
func flattenComputeInstanceGroupManagerCurrentActionsRefreshing(v interface{}) interface{} {
	return v
}
func flattenComputeInstanceGroupManagerCurrentActionsRestarting(v interface{}) interface{} {
	return v
}

func flattenComputeInstanceGroupManagerDescription(v interface{}) interface{} {
	return v
}

func flattenComputeInstanceGroupManagerId(v interface{}) interface{} {
	return v
}

func flattenComputeInstanceGroupManagerInstanceGroup(v interface{}) interface{} {
	return v
}

func flattenComputeInstanceGroupManagerInstanceTemplate(v interface{}) interface{} {
	return v
}

func flattenComputeInstanceGroupManagerName(v interface{}) interface{} {
	return v
}

func flattenComputeInstanceGroupManagerNamedPorts(v interface{}) interface{} {
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		transformed = append(transformed, map[string]interface{}{
			"name": flattenComputeInstanceGroupManagerNamedPortsName(original["name"]),
			"port": flattenComputeInstanceGroupManagerNamedPortsPort(original["port"]),
		})
	}
	return transformed
}
func flattenComputeInstanceGroupManagerNamedPortsName(v interface{}) interface{} {
	return v
}
func flattenComputeInstanceGroupManagerNamedPortsPort(v interface{}) interface{} {
	return v
}

func flattenComputeInstanceGroupManagerRegion(v interface{}) interface{} {
	return v
}

func flattenComputeInstanceGroupManagerTargetPools(v interface{}) interface{} {
	return v
}

func flattenComputeInstanceGroupManagerTargetSize(v interface{}) interface{} {
	return v
}

func flattenComputeInstanceGroupManagerZone(v interface{}) interface{} {
	return v
}

func expandComputeInstanceGroupManagerBaseInstanceName(v interface{}) interface{} {
	return v
}

func expandComputeInstanceGroupManagerDescription(v interface{}) interface{} {
	return v
}

func expandComputeInstanceGroupManagerInstanceTemplate(v interface{}) interface{} {
	return v
}

func expandComputeInstanceGroupManagerName(v interface{}) interface{} {
	return v
}

func expandComputeInstanceGroupManagerNamedPorts(v interface{}) interface{} {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformed["name"] =
			expandComputeInstanceGroupManagerNamedPortsName(original["name"])
		transformed["port"] =
			expandComputeInstanceGroupManagerNamedPortsPort(original["port"])

		req = append(req, transformed)
	}
	return req
}
func expandComputeInstanceGroupManagerNamedPortsName(v interface{}) interface{} {
	return v
}
func expandComputeInstanceGroupManagerNamedPortsPort(v interface{}) interface{} {
	return v
}

func expandComputeInstanceGroupManagerTargetPools(v interface{}) interface{} {
	return v
}

func expandComputeInstanceGroupManagerTargetSize(v interface{}) interface{} {
	return v
}

func expandComputeInstanceGroupManagerZone(v interface{}) interface{} {
	return v
}

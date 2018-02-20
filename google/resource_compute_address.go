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

func resourceComputeAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeAddressCreate,
		Read:   resourceComputeAddressRead,
		Update: resourceComputeAddressUpdate,
		Delete: resourceComputeAddressDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeAddressImport,
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
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
			"users": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

func resourceComputeAddressCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := map[string]interface{}{
		"address":     expandComputeAddressAddress(d.Get("address")),
		"description": expandComputeAddressDescription(d.Get("description")),
		"name":        expandComputeAddressName(d.Get("name")),
		"region":      expandComputeAddressRegion(d.Get("region")),
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/regions/{{region}}/addresses")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Address: %#v", obj)
	res, err := Post(config, url, obj)
	if err != nil {
		return fmt.Errorf("Error creating Address: %s", err)
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

	waitErr := computeOperationWait(config.clientCompute, op, project, "Creating Address")
	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return waitErr
	}

	return resourceComputeAddressRead(d, meta)
}

func resourceComputeAddressRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/regions/{{region}}/addresses/{{name}}")
	if err != nil {
		return err
	}

	res, err := Get(config, url)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeAddress %q", d.Id()))
	}

	d.Set("address", flattenComputeAddressAddress(res["address"]))
	d.Set("creation_timestamp", flattenComputeAddressCreationTimestamp(res["creationTimestamp"]))
	d.Set("description", flattenComputeAddressDescription(res["description"]))
	d.Set("id", flattenComputeAddressId(res["id"]))
	d.Set("name", flattenComputeAddressName(res["name"]))
	d.Set("users", flattenComputeAddressUsers(res["users"]))
	d.Set("region", flattenComputeAddressRegion(res["region"]))
	d.Set("self_link", res["selfLink"])
	d.Set("project", project)

	return nil
}

func resourceComputeAddressUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := map[string]interface{}{
		"address":     expandComputeAddressAddress(d.Get("address")),
		"description": expandComputeAddressDescription(d.Get("description")),
		"name":        expandComputeAddressName(d.Get("name")),
		"region":      expandComputeAddressRegion(d.Get("region")),
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/regions/{{region}}/addresses/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Address %q: %#v", d.Id(), obj)
	res, err := sendRequest(config, "PUT", url, obj)

	if err != nil {
		return fmt.Errorf("Error updating Address %q: %s", d.Id(), err)
	}
	err = Convert(res, &op)
	if err != nil {
		return err
	}

	err = computeOperationWait(config.clientCompute, &op, project, "Updating Address")
	if err != nil {
		return err
	}

	return resourceComputeAddressRead(d, meta)
}

func resourceComputeAddressDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/regions/{{region}}/addresses/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting Address %q", d.Id())
	res, err := Delete(config, url)
	if err != nil {
		return fmt.Errorf("Error deleting Address %q: %s", d.Id(), err)
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWait(config.clientCompute, op, project, "Deleting Address")
	if err != nil {
		return err
	}

	return nil
}

func resourceComputeAddressImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("", d.Id())
	return []*schema.ResourceData{d}, nil
}

func flattenComputeAddressAddress(v interface{}) interface{} {
	return v
}

func flattenComputeAddressCreationTimestamp(v interface{}) interface{} {
	return v
}

func flattenComputeAddressDescription(v interface{}) interface{} {
	return v
}

func flattenComputeAddressId(v interface{}) interface{} {
	return v
}

func flattenComputeAddressName(v interface{}) interface{} {
	return v
}

func flattenComputeAddressUsers(v interface{}) interface{} {
	return v
}

func flattenComputeAddressRegion(v interface{}) interface{} {
	return v
}

func expandComputeAddressAddress(v interface{}) interface{} {
	return v
}

func expandComputeAddressDescription(v interface{}) interface{} {
	return v
}

func expandComputeAddressName(v interface{}) interface{} {
	return v
}

func expandComputeAddressRegion(v interface{}) interface{} {
	return v
}

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

func resourceComputeBackendBucket() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeBackendBucketCreate,
		Read:   resourceComputeBackendBucketRead,
		Update: resourceComputeBackendBucketUpdate,
		Delete: resourceComputeBackendBucketDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeBackendBucketImport,
		},

		Schema: map[string]*schema.Schema{
			"bucket_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateRegexp(`^(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?)$`),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enable_cdn": {
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

func resourceComputeBackendBucketCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := map[string]interface{}{
		"bucketName":  expandComputeBackendBucketBucketName(d.Get("bucket_name")),
		"description": expandComputeBackendBucketDescription(d.Get("description")),
		"enableCdn":   expandComputeBackendBucketEnableCdn(d.Get("enable_cdn")),
		"name":        expandComputeBackendBucketName(d.Get("name")),
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/backendBuckets")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new BackendBucket: %#v", obj)
	res, err := Post(config, url, obj)
	if err != nil {
		return fmt.Errorf("Error creating BackendBucket: %s", err)
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

	waitErr := computeOperationWait(config.clientCompute, op, project, "Creating BackendBucket")
	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return waitErr
	}

	return resourceComputeBackendBucketRead(d, meta)
}

func resourceComputeBackendBucketRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/backendBuckets/{{name}}")
	if err != nil {
		return err
	}

	res, err := Get(config, url)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeBackendBucket %q", d.Id()))
	}

	d.Set("bucket_name", flattenComputeBackendBucketBucketName(res["bucketName"]))
	d.Set("creation_timestamp", flattenComputeBackendBucketCreationTimestamp(res["creationTimestamp"]))
	d.Set("description", flattenComputeBackendBucketDescription(res["description"]))
	d.Set("enable_cdn", flattenComputeBackendBucketEnableCdn(res["enableCdn"]))
	d.Set("name", flattenComputeBackendBucketName(res["name"]))
	d.Set("self_link", res["selfLink"])
	d.Set("project", project)

	return nil
}

func resourceComputeBackendBucketUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := map[string]interface{}{
		"bucketName":  expandComputeBackendBucketBucketName(d.Get("bucket_name")),
		"description": expandComputeBackendBucketDescription(d.Get("description")),
		"enableCdn":   expandComputeBackendBucketEnableCdn(d.Get("enable_cdn")),
		"name":        expandComputeBackendBucketName(d.Get("name")),
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/backendBuckets/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating BackendBucket %q: %#v", d.Id(), obj)
	res, err := sendRequest(config, "PUT", url, obj)

	if err != nil {
		return fmt.Errorf("Error updating BackendBucket %q: %s", d.Id(), err)
	}
	err = Convert(res, &op)
	if err != nil {
		return err
	}

	err = computeOperationWait(config.clientCompute, &op, project, "Updating BackendBucket")
	if err != nil {
		return err
	}

	return resourceComputeBackendBucketRead(d, meta)
}

func resourceComputeBackendBucketDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/backendBuckets/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting BackendBucket %q", d.Id())
	res, err := Delete(config, url)
	if err != nil {
		return fmt.Errorf("Error deleting BackendBucket %q: %s", d.Id(), err)
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWait(config.clientCompute, op, project, "Deleting BackendBucket")
	if err != nil {
		return err
	}

	return nil
}

func resourceComputeBackendBucketImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("{{name}}", d.Id())
	return []*schema.ResourceData{d}, nil
}

func flattenComputeBackendBucketBucketName(v interface{}) interface{} {
	return v
}

func flattenComputeBackendBucketCreationTimestamp(v interface{}) interface{} {
	return v
}

func flattenComputeBackendBucketDescription(v interface{}) interface{} {
	return v
}

func flattenComputeBackendBucketEnableCdn(v interface{}) interface{} {
	return v
}

func flattenComputeBackendBucketName(v interface{}) interface{} {
	return v
}

func expandComputeBackendBucketBucketName(v interface{}) interface{} {
	return v
}

func expandComputeBackendBucketDescription(v interface{}) interface{} {
	return v
}

func expandComputeBackendBucketEnableCdn(v interface{}) interface{} {
	return v
}

func expandComputeBackendBucketName(v interface{}) interface{} {
	return v
}

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
		"bucketName":  d.Get("bucket_name").(string),
		"description": d.Get("description").(string),
		"enableCdn":   d.Get("enable_cdn").(bool),
		"name":        d.Get("name").(string),
	}

	url, err := constructUrl(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/backendBuckets")
	if err != nil {
		return err
	}
	res, err := Post(config, url, obj)
	if err != nil {
		return fmt.Errorf("Error creating BackendBucket: %s", err)
	}

	// Store the ID now
	d.SetId(d.Get("name").(string))

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

	url, err := constructUrl(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/backendBuckets/{{name}}")
	if err != nil {
		return err
	}
	res, err := Get(config, url)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeBackendBucket %q", d.Id()))
	}

	d.Set("bucket_name", res["bucketName"])
	d.Set("creation_timestamp", res["creationTimestamp"])
	d.Set("description", res["description"])
	d.Set("enable_cdn", res["enableCdn"])
	d.Set("name", res["name"])
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
		"bucketName":  d.Get("bucket_name").(string),
		"description": d.Get("description").(string),
		"enableCdn":   d.Get("enable_cdn").(bool),
		"name":        d.Get("name").(string),
	}

	url, err := constructUrl(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/backendBuckets/{{name}}")
	if err != nil {
		return err
	}
	res, err := Put(config, url, obj)
	if err != nil {
		return fmt.Errorf("Error updating BackendBucket %q: %s", d.Id(), err)
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWait(config.clientCompute, op, project, "Updating BackendBucket")
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

	url, err := constructUrl(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/backendBuckets/{{name}}")
	if err != nil {
		return err
	}
	res, err := Delete(config, url)
	if err != nil {
		return fmt.Errorf("Error deleting BackendBucket %q: %s", d.Id(), err)
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWait(config.clientCompute, op, project, "Updating BackendBucket")
	if err != nil {
		return err
	}

	return nil
}

func resourceComputeBackendBucketImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("name", d.Id())
	return []*schema.ResourceData{d}, nil
}

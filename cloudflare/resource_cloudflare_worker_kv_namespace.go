package cloudflare

import (
	"fmt"
	"log"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareWorkerKvNamespace() *schema.Resource {
  return &schema.Resource{
    Create: resourceCloudflareWorkerKvNamespaceCreate,
    Read: resourceCloudflareWorkerKvNamespaceRead,
    Update: resourceCloudflareWorkerKvNamespaceUpdate,
    Delete: resourceCloudflareWorkerKvNamespaceDelete,
    Importer: &schema.ResourceImporter{
      State: resourceCloudflareWorkerKvNamespaceImport,
    }
    Schema: map[string]*schema.Schema{
      "title": {
        Type: schema.TypeString,
        Required: true,
        Description: "The title used to help you identify a worker kv namespace."
      }
    }
  }
}

func resourceCloudflareWorkerKvNamespaceCreate(d *schema.ResourceData, meta interface{}) error {
}

func resourceCloudflareWorkerKvNamespaceRead(d *schema.ResourceData, meta interface{}) error {
}

func resourceCloudflareWorkerKvNamespaceUpdate(d *schema.ResourceData, meta interface{}) error {
}

func resourceCloudflareWorkerKvNamespaceDelete(d *schema.ResourceData, meta interface{}) error {
}

func resourceCloudflareWorkerKvNamespaceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
}

package eventhub

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-provider-azurerm/internal/services/eventhub/sdk/2018-01-01-preview/eventhubsclusters"

	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/terraform-provider-azurerm/helpers/azure"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/location"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/timeouts"
)

func dataSourceEventHubCluster() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Read: dataSourceEventHubClusterRead,

		Timeouts: &pluginsdk.ResourceTimeout{
			Read: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"name": {
				Type:     pluginsdk.TypeString,
				Required: true,
			},

			"resource_group_name": azure.SchemaResourceGroupName(),

			"location": azure.SchemaLocationForDataSource(),

			"sku_name": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceEventHubClusterRead(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Eventhub.ClusterClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)
	resourceGroup := d.Get("resource_group_name").(string)

	subscriptionId := meta.(*clients.Client).Account.SubscriptionId
	id := eventhubsclusters.NewClusterID(subscriptionId, resourceGroup, name)
	resp, err := client.ClustersGet(ctx, id)
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return fmt.Errorf("%s was not found", id)
		}
		return fmt.Errorf("making Read request on Azure EventHub Cluster %q (Resource Group %q): %+v", name, resourceGroup, err)
	}
	d.SetId(id.ID())

	d.Set("name", id.ClusterName)
	d.Set("resource_group_name", id.ResourceGroupName)

	if model := resp.Model; model != nil {
		d.Set("sku_name", flattenEventHubClusterSkuName(model.Sku))
		d.Set("location", location.NormalizeNilable(model.Location))
	}

	return nil
}

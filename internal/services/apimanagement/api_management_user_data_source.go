// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagement

import (
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonschema"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/user"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/apimanagement/schemaz"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/timeouts"
)

func dataSourceApiManagementUser() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Read: dataSourceApiManagementUserRead,

		Timeouts: &pluginsdk.ResourceTimeout{
			Read: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"user_id": schemaz.SchemaApiManagementUserDataSourceName(),

			"api_management_name": schemaz.SchemaApiManagementDataSourceName(),

			"resource_group_name": commonschema.ResourceGroupNameForDataSource(),

			"first_name": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"email": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"last_name": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"note": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"state": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceApiManagementUserRead(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).ApiManagement.UsersClient
	subscriptionId := meta.(*clients.Client).Account.SubscriptionId
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id := user.NewUserID(subscriptionId, d.Get("resource_group_name").(string), d.Get("api_management_name").(string), d.Get("user_id").(string))

	resp, err := client.Get(ctx, id)
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return fmt.Errorf("%s was not found", id)
		}

		return fmt.Errorf("making read request %s: %+v", id, err)
	}

	d.SetId(id.ID())

	if model := resp.Model; model != nil {
		if props := model.Properties; props != nil {
			d.Set("first_name", pointer.From(props.FirstName))
			d.Set("last_name", pointer.From(props.LastName))
			d.Set("email", pointer.From(props.Email))
			d.Set("note", pointer.From(props.Note))
			d.Set("state", string(pointer.From(props.State)))
		}
	}

	return nil
}
package nexus

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NEXUS_REPOSITORY_MANAGER_URL", nil),
			},
			"auth_username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NEXUS_REPOSITORY_MANAGER_AUTH_USERNAME", nil),
			},
			"auth_password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NEXUS_REPOSITORY_MANAGER_AUTH_PASSWORD", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{},
	}
}

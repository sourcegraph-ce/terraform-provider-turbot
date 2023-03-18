package turbot

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-turbot/apiClient"
	log "github.com/sourcegraph-ce/logrus"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"secret_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"workspace": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"profile": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"credentials_file": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"turbot_policy_setting":          resourceTurbotPolicySetting(),
			"turbot_mod":                     resourceTurbotMod(),
			"turbot_folder":                  resourceTurbotFolder(),
			"turbot_resource":                resourceTurbotResource(),
			"turbot_local_directory":         resourceTurbotLocalDirectory(),
			"turbot_profile":                 resourceTurbotProfile(),
			"turbot_local_directory_user":    resourceTurbotLocalDirectoryUser(),
			"turbot_google_directory":        resourceGoogleDirectory(),
			"turbot_saml_directory":          resourceTurbotSamlDirectory(),
			"turbot_shadow_resource":         resourceTurbotShadowResource(),
			"turbot_smart_folder":            resourceTurbotSmartFolder(),
			"turbot_smart_folder_attachment": resourceTurbotSmartFolderAttachemnt(),
			"turbot_grant":                   resourceTurbotGrant(),
			"turbot_grant_activation":        resourceTurbotGrantActivation(),
			"turbot_turbot_directory":        resourceTurbotTurbotDirectory(),
			"turbot_file":                    resourceTurbotFile(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"turbot_policy_value": dataSourceTurbotPolicyValue(),
			"turbot_resource":     dataSourceTurbotResource(),
			"turbot_control":      dataSourceTurbotControl(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := apiClient.ClientConfig{
		Credentials: apiClient.ClientCredentials{
			AccessKey: d.Get("access_key").(string),
			SecretKey: d.Get("secret_key").(string),
			Workspace: d.Get("workspace").(string),
		},
		Profile:         d.Get("profile").(string),
		CredentialsPath: d.Get("credentials_file").(string),
	}

	client, err := apiClient.CreateClient(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %s", err.Error())
	}
	log.Println("[INFO] Turbot API client initialized, now validating...", client)
	if err = client.Validate(); err != nil {
		return nil, err
	}
	return client, nil
}

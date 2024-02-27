package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mrparkers/terraform-provider-keycloak/keycloak"
)

func resourceKeycloakRealmKeystoreGeneric() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceKeycloakRealmKeystoreGenericCreate,
		ReadContext:   resourceKeycloakRealmKeystoreGenericRead,
		UpdateContext: resourceKeycloakRealmKeystoreGenericUpdate,
		DeleteContext: resourceKeycloakRealmKeystoreGenericDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceKeycloakRealmKeystoreGenericImport,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Display name of provider when linked in admin console.",
			},
			"realm_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"active": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "TODO",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Set if the keys are enabled",
			},
			"priority": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Priority for the provider",
			},
			"provider_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "TODO",
			},
			"provider_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "TODO",
			},
		},
	}
}

func getRealmKeystoreGenericFromData(data *schema.ResourceData) (*keycloak.RealmKeystoreGeneric, error) {
	keystore := &keycloak.RealmKeystoreGeneric{
		Id:      data.Id(),
		Name:    data.Get("name").(string),
		RealmId: data.Get("realm_id").(string),

		Active:       data.Get("active").(bool),
		Enabled:      data.Get("enabled").(bool),
		Priority:     data.Get("priority").(int),
		ProviderId:   data.Get("provider_id").(string),
		ProviderType: data.Get("provider_type").(string),
	}

	return keystore, nil
}

func setRealmKeystoreGenericData(data *schema.ResourceData, realmKey *keycloak.RealmKeystoreGeneric) error {
	data.SetId(realmKey.Id)

	data.Set("name", realmKey.Name)
	data.Set("realm_id", realmKey.RealmId)

	data.Set("active", realmKey.Active)
	data.Set("enabled", realmKey.Enabled)
	data.Set("priority", realmKey.Priority)
	data.Set("provider_id", realmKey.ProviderId)
	data.Set("provider_type", realmKey.ProviderType)

	return nil
}

func resourceKeycloakRealmKeystoreGenericCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	keycloakClient := meta.(*keycloak.KeycloakClient)

	realmKey, err := getRealmKeystoreGenericFromData(data)
	if err != nil {
		return diag.FromErr(err)
	}

	err = keycloakClient.NewRealmKeystoreGeneric(ctx, realmKey)
	if err != nil {
		return diag.FromErr(err)
	}

	err = setRealmKeystoreGenericData(data, realmKey)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceKeycloakRealmKeystoreGenericRead(ctx, data, meta)
}

func resourceKeycloakRealmKeystoreGenericRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	keycloakClient := meta.(*keycloak.KeycloakClient)

	realmId := data.Get("realm_id").(string)
	id := data.Id()

	realmKey, err := keycloakClient.GetRealmKeystoreGeneric(ctx, realmId, id)
	if err != nil {
		return handleNotFoundError(ctx, err, data)
	}

	err = setRealmKeystoreGenericData(data, realmKey)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceKeycloakRealmKeystoreGenericUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	keycloakClient := meta.(*keycloak.KeycloakClient)

	realmKey, err := getRealmKeystoreGenericFromData(data)
	if err != nil {
		return diag.FromErr(err)
	}

	err = keycloakClient.UpdateRealmKeystoreGeneric(ctx, realmKey)
	if err != nil {
		return diag.FromErr(err)
	}

	err = setRealmKeystoreGenericData(data, realmKey)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceKeycloakRealmKeystoreGenericDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	keycloakClient := meta.(*keycloak.KeycloakClient)

	realmId := data.Get("realm_id").(string)
	id := data.Id()

	return diag.FromErr(keycloakClient.DeleteRealmKeystoreGeneric(ctx, realmId, id))
}

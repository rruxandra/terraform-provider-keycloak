package keycloak

import (
	"context"
	"fmt"
	"strconv"
)

type RealmKeystoreGeneric struct {
	Id      string
	Name    string
	RealmId string

	Active   bool
	Enabled  bool
	Priority int

	ProviderId   string
	ProviderType string
}

func convertFromRealmKeystoreGenericToComponent(realmKey *RealmKeystoreGeneric) *component {
	componentConfig := map[string][]string{
		"active": {
			strconv.FormatBool(realmKey.Active),
		},
		"enabled": {
			strconv.FormatBool(realmKey.Enabled),
		},
		"priority": {
			strconv.Itoa(realmKey.Priority),
		},
		"providerId": {
			realmKey.ProviderId,
		},
		"providerType": {
			realmKey.ProviderId,
		},
	}

	return &component{
		Id:           realmKey.Id,
		Name:         realmKey.Name,
		ParentId:     realmKey.RealmId,
		ProviderId:   realmKey.ProviderId,
		ProviderType: realmKey.ProviderType,
		Config:       componentConfig,
	}
}

func convertFromComponentToRealmKeystoreGeneric(component *component, realmId string) (*RealmKeystoreGeneric, error) {
	active, err := parseBoolAndTreatEmptyStringAsFalse(component.getConfig("active"))
	if err != nil {
		return nil, err
	}

	enabled, err := parseBoolAndTreatEmptyStringAsFalse(component.getConfig("enabled"))
	if err != nil {
		return nil, err
	}

	priority := 0 // Default priority
	if component.getConfig("priority") != "" {
		priority, err = strconv.Atoi(component.getConfig("priority"))
		if err != nil {
			return nil, err
		}
	}

	// TODO validate providerId and providerType
	providerId := component.getConfig("providerId")
	providerType := component.getConfig("providerType")

	realmKey := &RealmKeystoreGeneric{
		Id:      component.Id,
		Name:    component.Name,
		RealmId: realmId,

		Active:       active,
		Enabled:      enabled,
		Priority:     priority,
		ProviderId:   providerId,
		ProviderType: providerType,
	}

	return realmKey, nil
}

func (keycloakClient *KeycloakClient) NewRealmKeystoreGeneric(ctx context.Context, realmKey *RealmKeystoreGeneric) error {
	_, location, err := keycloakClient.post(ctx, fmt.Sprintf("/realms/%s/components", realmKey.RealmId), convertFromRealmKeystoreGenericToComponent(realmKey))
	if err != nil {
		return err
	}

	realmKey.Id = getIdFromLocationHeader(location)

	return nil
}

func (keycloakClient *KeycloakClient) GetRealmKeystoreGeneric(ctx context.Context, realmId, id string) (*RealmKeystoreGeneric, error) {
	var component *component

	err := keycloakClient.get(ctx, fmt.Sprintf("/realms/%s/components/%s", realmId, id), &component, nil)
	if err != nil {
		return nil, err
	}

	return convertFromComponentToRealmKeystoreGeneric(component, realmId)
}

func (keycloakClient *KeycloakClient) UpdateRealmKeystoreGeneric(ctx context.Context, realmKey *RealmKeystoreGeneric) error {
	return keycloakClient.put(ctx, fmt.Sprintf("/realms/%s/components/%s", realmKey.RealmId, realmKey.Id), convertFromRealmKeystoreGenericToComponent(realmKey))
}

func (keycloakClient *KeycloakClient) DeleteRealmKeystoreGeneric(ctx context.Context, realmId, id string) error {
	return keycloakClient.delete(ctx, fmt.Sprintf("/realms/%s/components/%s", realmId, id), nil)
}

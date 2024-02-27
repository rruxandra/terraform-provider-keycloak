---
page_title: "keycloak_realm_keystore_generic Resources"
---

# keycloak\_realm\_keystore\_generic Resources

Allows for creating and managing generic Realm keystores within Keycloak.

A realm keystore manages keys that are used by Keycloak to perform cryptographic signatures and encryption.

## Example Usage

```hcl
resource "keycloak_realm" "realm" {
	realm = "my-realm"
}

resource "keycloak_realm_keystore_generic" "keystore_generic" {
	name      = "my-generic-keystore"
	realm_id  = keycloak_realm.realm.id

	enabled = true
	active  = true

	priority       = 100
	provider_id    = "custom-keystore"
	provider_type  = "org.company.keys.KeyProvider"
}
```

## Argument Reference

- `name` - (Required) Display name of provider when linked in admin console.
- `realm_id` - (Required) The realm this keystore exists in.
- `enabled` - (Optional) When `false`, key is not accessible in this realm. Defaults to `true`.
- `active` - (Optional) When `false`, key in not used for signing. Defaults to `true`.
- `priority` - (Optional) Priority for the provider. Defaults to `0`
- `providerId` - (Required) The ID of the keystore provider.
- `providerType` - (Required) The type of the keystore provider.

## Import

Realm keys can be imported using realm name and keystore id, you can find it in web UI.

Example:

```bash
$ terraform import keycloak_realm_keystore_generic.keystore_generic my-realm/618cfba7-49aa-4c09-9a19-2f699b576f0b
```

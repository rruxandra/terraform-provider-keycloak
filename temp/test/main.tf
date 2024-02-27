provider "keycloak" {
  client_id = "admin-cli"
  username  = "admin"
  password  = "admin"
  url       = "http://localhost:8080"
}

resource "keycloak_realm" "realm" {
  realm = "my-realm"
}

resource "keycloak_realm_keystore_generic" "keystore_generic" {
  name     = "my-generic-keystore"
  realm_id = keycloak_realm.realm.id

  enabled = true
  active  = true

  priority      = 100
  provider_id   = "rsa-generated"
  provider_type = "org.keycloak.keys.KeyProvider"
}

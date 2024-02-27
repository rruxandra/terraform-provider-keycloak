# Test Local Terraform Plugin

1. Create the folder where the local Terraform plugin will be installed: `~/.terraform.d/plugins/<HOSTNAME>/<NAMESPACE>/<NAME>/<VERSION>/<OS_ARCH>/`
Example: `mkdir -p ~/.terraform.d/plugins/terraform.local/local/keycloak/4.4.1/darwin_arm64/`

2. Copy [terraform-provider-keycloak_v4.4.1](terraform-provider-keycloak_v4.4.1) into the directory created in the previous step.

3. Run Keycloak: `docker compose up`.

4. Run `terraform init` and `terraform apply` in `test` directory.

5. Go to `http://localhost:8080/admin/master/console/#/my-realm/realm-settings/keys` and check that the provider named `my-generic-keystore` exists.

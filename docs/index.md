---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "singlestoredb Provider"
subcategory: ""
description: |-
  The Terraform provider plugin for managing SingleStoreDB workspace groups and workspaces.
---

# singlestoredb Provider

The Terraform provider plugin for managing SingleStoreDB workspace groups and workspaces.

## Example Usage

```terraform
provider "singlestoredb" {
  // The SingleStoreDB Terraform provider uses the SINGLESTOREDB_API_KEY environment variable for authentication. 
  // Please set this environment variable with your SingleStore Management API key.
  // You can generate this key from the SingleStore Portal at https://portal.singlestore.com/organizations/org-id/api-keys.
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `api_key` (String, Sensitive) The SingleStore Management API key. This key is used for authentication. If this key is not set, the value from the environment variable SINGLESTOREDB_API_KEY will be used. You can generate this key from the SingleStore Portal at https://portal.singlestore.com/organizations/org-id/api-keys.
- `api_service_url` (String, Deprecated) The URL of the SingleStore Management API service. This URL is used by the provider to interact with the API.
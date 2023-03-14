# regions

`workspacegroups/get` shows workspace group lookup by ID.

~~~ shell
# Replace `workspace_group_id` in `main.tf` with the ID of the workspace group that exists.
# To fetch the ID, consider leveraging the `workspace_groups` data source.
terraform apply -auto-approve
~~~

**Note: `terraform init` does not work with `dev_overrides` for local development.**
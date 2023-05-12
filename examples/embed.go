package examples

import (
	"embed"
	"fmt"
)

//go:embed data-sources/*/data-source.tf resources/*/resource.tf
var f embed.FS

var (
	Regions                       = mustRead("data-sources/singlestoredb_regions/data-source.tf")
	WorkspaceGroupsListDataSource = mustRead("data-sources/singlestoredb_workspace_groups/data-source.tf")
	WorkspaceGroupsGetDataSource  = mustRead("data-sources/singlestoredb_workspace_group/data-source.tf")
	WorkspacesListDataSource      = mustRead("data-sources/singlestoredb_workspaces/data-source.tf")
	WorkspacesGetDataSource       = mustRead("data-sources/singlestoredb_workspace/data-source.tf")
	WorkspaceGroupsResource       = mustRead("resources/singlestoredb_workspace_group/resource.tf")
	WorkspacesResource            = mustRead("resources/singlestoredb_workspace/resource.tf")
)

func mustRead(path string) string {
	result, err := f.ReadFile(path)
	if err != nil {
		panic(err)
	}

	if string(result) == "" {
		panic(fmt.Sprintf("%s should have content but is empty", path))
	}

	return string(result)
}

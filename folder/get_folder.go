package folder

import (
	// "fmt"
	"slices"
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {

	org_folders := f.GetFoldersByOrgID(orgID)
	child_folders := []Folder{}


	for _, f := range org_folders {
		path_split := strings.Split(f.Paths, ".")
		path_last := len(path_split) - 1
		if slices.Contains(path_split, name) && path_split[path_last] != name{
			child_folders = append(child_folders, f)
			
		}
	}
	
	return child_folders
}

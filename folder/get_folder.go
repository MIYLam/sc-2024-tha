package folder

import (
	// "fmt"
	"errors"
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

func (f *driver) FindFolderByName(name string, folder_range []Folder) Folder{
	var target_folder Folder
	for _, f := range folder_range{
		if f.Name == name{
			target_folder = f
			break
		}
	}
	return target_folder
}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {

	org_folders := f.GetFoldersByOrgID(orgID)
	child_folders := []Folder{}
	parent := f.FindFolderByName(name, org_folders)

	if (parent.Name == ""){
		return nil, errors.New("folder does not exist")
	}

	for _, f := range org_folders {
		path_split := strings.Split(f.Paths, ".")
		path_last := len(path_split) - 1
		if slices.Contains(path_split, name) && path_split[path_last] != name && parent.OrgId == f.OrgId{
			child_folders = append(child_folders, f)
			
		}
	}
	
	return child_folders, nil
}


